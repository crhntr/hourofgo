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

	if g, err := es.Greet(""); g != es.Greetings[0] {
		t.Fail()
	} else if err != nil {
		t.Error(err)
	}

	if g, err := es.Greet("Orange"); g != "¡Hola Orange!" {
		t.Fail()
	} else if err != nil {
		t.Error(err)
	}
}

func TestLanguage_Greet02(t *testing.T) {
	es := Language{
		Name:      "klingon",
		Greetings: []string{"qo' vIvan"},
	}

	if g, err := es.Greet(""); g != es.Greetings[0] {
		t.Fail()
	} else if err != nil {
		t.Error(err)
	}

	if g, err := es.Greet("Orange"); g != "qo' Orange!" {
		t.Fail()
	} else if err != nil {
		t.Error(err)
	}
}
