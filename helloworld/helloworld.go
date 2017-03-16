// helloworld is a package showing one way to write a hello world program in go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"log"
)

var (
	languages       = map[string]Language{}
	idiomaPrincipal *string
)

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

func init() {
	// load languages
	f, err := os.Open("hw.json")
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(f)
	if err := dec.Decode(&languages); err != nil {
		panic(err)
	}

	// parse flags
	idiomaPrincipal = flag.String("lang", "es", "default language as standard abreviation")
	// httpPort := flag.String("port", "8080", "default language as standard abreviation")
	flag.Parse()

	if _, ok := languages[*idiomaPrincipal]; !ok {
		panic("unknown defualt language")
	}
}

func main() {
	switch flag.Arg(0) {
	case "server":
		http.HandleFunc("/", HelloHandler)
		http.ListenAndServe(":8080", nil)
	default:
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
	fmt.Println(l)
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
