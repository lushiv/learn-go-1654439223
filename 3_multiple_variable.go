// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a, b, c, d int = 1, 2, 3, 4
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a, b = 6, "Hello"
// 	c, d := 7, "World!"

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

// Go Variable Declaration in a Block

package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b int    = 2
		c string = "janak"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

/*

Go Variable Naming Rules
A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

# Go variable naming rules:

- A variable name must start with a letter or an underscore character (_)
- A variable name cannot start with a digit
- A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
- Variable names are case-sensitive (age, Age and AGE are three different variables)
- There is no limit on the length of the variable name
- A variable name cannot contain spaces
- The variable name cannot be any Go keywords

## Multi-Word Variable Names

## Camel Case
- myVariableName = "John"

## Pascal Case

MyVariableName = "John"
## Snake Case

my_variable_name = "John"

*/
