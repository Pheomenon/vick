package vick

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"math"
	"os"
	"sync"
	"sync/atomic"

	proto_memtable "github.com/Pheomenon/vick/proto/schemas/memtable"
	"github.com/sirupsen/logrus"
)

var CrcTable = crc32.MakeTable(crc32.Castagnoli)

type memtable struct {
	buf           []byte
	currentOffset int
	minRange      uint32
	maxRange      uint32
	concurrentMap map[uint32]uint32
	size          int
	records       uint32
	name          uint32
	sync.RWMutex
}

func newMemtable(size int, name uint32) *memtable {
	return &memtable{
		buf:           make([]byte, size),
		concurrentMap: make(map[uint32]uint32, 0),
		minRange:      math.MaxUint32,
		maxRange:      0,
		name:          name,
		size:          size,
		RWMutex:       sync.RWMutex{},
	}
}

func (m *memtable) Set(key, value []byte) {
	m.Lock()
	defer m.Unlock()
	hash := hashing(key)
	oldOffset := m.currentOffset
	keyLen := len(key)
	valLen := len(value)

	binary.LittleEndian.PutUint32(m.buf[m.currentOffset:], uint32(keyLen))
	m.currentOffset += 4

	binary.LittleEndian.PutUint32(m.buf[m.currentOffset:], uint32(valLen))
	m.currentOffset += 4

	copy(m.buf[m.currentOffset:m.currentOffset+keyLen], key)
	m.currentOffset += keyLen

	copy(m.buf[m.currentOffset:m.currentOffset+valLen], value)
	m.currentOffset += valLen

	m.concurrentMap[hash] = uint32(oldOffset)
	if hash > m.maxRange {
		m.maxRange = hash
	} else if hash < m.minRange {
		m.minRange = hash
	}

	atomic.AddUint32(&m.records, 1)
}

func (m *memtable) Get(item []byte) ([]byte, bool) {
	m.RLock()
	defer m.RUnlock()
	hash := hashing(item)
	position, ok := m.concurrentMap[hash]
	if !ok {
		return nil, false
	}

	keyLen := binary.LittleEndian.Uint32(m.buf[position : position+4])
	position += 4
	key := m.buf[position+4 : position+4+keyLen]
	if bytes.Compare(key, item) != 0 {
		return nil, false
	}

	valLen := binary.LittleEndian.Uint32(m.buf[position : position+4])
	position += 4
	start := position + keyLen
	end := start + valLen
	return m.buf[start:end], true
}

func (m *memtable) isFull(size int) bool {
	m.RLock()
	defer m.RUnlock()
	left := m.size - m.currentOffset
	if left < size {
		return false
	}
	return true
}

func (m *memtable) dump(path string) {
	m.Lock()
	defer m.Unlock()
	dataFp, err := os.Create(fmt.Sprintf("%s/%d.wic", path, m.name))
	if err != nil {
		logrus.Errorf("memtable: unable to create file: [%v]", err)
	}
	idxFp, err := os.Create(fmt.Sprintf("%s/%d.idx", path, m.name))
	if err != nil {
		logrus.Errorf("memtable: unable to create file: [%v]", err)
	}
	defer func(data, idx *os.File) {
		err := data.Close()
		if err != nil {
			logrus.Errorf("memtable: unable to close file: [%v]", err)
		}
		err = idx.Close()
		if err != nil {
			logrus.Errorf("mumtable unable to close file: [%v]", err)
		}
	}(dataFp, idxFp)

	positionRuler := uint32(0)
	ds := &proto_memtable.DataSchemas{}
	is := &proto_memtable.IndexSchemas{}
	for hash, position := range m.concurrentMap {
		keyLen := binary.LittleEndian.Uint32(m.buf[position : position+4])
		valLen := binary.LittleEndian.Uint32(m.buf[position+4 : position+8])
		ds.Es = append(ds.Es, &proto_memtable.EntrySchemas{
			KeyLen: keyLen,
			ValLen: valLen,
			Key:    m.buf[position+8 : position+8+keyLen],
			Val:    m.buf[position+8+keyLen : position+8+keyLen+valLen],
		})
		positionRuler = positionRuler + 8 + keyLen + valLen
		m.concurrentMap[hash] = positionRuler

		is.P = append(is.P, &proto_memtable.PositionSchemas{
			Crc32:  hash,
			Offset: position,
		})
	}

	info := &proto_memtable.InfoSchemas{
		MetaOffset: positionRuler,
		Entries:    int32(len(m.concurrentMap)),
		MinRange:   m.minRange,
		MaxRange:   m.maxRange,
	}
	ds.Info = info

	table, err := ds.Marshal()
	_, err = dataFp.Write(table)
	if err != nil {
		logrus.Errorf("memtable: dump memtable error: [%v]", err)
	}

	idx, err := is.Marshal()
	_, err = idxFp.Write(idx)
	if err != nil {
		logrus.Errorf("memtable: marshal memtable error: [%v]", err)
	}
}
