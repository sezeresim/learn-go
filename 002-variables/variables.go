package main

import "fmt"

func main() {
	var name string = "Sezer"
	fmt.Println("Hello", name)

	var total int = 4 + 2
	fmt.Println("4 + 2 = ", total)

	var isChecked bool = true

	if isChecked {
		fmt.Println("ISCHECKED")
	}

	var n64 int64 = 999999999999999999

	fmt.Println(n64)

	var n16 int16 = 9999

	fmt.Println(n16)
	n16 = n16 + 2
	fmt.Println(n16)

	/* Constant */
	const n17 uint = 10
	// n17 = n17 + 2
	fmt.Println(n17)

	/* short syntax for variable declaration */
	test2 := "tdsaest"
	fmt.Println(test2)

}
