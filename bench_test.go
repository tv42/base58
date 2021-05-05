package base58_test

import (
	"math/big"
	"testing"

	"github.com/tv42/base58"
	"strconv"
)

func fixedBigInt(numBytes int) *big.Int {
	data := make([]byte, numBytes)
	data[0] = 0xff // without this, it's just an inefficient zero
	num := new(big.Int)
	num.SetBytes(data)
	return num
}

// 8192-byte benchmarks mirror the ones in encoding/base64, and
// results should be comparable to those.
//
// 8-byte benchmarks are more like the uses that this library was
// meant for: smallish identifiers.

func BenchmarkEncode(b *testing.B) {
	run := func(numBytes int) {
		b.Run(strconv.Itoa(numBytes), func(b *testing.B) {
			num := fixedBigInt(numBytes)
			b.SetBytes(int64(len(num.Bytes())))
			buf := make([]byte, 12000)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = base58.EncodeBig(buf[:0], num)
			}
		})
	}
	run(8)
	run(8192)
}

func BenchmarkBase58Decode(b *testing.B) {
	run := func(numBytes int) {
		b.Run(strconv.Itoa(numBytes), func(b *testing.B) {
			num := fixedBigInt(numBytes)
			encoded := base58.EncodeBig(nil, num)
			b.SetBytes(int64(len(encoded)))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = base58.DecodeToBig(encoded)
			}
		})
	}
	run(8)
	run(8192)
}
