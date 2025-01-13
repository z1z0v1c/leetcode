package minstack

import "testing"

func TestMinStack(t *testing.T) {
	obj := Constructor()

	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	min := obj.GetMin()
	if min != -3 {
		t.Errorf("obj.GetMin() returned %d instead -3.", min)
	}

	obj.Pop()

	top := obj.Top()
	if top != 0 {
		t.Errorf("obj.Top() returned %d instead 0.", top)
	}

	min = obj.GetMin()
	if min != -2 {
		t.Errorf("obj.GetMin() returned %d instead -2.", min)
	}
}