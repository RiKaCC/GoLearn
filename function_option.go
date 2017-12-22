package main

import (
	"fmt"
)

type option struct {
	week  int
	str   string
	score int
}

func NewOption(opt ...TestOption) *option {
	in := new(option)
	for _, optemp := range opt {
		optemp(in)
	}

	return in
}

type TestOption func(*option)

func Week(a int) TestOption {
	return func(o *option) {
		o.week = a
	}
}

func Str(s string) TestOption {
	return func(o *option) {
		o.str = s
	}
}

func Score(s int) TestOption {
	return func(o *option) {
		o.score = s
	}
}

func main() {
	op1 := Week(1)
	op2 := Str("Monday")
	op3 := Score(100)

	op := NewOption(op1, op2, op3)
	fmt.Println(op)
}
