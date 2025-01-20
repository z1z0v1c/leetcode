package reorderlist

import (
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestReorderList(t *testing.T) {
	// Example one
	head := &cs.ListNode{
		Val: 1,
		Next: &cs.ListNode{
			Val: 2,
			Next: &cs.ListNode{
				Val: 3,
				Next: &cs.ListNode{
					Val: 4,
				},
			},
		},
	}
	
	expected := &cs.ListNode{
		Val: 1,
		Next: &cs.ListNode{
			Val: 4,
			Next: &cs.ListNode{
				Val: 2,
				Next: &cs.ListNode{
					Val: 3,
				},
			},
		},
	}

	reorderList(head)

	if !expected.Equals(head) {
		t.Error("reorderList(head) returned incorrect resut for example 1.")
	}

	// Example two
	head = &cs.ListNode{
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
	
	expected = &cs.ListNode{
		Val: 1,
		Next: &cs.ListNode{
			Val: 5,
			Next: &cs.ListNode{
				Val: 2,
				Next: &cs.ListNode{
					Val: 4,
					Next: &cs.ListNode{
						Val: 3,
					},
				},
			},
		},
	}

	reorderList(head)

	if !expected.Equals(head) {
		t.Error("reorderList(head) returned incorrect resut for example 2.")
	}
}