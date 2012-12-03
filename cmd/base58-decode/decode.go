package main

import (
	"flag"
	"fmt"
	"github.com/tv42/base58"
	"log"
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
