package main

import "fmt"

func main() {

	var a = "intial"
	fmt.Println(a)

	var b, c int = 1, 2 //you can do var and then type the type and this would tell go explicitly what the tupe is
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int //go  intialises the int variable to  zero on its own
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
