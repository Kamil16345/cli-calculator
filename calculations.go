package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

type tokenKind string

const (
	numberToken   tokenKind = "number"
	operatorToken tokenKind = "operator"
	endToken      tokenKind = "eos"
)

type expression struct {
	kind tokenKind
	val  float64
	str  string
}

const operators = "+-*/(),"

func Calculate(equation string) ([]expression, error) {
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
		if isOperator(char) {
			i++
			equations = append(equations, expression{kind: operatorToken, str: string(char)})
		}
		if isLetter(char) {
			fmt.Println("In calculation there is a letter")
			break
		}
		if isNumber(char) {
			startIndex := i
			i++
			for i < n && isNumber(chars[i]) {
				i++
			}
			equations = append(equations, expression{kind: numberToken, str: string(chars[startIndex:i])})
			continue
		}
		if val, err := numberPrefix(chars, &i, n); err == nil {
			equations = append(equations, expression{kind: numberToken, val: val})
			continue
		}

	}
	equations = append(equations, expression{kind: endToken})
	fmt.Println("Equations: ", equations)
	return equations, nil
}
func isOperator(char rune) bool {
	for _, op := range operators {
		if char == op {
			return true
		}
	}
	return false
}
func isLetter(char rune) bool {
	regex := regexp.MustCompile(`^[A-Za-z]$`)
	if regex.MatchString(string(char)) {
		return true
	}
	return false
}
func isNumber(char rune) bool {
	regex := regexp.MustCompile(`\d`)
	if regex.MatchString(string(char)) {
		return true
	}
	return false
}
func numberPrefix(chars []rune, i *int, n int) (float64, error) {
	val := 0.0
	len := 0
	for *i < n {
		currentVal, err := strconv.ParseFloat(string(chars[*i-len:*i+1]), 64)
		if err != nil {
			break
		}
		val = currentVal
		len++
		*i++
	}
	if len > 0 {
		return val, nil
	}
	return 0, errors.New("Lack of a number")
}
