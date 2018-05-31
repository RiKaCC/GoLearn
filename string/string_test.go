package string

import (
	"testing"
)

func TestStringJoin(t *testing.T) {
	s1 := "hello"
	s2 := "world"
	s3 := "every"

	s := StringJoin(s1, s2, s3).BetweenWith("")
	if s != "helloworldevery" {
		t.Errorf("StringJoin")
	}

	ss1 := StringJoin(s1, s2, s3).BetweenWith(" ")
	if ss1 != "hello world every" {
		t.Errorf("StringJoin")
	}
}

func TestStringJoinMore(t *testing.T) {
	str := []string{"hello", "world", "every"}

	s := StringJoinMore(str).BetweenWith("")
	if s != "helloworldevery" {
		t.Errorf("StringJoinMore")
	}

	s1 := StringJoinMore(str).BetweenWith(" ")
	if s1 != "hello world every" {
		t.Errorf("StringJoinMore")
	}
}

func TestSubString(t *testing.T) {
	str := "Monday"
	ret, _ := SubString(str).Start(0).End(4)
	if ret != "Monda" {
		t.Errorf("SubString")
	}

	ret1, _ := SubString(str).OnlyStart(2)
	if ret1 != "nday" {
		t.Errorf("SubString")
	}

	ret2, _ := SubString(str).OnlyEnd(4)
	if ret2 != "Monda" {
		t.Errorf("SubString")
	}
}
