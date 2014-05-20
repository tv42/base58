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
		num := new(big.Int)
		_, err := fmt.Sscan(dec, num)
		if err != nil {
			log.Fatal(err)
		}

		buf := base58.EncodeBig(nil, num)
		_, err = fmt.Printf("%s\n", buf)
		if err != nil {
			log.Fatal(err)
		}
	}
}
