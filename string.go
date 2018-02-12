package main

import (
	"bytes"
	"fmt"
)

//字符串反转
func ReverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

type String struct {
	JoinString string
	JoinArr    []string
}

func (s *String) BetweenWith(joinstring string) string {

	var buf bytes.Buffer
	s.JoinString = joinstring

	len := len(s.JoinArr)

	for i, str := range s.JoinArr {
		buf.WriteString(str)
		if i != (len - 1) {
			buf.WriteString(s.JoinString)
		}
	}

	return buf.String()
}

//字符串拼接
func StringJoin(strs ...string) *String {

	s := new(String)
	s.JoinArr = strs

	return s
}

func StringJoinMore(str interface{}) string {
	return ""
}

func main() {
	str := "hello"
	fmt.Println(ReverseString(str))
	fmt.Println(Reverse(str))

	s1 := "hey"
	s2 := "you"

	s := StringJoin(s1, s2).BetweenWith("+")
	fmt.Println(s)
}
