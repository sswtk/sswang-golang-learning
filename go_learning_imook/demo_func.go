package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 函数之返回一个值
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)
	}
}

//改良版
func eval1(a, b int, op string) (int, error){
	switch op{
	case "+":
		return a + b , nil
	case "-":
		return a - b , nil
	case "*":
		return a * b , nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// 函数之返回多值情况
func div(a, b int) (int, int) {
	return a / b , a % b
}
// 函数之给返回值取名
func div1(a, b int) (q, r int ) {
	//return a / b , a % b // 多值情况建议使用此方式返回
	q = a / b
	r = a % b
	return
}

//函数里面嵌套函数
func apply(op func(int, int) int, a, b int) int{
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args" + "(%d, %d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数列表
func sum(numbers ...int) int{
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
func main(){
	fmt.Println(sum(1, 2, 3, 4, 5))
	//fmt.Println(apply(pow, 3, 4))
	////不定义pow函数，使用匿名函数的方式
	//fmt.Println(apply(
	//	func(a int, b int) int{
	//		return int(math.Pow(float64(a), float64(b)))
	//	},3, 4))
	//if ret, err := eval1(3, 4, "x"); err != nil{
	//	fmt.Println("Error:", err)
	//} else{
	//	fmt.Println(ret)
	//}
	//q, r := div1(13, 3)
	//fmt.Println(q, r)
	//fmt.Println(div(13, 3))
	//fmt.Println(eval(3,6, "+"))
	//fmt.Println(eval(3,6, "*"))
	//fmt.Println(eval(3,6, "%"))
}