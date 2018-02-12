package main

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
