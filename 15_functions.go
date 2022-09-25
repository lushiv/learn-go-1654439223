//Function Return Example
// package main

// import (
// 	"fmt"
// )

// func myFunction(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(myFunction(1, 2))
// }

// named return

package main

import (
	"fmt"
)

func myFunction(x int, y int) (result int) {
	result = x + y
	return
}

func main() {
	fmt.Println(myFunction(1, 2))
}
