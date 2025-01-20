package copylistwithrandompointer

import (
	"reflect"
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestCopyRandomList(t *testing.T) {
	// Example one
	head := &cs.Node{Val: 7}

	node1 := &cs.Node{Val: 13}
	node2 := &cs.Node{Val: 11}
	node3 := &cs.Node{Val: 10}
	node4 := &cs.Node{Val: 1}

	head.Next = node1

	node1.Next = node2
	node1.Random = head

	node2.Next = node3
	node2.Random = node4 

	node3.Next = node4
	node3.Random = node2

	node4.Random = head
	
	copiedList := copyRandomList(head)

	if !reflect.DeepEqual(head, copiedList) {
		t.Error("copyRandomList(head) returned incorrect result for example 1.")
	}
	
	// Example two
	head = &cs.Node{Val: 7}

	node1 = &cs.Node{Val: 13}

	head.Next = node1
	head.Random = node1

	node1.Random = node1

	copiedList = copyRandomList(head)

	if !reflect.DeepEqual(head, copiedList) {
		t.Error("copyRandomList(head) returned incorrect result for example 2.")
	}

	// Example three
	head = &cs.Node{Val: 7}

	node1 = &cs.Node{Val: 13}
	node2 = &cs.Node{Val: 11}

	head.Next = node1

	node1.Next = node2
	node1.Random = head

	copiedList = copyRandomList(head)

	if !reflect.DeepEqual(head, copiedList) {
		t.Error("copyRandomList(head) returned incorrect result for example 3.")
	}
}