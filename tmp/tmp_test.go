package tmp_test

import (
	"fmt"
	"testing"
)

type model interface {
	add(a, b int) int
	div(a, b int) int
}

type math struct {
	res int
}

func (m *math) add(a, b int) int {
	m.res = a + b
	return a + b
}
func (m math) div(a, b int) int {
	return a / b
}

type math2 struct {
	next model
}

func (m math2) add(a, b int) int {
	println("math2")
	return m.next.add(a, b)
}
func (m math2) div(a, b int) int {
	println("math2")
	return m.next.div(a, b)
}

type math3 struct {
	next model
}

func (m math3) add(a, b int) int {
	println("math3")
	return m.next.add(a, b)
}
func (m math3) div(a, b int) int {
	println("math3")
	return m.next.div(a, b)
}

func TestInterface(t *testing.T) {
	var m model
	m = &math{}
	m = math2{m}
	m = math3{m}
	fmt.Println(m.add(1, 2))
}
