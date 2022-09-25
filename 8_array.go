package main

import "fmt"

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	arr2 := [3]int{1, 2, 3}
	var arr3 = [2]string{"jaank", "karan"}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	// Access Elements of an Array
	fmt.Println(arr3[0])
	fmt.Println(arr2[2])

	// Change Elements of an Array
	arr3[0] = "jenith"
	fmt.Println(arr3)

	// Array Initialization
	arr4 := [5]int{}              //not initialized
	arr5 := [5]int{1, 2}          //partially initialized
	arr6 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)

	// Initialize Only Specific Elements
	arr7 := [5]int{1: 10, 2: 40}
	fmt.Println(arr7)
	fmt.Println(len(arr7))

}
