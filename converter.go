package main

import "fmt"

type converter struct {
	expressions []expression
	i           int
}
type nodeKind string

const (
	addNode nodeKind = "+"
	subNode nodeKind = "-"
	mulNode nodeKind = "*"
	divNode nodeKind = "/"
	numNode nodeKind = "num"
)

type node struct {
	kind  nodeKind
	left  *node
	right *node

	funcName string
	args     []*node

	val float64
}

// Converter
func newConverter(expressions []expression) *converter {
	return &converter{expressions: expressions, i: 0}
}

func (c *converter) convert() (*node, error) {
	return c.add()
}
func (c *converter) add() (*node, error) {
	n, err := c.mul()
	if err != nil {
		return nil, err
	}
	for {
		if c.consume("+") {
			n, err = c.insert(n, c.mul, addNode)
			if err != nil {
				return nil, err
			}
		} else if c.consume("-") {
			n, err = c.insert(n, c.mul, subNode)
			if err != nil {
				return nil, err
			}
		} else {
			return n, nil
		}
	}
}
func (c *converter) mul() (*node, error) {
	n, err := c.primary()
	if err != nil {
		return nil, err
	}
	for {
		if c.consume("*") {
			n, err = c.insert(n, c.unary, mulNode)
			if err != nil {
				return nil, err
			}
		} else if c.consume("/") {
			n, err = c.insert(n, c.unary, divNode)
			if err != nil {
				return nil, err
			}
		} else {
			return n, nil
		}
	}
}
func (c *converter) primary() (*node, error) {
	if c.consume("(") {
		n, err := c.add()
		if err != nil {
			return nil, err
		}
		c.consume(")")
		return n, nil
	}
	return c.numberNode()
}

func (c *converter) consume(s string) bool {
	t := c.expressions[c.i]
	if t.kind != operatorToken || t.str != s {
		return false
	}
	c.i++
	return true
}

func (c *converter) insert(n *node, f func() (*node, error), kind nodeKind) (*node, error) {
	left := n
	right, err := f()
	if err != nil {
		return n, err
	}
	return &node{kind: kind, left: left, right: right}, err
}
func (c *converter) unary() (*node, error) {
	if c.consume("+") {
		return c.primary()
	} else if c.consume("-") {
		return c.insert(&node{kind: numNode, val: 0.0}, c.primary, subNode)
	}
	return c.primary()
}
func (c *converter) numberNode() (*node, error) {
	t := c.expressions[c.i]
	if t.kind != numberToken {
		return nil, fmt.Errorf("expected a number: %s", t.str)
	}
	c.i++
	return &node{kind: numNode, val: t.val}, nil
}
