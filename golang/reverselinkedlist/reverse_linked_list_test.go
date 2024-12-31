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

	for expected != nil {
		if expected.Val != actual.Val {
			t.Error("reverseList() function returned incorrect result.")
		}		
		expected = expected.Next
		actual = actual.Next
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

	for expected != nil {
		if expected.Val != actual.Val {
			t.Error("reverseList() function returned incorrect result.")
		}		
		expected = expected.Next
		actual = actual.Next
	}

	// Example thrVale
	actual = reverseList(nil)

	if actual != nil {
		t.Error("reverseList() function returned incorrect result.")
	}		
}