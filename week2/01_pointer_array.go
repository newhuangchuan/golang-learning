package main

import (
	"fmt"
	"unsafe"
)

func bigArrPointer(arr *[1e6]int64) {
	fmt.Printf("[array pointer copy : size : %d byte]\n", unsafe.Sizeof(arr))
}

func bigArr(arr [1e6]int64) {
	fmt.Printf("[array copy : size : %d byte]\n", unsafe.Sizeof(arr))
}

func main() {

	var arr [1e6]int64
	bigArr(arr)
	bigArrPointer(&arr)
}
