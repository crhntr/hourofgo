package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

func TestHTTP01(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	(*idiomaPrincipal) = "pl"
	HelloHandler(w, r)
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "<h1>Witaj świecie</h1>\n" {
		t.Fail()
	}
}

func TestHTTP02(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/Christopher", nil)
	(*idiomaPrincipal) = "pl"
	HelloHandler(w, r)
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "<h1>Witaj Christopher!</h1>\n" {
		t.Fail()
	}
}

func TestHTTP03(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/Christopher", nil)
	(*idiomaPrincipal) = "pl"
	r.Header.Set("Accept-Header", "pl")
	HelloHandler(w, r)
	body, _ := ioutil.ReadAll(w.Body)
	if string(body) != "<h1>Witaj Christopher!</h1>\n" {
		t.Fail()
	}
}

func TestHTTP04(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/Christopher", nil)
	(*idiomaPrincipal) = "pl"
	r.URL.Query().Set("Accept-Header", "pl")
	HelloHandler(w, r)
	body, _ := ioutil.ReadAll(w.Body)
	if string(body) != "<h1>Witaj Christopher!</h1>\n" {
		t.Fail()
	}
}
