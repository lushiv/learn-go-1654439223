/*
In Go, there are different types of variables, for example:

- int- stores integers (whole numbers), such as 123 or -123
- float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
- string - stores text, such as "Hello World". String values are surrounded by double quotes
- bool- stores values with two states: true or false

*/

//Declaring (Creating) Variables ====> 1. With the var keyword
// var variablename type = value

// Variable Declaration With Initial Value
// package main

// import "fmt"

// func main() {
// 	var name1 string = "Janak Raikhola"
// 	var name2 = "Jenith"
// 	a := 1

// 	fmt.Println(name1)
// 	fmt.Println(name2)
// 	fmt.Println(a)
// }

// 2. Variable Declaration Without Initial Value

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a string
// 	var b int
// 	var c bool
// 	var d float32

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)

// }

// 3. Value Assignment After Declaration

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var name string
// 	name = "janak"
// 	fmt.Println(name)
// }

//Difference Between var and :=

package main
import ("fmt")

var a int
var b int = 2
var c = 3

func main() {
	a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}