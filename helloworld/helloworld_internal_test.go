package main

import (
	"testing"
)

func TestLanguage_Greet01(t *testing.T) {
	es := Language{
		Name:      "spanish",
		Greetings: []string{"Hola Mundo"},
		Fmt:       "¡%s %s!",
	}

	if es.Greet("") != es.Greetings[0] {
		t.Fail()
	}

	if es.Greet("Orange") != "¡Hola Orange!" {
		t.Fail()
	}
}

func TestLanguage_Greet02(t *testing.T) {
	es := Language{
		Name:      "klingon",
		Greetings: []string{"qo' vIvan"},
	}

	if es.Greet("") != es.Greetings[0] {
		t.Fail()
	}

	if es.Greet("Orange") != "qo' Orange!" {
		t.Fail()
	}
}
