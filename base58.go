// Package base58 implements a human-friendly base58 encoding.
//
// As opposed to base64 and friends, base58 is typically used to
// convert integers. You can use big.Int.SetBytes to convert arbitrary
// bytes to an integer first, and big.Int.Bytes the other way around.
package base58

import (
	"math/big"
	"strconv"
)

const alphabet = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var decodeMap [256]*big.Int

var (
	// constant-like variables

	radix = big.NewInt(58)
	zero  = big.NewInt(0)
)

func init() {
	for i := 0; i < len(alphabet); i++ {
		decodeMap[alphabet[i]] = big.NewInt(int64(i))
	}
}

type CorruptInputError int64

func (e CorruptInputError) Error() string {
	return "illegal base58 data at input byte " + strconv.FormatInt(int64(e), 10)
}

// Decode a big integer from the bytes. Returns an error on corrupt
// input.
func DecodeToBig(src []byte) (*big.Int, error) {
	n := new(big.Int)
	for i, c := range src {
		b := decodeMap[c]
		if b == nil {
			return nil, CorruptInputError(i)
		}
		n.Mul(n, radix)
		n.Add(n, b)
	}
	return n, nil
}

// Encode encodes src, appending to dst. Be sure to use the returned
// new value of dst.
func EncodeBig(dst []byte, src *big.Int) []byte {
	start := len(dst)
	n := new(big.Int)
	n.Set(src)

	for n.Cmp(zero) > 0 {
		mod := new(big.Int)
		n.DivMod(n, radix, mod)
		dst = append(dst, alphabet[mod.Int64()])
	}

	for i, j := start, len(dst)-1; i < j; i, j = i+1, j-1 {
		dst[i], dst[j] = dst[j], dst[i]
	}
	return dst
}
