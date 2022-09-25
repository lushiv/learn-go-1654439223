/*
Go for Loop
Loops are handy if you want to run the same code over and over again, each time with a different value.

Each execution of a loop is called an iteration.

The for loop can take up to three statements:

Syntax
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}
statement1 Initializes the loop counter value.

statement2 Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.

statement3 Increases the loop counter value.

Note: These statements don't need to be present as loops arguments. However, they need to be present in the code in some form.

for Loop Examples
Example 1
This example will print the numbers from 0 to 4:

package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}
*/

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 8; i++ {
		fmt.Println(i)
	}
}
