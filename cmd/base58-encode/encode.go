package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/tv42/base58"
)

func main() {
	flag.Parse()
	for _, dec := range flag.Args() {
		num := new(big.Int)
		if _, ok := num.SetString(dec, 10); !ok {
			log.Fatalf("not a number: %s", dec)
		}

		buf := base58.EncodeBig(nil, num)
		if _, err := fmt.Printf("%s\n", buf); err != nil {
			log.Fatal(err)
		}
	}
}
