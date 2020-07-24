package stack

import "fmt"

type Stack struct {
	top  int
	data []interface{}
}

func New() Stack {
	st := Stack{top: -1}
	return st
}
func (st *Stack) Push(val interface{}) {
	st.data = append(st.data, val)
	st.top++
}
func (st *Stack) Pop() interface{} {
	val := st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	st.top--
	return val
}
func (st *Stack) Size() int {
	return len(st.data)
}
func (st *Stack) IsEmpty() bool {
	if st.top == -1 {
		return true
	}
	return false
}
func (st *Stack) Show() {
	fmt.Println("top:")
	fmt.Println("\t", st.top)
	fmt.Println("data:")
	fmt.Println("\t", st.data)
}
