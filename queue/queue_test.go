package queue

import (
	"github.com/tornyak/goalg/util"
	"testing"
)

func TestQueue(t *testing.T) {
	testCreateEmpty(t)
	//testAddOne(t)
	//testAddMultiple(t)
	//testAddDifferentTypes(t)
	//testIterateEmpty(t)
	//testIterateOne(t)
	//testIterateMultiple(t)
}

func testCreateEmpty(t *testing.T) {
	q := New()
	if !q.IsEmpty() {
		t.Error(util.GetFuncName(), ":Queue not empty")
	}
	if q.Size() != 0 {
		t.Errorf("%q :size expected 0 received %d", util.GetFuncName(), q.Size())
	}
}
