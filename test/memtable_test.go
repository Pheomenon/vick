package test

import (
	"testing"

	"github.com/Pheomenon/vick"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MemtableTest struct{}

var _ = Suite(&MemtableTest{})

func (m *MemtableTest) TestOpenMemtable(c *C) {
	db := vick.NewDB(1024, 0)
	db.Memtable.Set([]byte("Key1"), []byte("val1"))
	val, _ := db.Memtable.Get([]byte("Key1"))
	c.Assert(val, DeepEquals, []byte("val1"))
}

func (m *MemtableTest) TestPersistence(c *C) {
	db := vick.NewDB(1024, 0)
	db.Memtable.Set([]byte("Key1"), []byte("val1"))
	db.Memtable.Dump("./")
}

func (m *MemtableTest) TestDel(c *C) {
	db := vick.NewDB(1024, 0)
	db.Memtable.Set([]byte("Key1"), []byte("val1"))
	val, _ := db.Memtable.Get([]byte("Key1"))
	c.Assert(val, DeepEquals, []byte("val1"))
	db.Memtable.Del([]byte("Key1"))
	val, _ = db.Memtable.Get([]byte("Key1"))
	c.Assert(val, DeepEquals, []byte(nil))
}
