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
)

var (
	languages       = map[string]Language{}
	idiomaPrincipal *string
)

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
	idiomaPrincipal := flag.String("lang", "es", "default language as standard abreviation")
	// httpPort := flag.String("port", "8080", "default language as standard abreviation")
	flag.Parse()

	if _, ok := languages[*idiomaPrincipal]; !ok {
		panic("unknown defualt language")
	}
}

func main() {
	switch flag.Arg(1) {
	case "server":
		http.HandleFunc("/", HelloHandler)
		http.ListenAndServe(":8080", nil)
	default:
		lang, ok := languages[*idiomaPrincipal]
		if !ok {
			fmt.Println("😕")
			return
		}
		fmt.Println(lang.Greet(flag.Arg(1)))
	}

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	
}

type Language struct {
	Greetings []string `json:"greetings,omitempty"`
	Name      string   `json:"language,omitempty"`
	Fmt       string   `json:"fmt,omitempty"`
}

func (l Language) Greet(who string) string {
	greeting := l.Greetings[rand.Int()%len(l.Greetings)]
	if who == "" {
		return greeting
	}
	greeting = strings.Split(greeting, " ")[0]

	fmtStr := l.Fmt
	if fmtStr == "" {
		fmtStr = "%s %s!"
	}

	return fmt.Sprintf(fmtStr, greeting, who)
}
