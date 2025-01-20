package reverselinkedlist

import (
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestReverseList(t *testing.T) {
	// Example one
	head := & cs.ListNode{
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
		Val: 5,
		Next: &cs.ListNode{
			Val: 4,
			Next: &cs.ListNode{
				Val: 3,
				Next: &cs.ListNode{
					Val: 2,
					Next: &cs.ListNode{
						Val: 1,
					},
				},
			},
		},
	}

	actual := reverseList(head)

	if !expected.Equals(actual) {
		t.Error("reverseList(head) function returned incorrect result for example 1.")
	}		

	// Example two
	head = &cs.ListNode{
		Val: 1,
		Next: &cs.ListNode{
			Val: 2,
		},
	}

	expected = &cs.ListNode{
		Val: 2,
		Next: &cs.ListNode{
			Val: 1,
		},
	}

	actual = reverseList(head)

	if !expected.Equals(actual) {
		t.Error("reverseList(head) function returned incorrect result for example 2.")
	}		

	// Example three
	actual = reverseList(nil)

	if actual != nil {
		t.Error("reverseList(head) function returned incorrect result for example 3.")
	}		
}