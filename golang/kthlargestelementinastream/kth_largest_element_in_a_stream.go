/*
703 - Easy

You are part of a university admissions office and need to keep track of the kth highest test score from applicants in real-time.
This helps to determine cut-off marks for interviews and admissions dynamically as new applicants submit their scores.
You are tasked to implement a class which, for a given integer k, maintains a stream of test scores and continuously returns the kth highest
test score after a new score has been submitted. More specifically, we are looking for the kth highest score in the sorted list of all scores.

Implement the KthLargest class:

  - KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of test scores nums.
  - int add(int val) Adds a new test score val to the stream and returns the element representing the kth largest element in the pool of test scores so far.

Example 1:

	Input:

		["KthLargest", "add", "add", "add", "add", "add"]
		[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]

	Output:

		[null, 4, 5, 5, 8, 8]

	Explanation:

		KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
		kthLargest.add(3); // return 4
		kthLargest.add(5); // return 5
		kthLargest.add(10); // return 5
		kthLargest.add(9); // return 8
		kthLargest.add(4); // return 8

Example 2:

	Input:

		["KthLargest", "add", "add", "add", "add"]
		[[4, [7, 7, 7, 7, 8, 3]], [2], [10], [9], [9]]

	Output:

		[null, 7, 7, 7, 8]

	Explanation:

		KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);
		kthLargest.add(2); // return 7
		kthLargest.add(10); // return 7
		kthLargest.add(9); // return 7

Constraints:

  - 0 <= nums.length <= 10^4
  - 1 <= k <= nums.length + 1
  - -10^4 <= nums[i] <= 10^4
  - -10^4 <= val <= 10^4
  - At most 10^4 calls will be made to add.
*/
package kthlargestelementinastream

import "container/heap"

type Heap []int

func (h Heap) Len() int { return len(h) } 
func (h Heap) Less(i, j int) bool { return h[i] < h[j] } 
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] } 

func (h *Heap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *Heap) Pop() any {
	val := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return val	
}

type KthLargest struct {
	K int
	Heap *Heap
}


func Constructor(k int, nums []int) KthLargest {
	h := &Heap{}
	heap.Init(h)

	kthLargest := KthLargest{ K: k, Heap: h }

	for _, num := range nums {
		kthLargest.Add(num)
	}

	return kthLargest
}

func (kl *KthLargest) Add(val int) int {
	if kl.Heap.Len() < kl.K {
		heap.Push(kl.Heap, val)
	} else if (*kl.Heap)[0] < val {
		heap.Pop(kl.Heap)
		heap.Push(kl.Heap, val)	
	}
	
	if kl.Heap.Len() > 0 {
		return (*kl.Heap)[0]
	}
	
	return 0
}

/* ---- Slice sorting solution ----

import (
	"sort"
)

type KthLargest struct {
	K int
	Nums []int
}


func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)

	if len(nums)-k > 0 {
		nums = nums[len(nums)-k:]
	}

	return KthLargest{
		K: k,
		Nums: nums,
	}
}

func (kl *KthLargest) Add(val int) int {
	if len(kl.Nums) < kl.K || len(kl.Nums) == 0 {
		kl.Nums = append(kl.Nums, val)
	} else if val > kl.Nums[0] {
		kl.Nums[0] = val
	}
	sort.Ints(kl.Nums)

	return kl.Nums[0]
}
*/

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
