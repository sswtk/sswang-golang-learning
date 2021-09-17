package main

import "fmt"

func main() {
	//bool
	var w bool // 申声明变量，如果没有初始化，默认为false
	var ww bool = true
	var www = true //自动推断类型
	fmt.Println(w, ww, www)
	a := true
	b := false
	fmt.Println(a, b)
	c := a && b
	fmt.Println(c)
	d := a || b
	fmt.Println(d)

	//string
	var str1 string // 声明字符串类型
	str1 = "aaa"
	fmt.Println("str1=", str1)
	first := "Ronin"
	last := "ssWang"
	name := first + last
	fmt.Println("My name is", name)
	fmt.Printf("name 的类型是 %T\n", name)
	// len()，内建函数，可以返回字符串有多少个字符
	fmt.Println("len(name)=", len(name)) //len(name)=11
	//如果只想操作某个字符，可用如下方式
	fmt.Printf("name[0] = %c, name[1] = %c", name[0], name[1])

	//byte
	var n byte
	n = 87
	fmt.Printf("%c, %d\n", n, n) //%c:以字符的形式打印，%d:以整型方式打印
	var ch byte
	ch = 'a' // 单引号表示字符， 双引号表示字符串
	fmt.Printf("%c, %d\n", ch, ch)

	//rune
	str := "nihao你好"
	fmt.Println(len(str))
	k := []byte(str)
	for _, v := range k {
		fmt.Print(string(v)) //nihaoä½ å¥½
	}
	fmt.Println("")
	r := []rune(str)
	for _, v := range r {
		fmt.Print(string(v)) // nihao你好
	}
	fmt.Println("")
	r[0] = 'N'
	r[5] = '您'
	fmt.Println(string(r)) // Nihao您好

	// float
	var y float32
	y = 3.3
	yy := 6.555 // 自动推导类型，默认是float64,要比float32更精确
	fmt.Println("y=", y)
	fmt.Printf("yy的类型为 %T", yy) //yy的类型为 float64

	//complex
	var t complex128
	t = 2.1 + 3.14i
	fmt.Println("t=", t)
	t1 := 3.3 + 5.5i
	fmt.Printf("t1的类型是%T\n", t1)
	// 通过内建函数，取出实部和虚部
	fmt.Println("real(t1)=", real(t1), "imag(t1)=", imag(t1)) //real(t1)= 3.3 imag(t1)= 5.5

}
