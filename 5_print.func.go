/*
Go has three functions to output text:
- Print()
- Println()
- Printf()

*/

// The Print() Function

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var i string = "janak"
// 	var j int = 56
// 	fmt.Print(i, "\n")
// 	fmt.Print(j, "\n")
// 	fmt.Print(i, "\n", j)

// 	fmt.Print(i, "hello janak", j)

// }

// The Println() Function

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var i, j string = "Hello", "World"

// 	fmt.Println(i, j)
// }

// The Printf() Function
package main

import (
	"fmt"
)

func main() {
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)
}
