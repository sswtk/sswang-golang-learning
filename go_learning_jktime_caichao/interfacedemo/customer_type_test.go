package interfacedemo

import (
	"fmt"
	"testing"
	"time"
)

//func timeSpent(inner func(op int) int) func(op int) int{
//	return func(n int) int{
//		start := time.Now()
//		ret := inner(n)
//		fmt.Println("time spent:", time.Since(start).Seconds())
//		return ret
//	}
//}

//上面注释的函数可以改写为：
type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv{
	return func(n int) int{
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int{
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}