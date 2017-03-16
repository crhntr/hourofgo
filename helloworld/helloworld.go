// helloworld is a package showing one way to write a hello world program in go
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"flag"
)

var (
	languages = map[string]HWLang{}
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
	defaultLanguage := flag.String("lang", "es", "default language as standard abreviation")
	flag.Parse()
}


func main() {
	switch flag.Arg(1) {
	case "server":
		http.HandleFunc("/", )

	}

}




const helloWorldMessage = "Hello World!"

type HWLang struct {
	Hello []string `json:"hello,omitempty"`
	Fmt   string   `json:"fmt,omitempty"`
}

func sayHello(to, language string) string {
	if to == "" {
		to = "world"
	}

	if language != "" {

	}

	if langData, ok := languages[language]; ok {
		if langData.Fmt == "" {
			return fmt.Sprintf("%s %s!",
				langData.Hello[rand.Int()%len(langData.Hello)],
				to,
			)
		}

		return fmt.Sprintf(langData.Fmt,
			langData.Hello[rand.Int()%len(langData.Hello)],
			to,
		)
	}

	return ""
}
