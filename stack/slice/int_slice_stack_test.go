package slice

import (
	"testing"
	"reflect"
)

func TestIntStack_ToSlice(t *testing.T) {
	stack := Stack()
	exp := []int{1,2,3,4,5}

	t.Log(stack)
	stack.Push(exp...)
	if !reflect.DeepEqual(stack.ToSlice(), exp) {
		t.Errorf("ToSlice failed: %v got: %v", exp, stack.ToSlice())
	}
}
