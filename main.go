package main

import (
	"fmt"
)

func main() {
	var calculation string
	for {
		fmt.Scanln(&calculation)
		if calculation == "quit" || calculation == "end" || calculation == "exit" {
			break
		}
		val, err := Calculate(calculation)
		c := newConverter(val)
		n, err := c.convert()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(val)
		}
		result, err := calculate(n)
		fmt.Printf("Result: %f\n", result)
	}

}
