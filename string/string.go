package string

import (
	"bytes"
	"errors"
)

// 字符串反转
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
	s          string
	start      int
	end        int
}

// 字符串拼接
func StringJoin(strs ...string) *String {

	s := new(String)
	s.JoinArr = strs

	return s
}

func StringJoinMore(str []string) *String {

	s := new(String)
	s.JoinArr = str

	return s
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

// 获取子串
func SubString(str string) *String {
	st := []rune(str)

	s := &String{
		s: string(st),
	}
	return s
}

func (s *String) Start(n int) *String {
	if n < 0 {
		n = 0
	}
	s.start = n

	return s
}

func (s *String) OnlyStart(n int) (string, error) {
	if n < 0 {
		n = 0
	}

	len := len(s.s)
	if n >= len {
		return "", errors.New("start must less than length of string")
	}

	s.start = n

	return s.s[s.start:len], nil
}

func (s *String) End(n int) (string, error) {
	len := len(s.s)
	if n > len {
		n = len
	}

	// Start的时候不用去判断start是不是大于字符串 的长度，在这里一次就能判断了
	if s.start >= n {
		return "", errors.New("start larger than end")
	}
	s.end = n

	ret := s.s[s.start : s.end+1]
	return ret, nil
}
