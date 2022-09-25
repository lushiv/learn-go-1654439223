// package main

// import (
// 	"fmt"
// )

// func main() {
// 	myslc_1 := []int{}
// 	fmt.Println(len(myslc_1))
// 	fmt.Println(cap(myslc_1))
// 	fmt.Println(myslc_1)

// 	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
// 	fmt.Println(len(myslice2))
// 	fmt.Println(cap(myslice2))
// 	fmt.Println(myslice2)

// }

/// Create a Slice From an Array

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr1 := [6]int{10, 11, 12, 13, 14, 15}
// 	myslice := arr1[2:4]

// 	fmt.Printf("myslice = %v\n", myslice)
// 	fmt.Printf("length = %d\n", len(myslice))
// 	fmt.Printf("capacity = %d\n", cap(myslice))
// }

//Create a Slice With The make() Function
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	myslice1 := make([]int, 5, 10)
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))

// 	// with omitted capacity
// 	myslice2 := make([]int, 5)
// 	fmt.Printf("myslice2 = %v\n", myslice2)
// 	fmt.Printf("length = %d\n", len(myslice2))
// 	fmt.Printf("capacity = %d\n", cap(myslice2))
// }

//Access Elements of a Slice
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	prices := []int{10, 20, 30}

// 	fmt.Println(prices[0])
// 	fmt.Println(prices[2])
// }

//Append Elements To a Slice
package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}
