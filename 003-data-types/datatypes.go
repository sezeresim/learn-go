package main

import "fmt"

func main() {

	var name = "Sezer"
	var age int16 = -256
	var isMarried bool = true
	var weight float32 = 72.5
	const constty = "SEZER"
	isMarried = false

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)

	fmt.Println("---------------")

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
	fmt.Printf("%T\n", weight)
	fmt.Printf("%T\n", constty)

	// Using 8-bit unsigned int
	var X uint8 = 225
	fmt.Println(X, X-3)

	// Using 16-bit signed int
	var Y int16 = 32767
	fmt.Println(Y+2, Y-2)

	var a complex128 = complex(6, 2)
	var b complex64 = complex(9, 2)
	fmt.Println(a)
	fmt.Println(b)

	// Display the type
	fmt.Printf("The type of a is %T and "+
		"the type of b is %T", a, b)

}
