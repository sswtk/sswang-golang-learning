package fibdemo

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1
	a := 1
	b := 1
	//fmt.Print(a)
	t.Log(a)
	for i := 0; i< 5; i++ {
		//fmt.Println(" ", b)
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// 方式一
	// tmp := a
	// a = b
	// b = tmp
	// 方式二
	a, b = b, a
	t.Log(a, b)
}