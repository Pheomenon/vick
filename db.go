package vick

type DB struct {
	Memtable *memtable
}

func NewDB(size int, name uint32) *DB {
	m := newMemtable(size, name)
	return &DB{
		Memtable: m,
	}
}
