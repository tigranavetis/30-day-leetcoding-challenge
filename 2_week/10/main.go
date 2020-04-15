package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	min      int
	elements []int
}

func Constructor() MinStack {
	min := int(math.Inf(0)) - 1
	minStack := MinStack{min, make([]int, 0)}
	return minStack
}

func (this *MinStack) Push(x int) {
	this.elements = append(this.elements, x)
	if x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	lastElIndex := len(this.elements) - 1
	lastEl := this.elements[lastElIndex]
	this.elements = this.elements[:len(this.elements)-1]

	if lastEl == this.min {
		this.min = int(math.Inf(0)) - 1
		for _, v := range this.elements {
			if v < this.min {
				this.min = v
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
}
