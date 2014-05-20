package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tv42/base58"
)

func main() {
	flag.Parse()
	for _, enc := range flag.Args() {
		buf := []byte(enc)
		dec, err := base58.DecodeToBig(buf)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fmt.Println(dec)
		if err != nil {
			log.Fatal(err)
		}
	}
}
