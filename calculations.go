package main

import (
	"fmt"
	"regexp"
	"unicode"
)

type tokenKind string

const (
	numberToken tokenKind = "number"
	identToken  tokenKind = "ident"
)

type expression struct {
	kind tokenKind
	val  float64
	str  string
}

func Calculate(equation string) (int, error) {
	chars := []rune(equation)
	i := 0
	n := len(chars)
	equations := []expression{}
	for i < n {
		char := chars[i]
		if unicode.IsSpace(char) {
			i++
			continue
		}
		if isChar(char) {
			equations = append(equations, expression{kind: numberToken, str: string(char)})
		}
		if isNumber(char) {

		}
	}
	return 1
}
func isChar(char rune) bool {
	regex := regexp.MustCompile(`^[A-Za-z]$`)
	fmt.Println("Regex: ", regex)
	return regex.MatchString(string(char))
}
func isNumber(char rune) bool {

}
