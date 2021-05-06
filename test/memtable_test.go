package test

import (
	"fmt"
	"testing"

	"github.com/Pheomenon/vick"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MemtableTest struct{}

var _ = Suite(&MemtableTest{})

func (m *MemtableTest) TestOpenMemtable(c *C) {
	db := vick.NewDB(1024, 0)
	db.Memtable.Set([]byte("Key1"), []byte("val1"))
	val, _ := db.Memtable.Get([]byte("Key1"))
	fmt.Sprintf("%v", val)
	c.Assert(val, DeepEquals, []byte("val1"))
}
