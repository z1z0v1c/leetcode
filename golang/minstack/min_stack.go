/**
 * 155 - Medium
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
 *
 * Implement the MinStack class:
 *
 * 	- MinStack() initializes the stack object.
 * 	- void push(int val) pushes the element val onto the stack.
 * 	- void pop() removes the element on the top of the stack.
 * 	- int top() gets the top element of the stack.
 * 	- int getMin() retrieves the minimum element in the stack.
 *
 * You must implement a solution with O(1) time complexity for each function.
 *
 *
 * Example 1:
 *
 * 	Input
 *
 * 		["MinStack","push","push","push","getMin","pop","top","getMin"]
 * 		[[],[-2],[0],[-3],[],[],[],[]]
 *
 * 	Output
 *
 * 		[null,null,null,null,-3,null,0,-2]
 *
 * 	Explanation
 *
 * 		MinStack minStack = new MinStack();
 * 		minStack.push(-2);
 * 		minStack.push(0);
 * 		minStack.push(-3);
 * 		minStack.getMin(); // return -3
 * 		minStack.pop();
 * 		minStack.top();    // return 0
 * 		minStack.getMin(); // return -2
 *
 * Constraints:
 *
 * 	- -2^31 <= val <= 2^31 - 1
 * 	- Methods pop, top and getMin operations will always be called on non-empty stacks.
 * 	- At most 3 * 10^4 calls will be made to push, pop, top, and getMin.
 */
package minstack

type MinStack struct {
	min []int
	stack []int
}


func Constructor() MinStack {
	return MinStack{
		min: []int{},
		stack: []int{},
	}
}


func (ms *MinStack) Push(val int)  {
	ms.stack = append(ms.stack, val)
	// Append element to min array if it's less or equal to the previous one
	if len(ms.min) == 0 || ms.min[len(ms.min)-1] >= val {
		ms.min = append(ms.min, val)
	}
}


func (ms *MinStack) Pop()  {
	top := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	// Remove the top element from min array if present
	if ms.min[len(ms.min)-1] == top {
		ms.min = ms.min[:len(ms.min)-1]
	}
}


func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}


func (ms *MinStack) GetMin() int {
	return ms.min[len(ms.min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */