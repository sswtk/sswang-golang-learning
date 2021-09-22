package main

import "fmt"

const (
	Frist int = iota
	Second
	Third
	Fourth
)

func main() {
	// 申明和初始化数组
	var a1 [3]int
	fmt.Println(a1) // [0 0 0]
	var a2 = [3]int{1, 2, 3}
	fmt.Println(a2) // [1 2 3]
	var a3 = [3]string{"a", "b", "c"}
	fmt.Println(a3) //[a b c]
	var a4 = [5]string{"r", "t", "s"}
	fmt.Println(a4) //[r t s  ]  s后面是两个空字符串
	a5 := [5]int{1, 3, 5}
	fmt.Println(a5) //[1 3 5 0 0]
	a6 := [3]int{7}
	fmt.Println(a6) //[7 0 0]
	a7 := [...]int{11, 22, 33}
	fmt.Println(a7) //[11 22 33]  [...]表示省略数组长度，编辑器会自动计算长度
	a8 := [4]string{Frist: "1", Second: "2", Third: "3", Fourth: "4"}
	fmt.Println(a8) //[1 2 3 4]

	//数组的操作---比较
	//aa := [3]int{5, 78, 8}
	//var bb [5]int
	//bb = aa //cannot use aa (type [3]int) as type [5]int in assignment

	//数组的操作---取值
	cc := [...]string{"anhui", "jiangsu", "shanghai", "zhejiang"}
	dd := cc[1]
	fmt.Println(dd) //jiangsu

	//数组的操作---修改
	cc[1] = "beijing"
	fmt.Println(cc) //[anhui beijing shanghai zhejiang]

	//数组的操作---求长度
	ee := [...]float64{33.3, 45.2, 43, 89.1}
	fmt.Println("length of ee is", len(ee)) //length of ee is 4

	//数组的操作---迭代、遍历
	for i := 0; i < len(ee); i++ {
		fmt.Printf("%d th element of ee is %.2f\n", i, ee[i])
		//0 th element of ee is 33.30
		//1 th element of ee is 45.20
		//2 th element of ee is 43.00
		//3 th element of ee is 89.10
	}

	//数组的操作---使用range遍历
	sum := float64(0)
	for ii, v := range ee {
		fmt.Printf("%d the element of ee is %.2f\n", ii, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of ee", sum) //sum of all elements of ee 210.6

	//数组的操作---使用range遍历,且忽略索引只需要获取值
	for _, v := range ee {
		fmt.Println(v)
		//33.3
		//45.2
		//43
		//89.1
	}
}
