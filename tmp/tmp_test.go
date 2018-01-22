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
	m2 model
}

func (m math2) add(a, b int) int {
	return m.m2.add(a, b)
}
func (m math2) div(a, b int) int {
	return m.m2.div(a, b)
}

func TestInterface(t *testing.T) {
	// var md model
	// md = new(math)
	m := math{}
	m2 := math2{&m}
	fmt.Println(m2.add(1, 2))
}
