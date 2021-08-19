package main

import "fmt"
/*
*: 指针标识符
&：取值标识符
 */

// 此方式不能交换值
func swap(a, b int) {
	b, a = a, b
}

//正确的做法
func swap1(a, b *int){
	 *b, *a = *a, *b
}

//上述方式的改良版
func swap2(a, b int)(int, int){
	return b, a
}


func main(){
	a, b := 3, 5
	//swap(a,b)
	//swap1(&a, &b)
	a, b = swap2(a, b)
	fmt.Println(a, b)
}