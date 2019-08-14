package stack

type StackEl struct {
    value int
	next *StackEl
}

type Stack struct {
    top *StackEl
	size int
}

func New() *Stack {
    var p Stack;
	p.size = 0
	p.top = nil;
	return &p
}

func (p *Stack) Push(value int) {
	var top StackEl
	top.value = value
	p.size += 1
	if (p.top == nil) {
		top.next = nil
	} else {
	    top.next = p.top
	}
	p.top = &top 
}

func (p *Stack) Pop() int {
    value:=p.top.value
    p.size-=1
    p.top = p.top.next
    return value
}
