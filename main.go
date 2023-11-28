package main

import (
	"fmt"
)

func main() {
	var calculation string
	fmt.Scanln(&calculation)
	val, err := Calculate(calculation)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}
