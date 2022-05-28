package main

import (
	"fmt"
)

func mapMerge(m1 ...map[string]string) map[string]string {
	//m1 就是map的切片
	if len(m1) == 0 {
		return map[string]string{}
	}
	//fmt.Println(m1)
	mRes := make(map[string]string)
	for _, m := range m1 {
		//fmt.Println(m)
		for k, v := range m {
			//fmt.Printf("[k is %v   v = %v]\n",k , v)
			mRes[k] = v
		}
	}
	return mRes
}

func main() {

	m1 := map[string]string{
		"k1":"v1",
		"k2":"v3",
		"k3":"v3",
	}
	m2 := map[string]string{
		"k1":"v11",
		"k2":"v22",
		"k3":"v33",
	}
	m3 := map[string]string{
		"k14":"v11",
		"k24":"v22",
		"k34":"v33",
	}

	//第一种，直接塞入各个参数
	fmt.Println(mapMerge(m1, m2, m3))
	//第二种 slice 传入
	s1 := []map[string]string{m1, m2}
	fmt.Println(mapMerge(s1...))

}
