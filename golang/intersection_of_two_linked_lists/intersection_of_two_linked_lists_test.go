package intersectionoftwolinkedlists

import (
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestGetIntersectionNode(t *testing.T) {

	// Example one
	expected := &cs.ListNode{
		Val: 8,
		Next: &cs.ListNode{
			Val: 4,
			Next: &cs.ListNode{
				Val: 5,
			},
		},
	}

	listA := &cs.ListNode{
		Val: 4,
		Next: &cs.ListNode{
			Val:  1,
			Next: expected,
		},
	}

	listB := &cs.ListNode{
		Val: 5,
		Next: &cs.ListNode{
			Val: 6,
			Next: &cs.ListNode{
				Val:  1,
				Next: expected,
			},
		},
	}

	actual := getIntersectionNode(listA, listB)

	if actual != expected {
		t.Errorf("getIntersectionNode func returned incorrect value.")
	}

	// Example two
	expected = &cs.ListNode{
		Val: 2,
		Next: &cs.ListNode{
			Val: 4},
	}

	listA = &cs.ListNode{
		Val: 1,
		Next: &cs.ListNode{
			Val: 9,
			Next: &cs.ListNode{
				Val:  1,
				Next: expected,
			},
		},
	}

	listB = &cs.ListNode{
		Val: 3,
		Next: expected,
	}

	actual = getIntersectionNode(listA, listB)

	if actual != expected {
		t.Errorf("getIntersectionNode func returned incorrect value.")
	}
	
	// Example three
	listA = &cs.ListNode{
		Val: 2,
		Next: &cs.ListNode{
			Val: 6,
			Next: &cs.ListNode{
				Val: 4,
			},
		},
	}

	listB = &cs.ListNode{
		Val: 1,
		Next: &cs.ListNode{
			Val:  5,
		},
	}

	actual = getIntersectionNode(listA, listB)

	if actual != nil {
		t.Errorf("getIntersectionNode func returned incorrect value.")
	}
}
