package main

import (
	"fmt"
)

func main() {
	q := [...]int{1,2,3}
	sli := q[0:1]
	for k,v := range sli {
		fmt.Println(k,v)
	}

	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))

	var sss []int
	sss = append(sss, 1) // 追加1个元素
	sss = append(sss, 1, 2, 3) // 追加多个元素, 手写解包方式
	sss = append(sss, []int{1,2,3}...) // 追加一个切片, 切片需要解包

	var numbers []int
	for i := 0; i < 10; i++ {
	    numbers = append(numbers, i)
	    fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
		
}