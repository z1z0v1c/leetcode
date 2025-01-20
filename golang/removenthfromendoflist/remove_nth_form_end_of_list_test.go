package removenthfromendoflist

import (
	"reflect"
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestRemoveNthFromEnd(t *testing.T) {
	// Example one
	head := &cs.ListNode{
		Val: 1,
		Next: &cs.ListNode{
			Val: 2,
			Next: &cs.ListNode{
				Val: 3,
				Next: &cs.ListNode{
					Val: 4,
					Next: &cs.ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	
	expected := &cs.ListNode{
		Val: 1,
		Next: &cs.ListNode{
			Val: 2,
			Next: &cs.ListNode{
				Val: 3,
				Next: &cs.ListNode{
					Val: 5,
				},
			},
		},
	}

	actual := removeNthFromEnd(head, 2)

	if !reflect.DeepEqual(actual, expected) {
		t.Error("reorderList(head, n) returned incorrect resut for example 1.")
	}

	// Example two
	head = &cs.ListNode{
		Val: 1,
	}
	
	expected = nil
	actual = removeNthFromEnd(head, 1)

	if actual != expected {
		t.Error("reorderList(head, n) returned incorrect resut for example 2.")
	}

	// Example three
	head = &cs.ListNode{
		Val: 1,
		Next: &cs.ListNode{
			Val: 2,
		},
	}
	
	expected = &cs.ListNode{
		Val: 1,
	}

	actual = removeNthFromEnd(head, 1)

	if !reflect.DeepEqual(actual, expected) {
		t.Error("reorderList(head, n) returned incorrect resut for example 3.")
	}
}