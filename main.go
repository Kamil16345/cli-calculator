package main

import (
	"fmt"
	"log"
)

func main() {

	numbers := make([]float64, 0)
	operations := make([]rune, 0)
	var num float64
	var operation rune

	for {
		fmt.Println("Insert number: ")
		_, err := fmt.Scanln(&num)
		if err != nil {
			log.Fatal(err)
			break
		}
		numbers = append(numbers, num)

		fmt.Println("Insert operation: ")
		_, err = fmt.Scanf("%c\n", &operation)
		if err != nil {
			log.Fatal(err)
			break
		}
		if operation == '=' {
			operations = append(operations, operation)
			break
		}
	}
	calculateOperations(numbers, operations)
}

func calculateOperations(numbers []float64, operations []rune) {
	var result float64 = 0
	for i := 0; i < len(operations); i++ {
		result += numbers[i]
		for j:=0 j<len(operations); j++ {

		}
	}
}
