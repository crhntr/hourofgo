package main

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	tests := []struct {
		to, lang, ret string
	}{
		{"", "", "Hello World!"},
		{"Maria", "fr", "Bonjour Maria!"},
		{"Lex", "es", "¡Hola Lex!"},
	}

	for i, tst := range tests {
		ret := sayHello(tst.to, tst.lang)
		if ret != tst.ret {
			t.Errorf("say hello test %d failed sayHello(%q, %q) should return %q but got %q", i, tst.to, tst.lang, tst.ret, ret)
		}
	}
}
