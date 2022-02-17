package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("[s1 slice len == %d , cap == %d ]\n", len(s1), cap(s1))

	s2 := s1[2:6]
	fmt.Printf("[s2][从索引为2 就是3号往后切4个元素][值:%v][新切片的长度=%d 容量=%d]\n", s2, len(s2), cap(s2))
	s3 := s1[5:]
	fmt.Printf("[s3][从索引为5 就是6号 切到最后][值:%v][新切片的长度=%d 容量=%d]\n", s3, len(s3), cap(s3))
	s4 := s1[:4]
	fmt.Printf("[s4][从开头切4个元素][值:%v][新切片的长度=%d 容量=%d]\n", s4, len(s4), cap(s4))
	s5 := s1[:]
	fmt.Printf("[s5][复制整个切片][值:%v][新切片的长度=%d 容量=%d]\n", s5, len(s5), cap(s5))

	s6 := s1[2:6:8]
	fmt.Printf("[s6][从索引为2 就是3号往后切4个元素][值:%v][新切片的长度=%d 容量=%d]\n", s6, len(s6), cap(s6))

	s7 := s1[2:6:9]
	fmt.Printf("[s7][从索引为2 就是3号往后切4个元素][值:%v][新切片的长度=%d 容量=%d]\n", s7, len(s7), cap(s7))

	fmt.Printf("[改变某一个切片的元素。看看其他切片会受影响吗]\n")
	s6[1] = 80
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Println(s7)
}
