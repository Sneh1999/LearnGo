package main

import (
	"errors"
	"fmt"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	width, height int
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func zeroptr(iptr *int) {
	*iptr = 0
}

type people struct {
	name string
	age  int
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cant work with 42")
	}
	return arg + 3, nil
}

// func (r rect) perim() float64 {
//     return 2*r.width + 2*r.height
// }
func NewPerson(name string) *people {

	p := people{name: name}
	p.age = 42
	return &p
}

func (r *rect) area() int {
	return r.width * r.height
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// it basically makes all the arguments as inputs array and then we can apply the range function
func sum(nums ...int) {
	fmt.Println(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println((total))
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3", res)
	_, c := vals()
	fmt.Println(c)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// zeroptr(&i)
	// fmt.Println("zeroptr:", i)
	fmt.Println(NewPerson("Jon"))
}
