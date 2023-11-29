package main

type converter struct {
	expressions []expression
	i           int
}
type nodeKind string

const (
	addNode  nodeKind = "+"
	subNode  nodeKind = "-"
	mulNode  nodeKind = "*"
	divNode  nodeKind = "/"
	funcNode nodeKind = "func"
	numNode  nodeKind = "num"
)

type node struct {
	kind  nodeKind
	left  *node
	right *node

	funcName string
	args     []*node

	val float64
}

func newConverter(expressions []expression) *converter {
	return &converter{expressions: expressions, i: 0}
}
func (c *converter) convert() (*node, error) {
	return c.add()
}
func (c *converter) add() (*node, error) {
	n, err := c.mul()
}
func (c *converter) mul() (*node, error) {
	if c.consume("+") {

	} else if c.consume("-") {

	}
	return c.primary()
}
func (c *converter) primary() (*node, error) {
	if c.consume("(") {
		n, err := c.add()
		if err != nil {
			return nil, err
		}
		c.consume(")")
	}
	if c.expressions[c.i].kind == operatorToken {

	}
	return
}
func (c *converter) consume(s string) bool {
	t := c.expressions[c.i]
	if t.kind != operatorToken || t.str != s {
		return false
	}
	c.i++
	return true
}
