package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("Write", i, "as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {

	case time.Saturday, time.Sunday:
		fmt.Println("It is a weekedn")

	default:
		fmt.Println("Its is a weekday")
	}

	t := time.Now()

	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")

	default:
		fmt.Println("Its after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")

		default:
			fmt.Println("Dont know its type", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
