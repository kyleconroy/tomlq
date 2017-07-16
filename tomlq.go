package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	toml "github.com/pelletier/go-toml"
	"github.com/pelletier/go-toml/query"
)

func main() {
	flag.Parse()
	q := flag.Arg(0)
	if q == "" {
		log.Fatal(errors.New("Missing TOML query string argument"))
	}
	conf, err := toml.LoadReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	result, err := query.CompileAndExecute(q, conf)
	if err != nil {
		log.Fatal(err)
	}
	if len(result.Values()) == 0 {
		log.Fatal(errors.New("No match found for path"))
	}
	fmt.Println(result.Values()[0])
}
