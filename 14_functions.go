package main

import (
	"fmt"
)

func internalHelper(name string, age int64) {
	fmt.Println(name, age)
}

func main() {
	internalHelper("Janak", 89) // call the function
}



