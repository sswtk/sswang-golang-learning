package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 函数外定义变量
//var q = 4
//var u = 8
//var o = "nihao"
////上面的定义方式可以改为
//var (
//	qq = 44
//	uu = 88
//	oo = "nihao"
//)


//定义变量方式1：未设置值
func variableZeroValue() {
	var a int  //整型默认就是0
	var b string // 字符串默认是空字符串
	//fmt.Println(a, b)  // output: 0
	fmt.Printf("%d %q\n", a, b)  //output: 0 ""
}

//定义变量方式2：设置值
func variableInitialValue() {
	// 定义变量写法一
	// var a, b int = 3, 4
	// var s string = "aaa"
	// fmt.Println(a, b, s) // output:3 4 aaa

	// 定义变量写法二
	var a, b, c, d = 1, 2, "sss", true
	var s = "abab"
	fmt.Println(a, b, c, d, s) //output: 1 2 sss true abab

	// 定义变量写法三(此写法只能在函数内使用，函数外必须使用var)
	w, v, t := 7, false, "ttt"
	y := 9
	fmt.Println(w, v, t, y)  // output: 7 false ttt 9
}

//go会自动推断变量类型
func variableTypeDeduction() {
	var a, b, c, d = 1, 2, "sss", true
	var s = "abab"
	fmt.Println(a, b, c, d, s) //output: 1 2 sss true abab
}

// 复数型
func euler() {
	// c := 3 +4i
	// fmt.Println(cmplx.Abs(c))
	// fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
}

//强制类型转换
func triangle() {
	var a, b int = 1, 2
	var c int
	// c = int(math.Sqrt(a*a + b*b))  // 此写法直接编译错误
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量定义
const filename = "abc.txt"
func consts() {
	const a, b = 1, 2
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//常量之枚举类型
func enums() {
	//写法一
	//const (
	//	cpp = 0
	//	java = 1
	//	python = 2
	//	golang = 3
	//)
	// 方式二
	const (
		cpp = iota
		java
		python
		golang
	)
	// 方式三 b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb)
}


func main()  {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	euler()
	triangle()
	consts()
	enums()
}


