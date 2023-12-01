package main

import (
	"fmt"
)

func main() {
	var calculation string
	fmt.Scanln(&calculation)
	val, err := Calculate(calculation)
	c := newConverter(val)
	n, err := c.convert()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
	result, err := calculate(n)
	fmt.Printf("Result: %f", result)
}
