// helloworld is a package showing one way to write a hello world program in go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var (
	languages       = map[string]Language{}
	idiomaPrincipal *string
)

// init loads languages from json file and sets up flag
// ensure ∃ "hw.json" with map[string]Language.
// allow a "-lang" flag to set default language
func init() {
	f, err := os.Open("languages.json")
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(f)
	if err := dec.Decode(&languages); err != nil {
		panic(err)
	}

	idiomaPrincipal = flag.String("lang", "es", "default language as standard abreviation")
	flag.Parse()
	if _, ok := languages[*idiomaPrincipal]; !ok {
		panic("unknown defualt language")
	}
}

func main() {
	if flag.Arg(0) == "server" {
		http.HandleFunc("/", HelloHandler)
		http.ListenAndServe(":8080", nil)
	}

	lang, ok := languages[*idiomaPrincipal]
	if !ok {
		fmt.Println(UnknownLanguage.Error())
		return
	}
	greeting, err := lang.Greet(flag.Arg(0))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(greeting)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var (
		l, greeting, who string
	)

	fmt.Fprintf(w, "<h1>")
	defer fmt.Fprintf(w, "</h1>\n")

	l = r.URL.Query().Get("lang")
	if l == "" {
		l = r.Header.Get("Accept-Language")
		if len(l) > 2 {
			l = l[:2]
		}
		if l == "" {
			l = *idiomaPrincipal
		}
	}

	lang, ok := languages[l]
	if !ok {
		fmt.Fprintf(w, UnknownLanguage.Error())
		return
	}

	who = r.URL.Path[1:]
	who = strings.Title(who)

	greeting, err := lang.Greet(who)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, greeting)

	if who == "" {
		who = "World"
	}
	log.Printf("%s was greeted with %q in %s", who, greeting, lang.Name)
}

// Language represents a lanugage that can be used to generate greeting
type Language struct {
	Greetings []string `json:"greetings,omitempty"`
	Name      string   `json:"language,omitempty"`
	Fmt       string   `json:"fmt,omitempty"`
}

func (l Language) Greet(who string) (string, error) {
	if len(l.Greetings) == 0 {
		return "", VowOfSilence
	}

	greeting := l.Greetings[rand.Int()%len(l.Greetings)]
	if who == "" {
		return greeting, nil
	}
	greeting = strings.Split(greeting, " ")[0]

	fmtStr := l.Fmt
	if fmtStr == "" {
		fmtStr = "%s %s!"
	}

	return fmt.Sprintf(fmtStr, greeting, who), nil
}

type Sentiment int

const (
	NotImplemented Sentiment = iota
	UnknownLanguage
	VowOfSilence
)

func (s Sentiment) Error() string {
	switch s {
	case UnknownLanguage:
		return "😬"
	case VowOfSilence:
		return "🙊"
	case NotImplemented:
		fallthrough
	default:
		return "😶"
	}
}
func (s Sentiment) HTTPStatus() int {
	switch s {
	case UnknownLanguage:
		return http.StatusNotFound
	case VowOfSilence:
		return http.StatusUnauthorized
	case NotImplemented:
		return http.StatusNotImplemented
	default:
		return http.StatusOK
	}
}
