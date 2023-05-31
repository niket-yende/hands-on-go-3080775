// packages/basics/main.go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by 0")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Hello Niket!")

	// fmt.Printf("Today is %s", time.Now().Weekday())

	// Invoke divide function
	fmt.Println(divide(10, 2))

	// Invide divide function with 0 as denominator to return error
	fmt.Println(divide(5, 0))
}
