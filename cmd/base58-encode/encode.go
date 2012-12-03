package main

import (
	"math/big"
	"flag"
	"fmt"
	"github.com/tv42/base58"
	"log"
)

func main() {
	flag.Parse()
	for _, dec := range flag.Args() {
		n := new(big.Int)
		_, err := fmt.Sscan(dec, n)
		if err != nil {
			log.Fatal(err)
		}

		buf := base58.EncodeBig(nil, n)
		_, err = fmt.Printf("%s\n", buf)
		if err != nil {
			log.Fatal(err)
		}
	}
}
