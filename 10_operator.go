// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a = 15 + 25
// 	var b = 15 - 2
// 	var c = 15 * 2
// 	var d = 15 / 2
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

package main

import (
	"fmt"
)

func main() {
	var (
		sum1 = 100 + 50    // 150 (100 + 50)
		sum2 = sum1 + 250  // 400 (150 + 250)
		sum3 = sum2 + sum2 // 800 (400 + 400)
	)
	fmt.Println(sum3)
}
