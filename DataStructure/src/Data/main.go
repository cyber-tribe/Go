package main

import (
	"Data/stack"
	"fmt"
)

func main() {
	st := stack.New()
	st.Push(4)
	st.Push(4)
	st.Show()
	st.Pop()
	st.Show()
	st.Push(2)
	st.Show()
	fmt.Println(st.Size())
}
