package main 

import "fmt"
// -------------------------------

// 语法：
//  ①：使用var 关键字
	// 	var a, b, c bool
	// 	var s1, s2, string = "hello", "world"
// ②：使用:= 定义变量(只能在函数内部使用)
	// 	a, b, i, s1, s2 := true, false, 3, "hello", "world" 
// 变量可以放在函数内，或者直接放在包内
// 使用var关键字集中定义变量

// -------------------------------

// var aa = 3
// var kk = "8"
// var dd = true
// 上面的多个变量的定义可以换成下列方式
var (
	aa = 3 
	kk = "sdfwefw"
	kk4 = "6"
	dd = true
)

func variableZeroValue() {
	var a int 
	var b string 
	fmt.Printf("%d %q\n", a, b)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "ABC"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, ""
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"  // 可以用： 来定义变量，和var的效果是一样的
	b = 9
	fmt.Println(a, b, c, s)
}
// func traingle() {
// 	var a, b int = 3, 4 
// 	var c int 
// 	c = int(math.Sqrt(float64(a*a + b*b)))
// 	fmt.Println(c)
// }



func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, kk, dd, kk4)
	// traingle()
}