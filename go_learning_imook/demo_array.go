package main

import "fmt"

func printArray(arr [5]int){
	for i,v := range arr{
		fmt.Println(i, v)
	}
	// 验证数组是值类型
	arr[0] = 100
}


func main(){
	//定义数组
	var arr1 [5]int
	arr2 := [4]int{1, 2, 4, 5}
	arr3 := [...]int{1, 2, 4, 5, 6}
	var grid [4][5]int // 二维数组4行5列
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	//遍历数组方式一
	//for i := 0; i < len(arr3); i++{
	//	fmt.Println(arr3[i])
	//}

	//遍历数组方式二 range
	//for i := range arr3{
	//	fmt.Println(arr3[i])
	//}

	//遍历数组并获取下标和值
	//for i, v := range arr3{
	//	fmt.Println(i, v)
	//}

	//遍历数组只获取值
	//for _, v := range arr3{
	//	fmt.Println(v)
	//}

	printArray(arr1)
	printArray(arr3)
	fmt.Println(arr1, arr3)
}



