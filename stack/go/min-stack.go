// https://leetcode.com/problems/min-stack
// TODO: solve again
package main

import "fmt"

func main() {
	stack := Constructor()

	stack.Push(-2)
	fmt.Println(stack)
	stack.Push(0)
	fmt.Println(stack)
	stack.Push(-3)
	fmt.Println(stack)
	stack.GetMin()
	stack.Pop()
	fmt.Println(stack)
	stack.Top()
	stack.GetMin()
}

type MinStack struct {
	stackSlice []int
	minSlice   []int
	topPointer int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}, -1}
}

func (this *MinStack) Push(val int) {
	this.stackSlice = append(this.stackSlice, val)

	if this.topPointer < 0 {
		this.minSlice = append(this.minSlice, val)
		this.topPointer++

		return
	}

	if val < this.minSlice[this.topPointer] {
		this.minSlice = append(this.minSlice, val)
		this.topPointer++

		return
	}

	this.minSlice = append(this.minSlice, this.minSlice[this.topPointer])
	this.topPointer++
}

func (this *MinStack) Pop() {
	this.stackSlice = this.stackSlice[:len(this.stackSlice)-1]
	this.minSlice = this.minSlice[:len(this.minSlice)-1]
	this.topPointer--
}

func (this *MinStack) Top() int {
	return this.stackSlice[this.topPointer]
}

func (this *MinStack) GetMin() int {
	return this.minSlice[this.topPointer]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
