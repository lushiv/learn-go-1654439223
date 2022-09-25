// package main

// import (
// 	"fmt"
// )

// const _USER_ID_ = "4bb75e74-3ccc-11ed-bc19-6b9f3142ec0a"

// func main() {
// 	fmt.Println(_USER_ID_)
// }

// Multiple Constants Declaration

package main

import (
	"fmt"
)

const (
	username string = "janak raikhola"
	password        = "Test@1234"
	Id       int    = 12123344
	status          = true
)

func main() {
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(Id)
	fmt.Println(status)
}
