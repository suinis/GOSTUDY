package main

import (
	"fmt"
	"unsafe"
)

const (
	A = 1
	B = 2
	C = 3
)

const (
	AA = iota // iota 默认值(0)
	BB
	CC
)

const (
	AAA = "abc"
	BBB = len(AAA)
	CCC = unsafe.Sizeof(AAA)
)

const (
	first, second = iota + 1, iota + 2
	third, forth
	fifth, sixth
)

func main() {
	// const 
	const len int = 10
	fmt.Println("len = ", len)

	// len = 100 // const variables

	fmt.Println("A = ", A , ", B = ", B, ", C = ", C)

	fmt.Println("AA = ", AA , ", BB = ", BB, ", CC = ", CC) // 0 1 2

	fmt.Println("AAA = ", AAA, ", BBB = ", BBB, ", CCC = ", CCC) // 0 1 2

	fmt.Println("first = ", first, ", second = ", second, 
	", third = ", third, ", forth = ", forth, 
	", fifth = ", fifth, ", sixth = ", sixth)
}