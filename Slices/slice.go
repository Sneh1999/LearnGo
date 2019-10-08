package main

import "fmt"

// a slice is a array of strings

func main() {

	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "champ"
	s[1] = "camp"
	s[2] = "hurray"

	fmt.Println("set:", s)
	fmt.Println("get", s[2])

	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := (make([]string, 3))
	copy(c, s) // slices are good becoase we  can easily copy them

	fmt.Println("cpy", c)

	l := s[2:5] // this excludes  5 but includes 2
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl1", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("sl3", t)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
