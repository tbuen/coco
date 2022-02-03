package coco

import "testing"

func TestLexer(t *testing.T) {
	l := Lexer{}
	if l.Result() != 5 {
		t.Error("Invalid Result")
	}
}
