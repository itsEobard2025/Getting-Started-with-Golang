package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	ptr := new(int)
	*ptr = 3
	fmt.Println(*ptr)

	var z int32 = 1
	var y int64 = 2
	z = int32(y) // type conversion from int64 to int32
	fmt.Println(z)

	var x float32 = 2.345e2
	fmt.Println(x)

	var w complex128 = complex(2, 3)
	fmt.Println(w)

	// count := 10
	// for i := 0; i < count; i++ {
	// 	fmt.Println(i)
	// }

	switch x {
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
	default:
		fmt.Println("Default Case")
	}

	var a int
	var b *int
	c := 3
	b = &c
	a = &b

}
