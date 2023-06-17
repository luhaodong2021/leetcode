package minstack

type MinStack struct {
	len   int
	min   []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.len == 0 {
		this.min = append(this.min, 0)
		this.len++
		return
	}
	if val <= this.GetMin() {
		this.min = append(this.min, this.len)
	} else {
		this.min = append(this.min, this.min[this.len-1])
	}
	this.len++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.len-1]
	this.min = this.min[:this.len-1]
	this.len--
}

func (this *MinStack) Top() int {
	return this.stack[this.len-1]
}

func (this *MinStack) GetMin() int {
	return this.stack[this.min[this.len-1]]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
