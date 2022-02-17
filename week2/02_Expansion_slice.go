package main

import "fmt"

func main() {
	s1 := []int{1, 3, 4, 5, 7}

	fmt.Printf("[this slice len is %d , cap is %d ]\n", len(s1), cap(s1))

	s1 = append(s1, 100)
	fmt.Printf("[this slice len is %d , cap is %d ]\n", len(s1), cap(s1))
	//当这个 append 操作完成后，newSlice 拥有一个全新的底层数组，这个数组的容量是原来的两倍
	//append() 会智能地处理底层数组的容量增长。在切片的容量小于 1000 个元素时，总是会成倍地增加容量。一旦元素个数超过 1000，容量的增长因子会设为 1.25，也就是会每次增加 25%的容量(随着语言的演化，这种增长算法可能会有所改变)。

	a1 := []int{10, 11, 12, 13}
	fmt.Printf("len：%d cap：%d\n", len(a1), cap(a1))
	a1 = append(a1, 22)
	fmt.Printf("len：%d cap：%d\n", len(a1), cap(a1))
	a2 := make([]int, 1000)
	a1 = append(a1, a2...)
	fmt.Printf("len：%d cap：%d\n", len(a1), cap(a1))
	a3 := make([]int, 1000)
	a1 = append(a1, a3...)
	fmt.Printf("len：%d cap：%d\n", len(a1), cap(a1))

}
