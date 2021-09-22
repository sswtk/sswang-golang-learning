package main

import "fmt"

// 切片

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	s1 := arr[2:]
	fmt.Println("arr[:] = ", arr[:])
	s2 := arr[:]
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	updateSlice(s2)
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:4]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Println("arr =", arr)
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Printf("s1=%v ,len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v ,len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	fmt.Println("Append slice elems ")
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3=, s4=, s5=", s3, s4, s5)
	// s5 no longer view arr (添加元素时，如果超越cap,系统会重新非配更大的底层数组, 并将原来的元素拷贝过去)
	fmt.Println("arr=", arr)

	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}

}
