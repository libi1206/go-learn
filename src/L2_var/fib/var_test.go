package fib

import (
	"fmt"
	"testing"
)

const (
	One int = 1 + iota
	Two
	Three
)

func TestFibList(t *testing.T)  {
	var (
		a int = 1
		b int = 1
	)
	fmt.Println(a)
	for i:=0;i<5;i++ {
		fmt.Print(" ",b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func TestConst(t *testing.T)  {
	t.Log(One)
	t.Log(Two)
	t.Log(Three)
}