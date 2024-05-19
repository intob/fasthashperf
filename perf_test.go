package fasthashperf

import (
	"crypto/rand"
	"hash"
	"hash/fnv"
	"testing"

	murmur3spaghetti "github.com/spaolacci/murmur3"
	murmur3twmb "github.com/twmb/murmur3"
)

func run(h hash.Hash, l int, b *testing.B) {
	v := make([]byte, l)
	rand.Read(v)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Reset()
		h.Write(v)
		h.Sum(nil)
	}
}

func run64(h hash.Hash64, l int, b *testing.B) {
	v := make([]byte, l)
	rand.Read(v)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Reset()
		h.Write(v)
		h.Sum64()
	}
}

// Byte slices

func Benchmark_murmur3twmb_64_l16(b *testing.B) {
	run(murmur3twmb.New64(), 16, b)
}

func Benchmark_murmur3spaolacci_64_l16(b *testing.B) {
	run(murmur3spaghetti.New64(), 16, b)
}

func Benchmark_fnva_64_l16(b *testing.B) {
	run(fnv.New64a(), 16, b)
}

func Benchmark_fnv_64_l16(b *testing.B) {
	run(fnv.New64(), 16, b)
}

func Benchmark_murmur3twmb_64_l32(b *testing.B) {
	run(murmur3twmb.New64(), 32, b)
}

func Benchmark_murmur3spaolacci_64_l32(b *testing.B) {
	run(murmur3spaghetti.New64(), 32, b)
}

func Benchmark_fnva_64_l32(b *testing.B) {
	run(fnv.New64a(), 32, b)
}

func Benchmark_fnv_64_l32(b *testing.B) {
	run(fnv.New64(), 32, b)
}

func Benchmark_murmur3twmb_64_l512(b *testing.B) {
	run(murmur3twmb.New64(), 512, b)
}

func Benchmark_murmur3spaolacci_64_l512(b *testing.B) {
	run(murmur3spaghetti.New64(), 512, b)
}

func Benchmark_fnva_64_l512(b *testing.B) {
	run(fnv.New64a(), 512, b)
}

func Benchmark_fnv_64_l512(b *testing.B) {
	run(fnv.New64(), 512, b)
}

// Uint64 (much faster on 64-bit registers)

func Benchmark_uint64_murmur3twmb_64_l16(b *testing.B) {
	run64(murmur3twmb.New64(), 16, b)
}

func Benchmark_uint64_murmur3spaolacci_64_l16(b *testing.B) {
	run64(murmur3spaghetti.New64(), 16, b)
}

func Benchmark_uint64_fnva_64_l16(b *testing.B) {
	run64(fnv.New64a(), 16, b)
}

func Benchmark_uint64_fnv_64_l16(b *testing.B) {
	run64(fnv.New64(), 16, b)
}

func Benchmark_uint64_murmur3twmb_64_l32(b *testing.B) {
	run64(murmur3twmb.New64(), 32, b)
}

func Benchmark_uint64_murmur3spaolacci_64_l32(b *testing.B) {
	run64(murmur3spaghetti.New64(), 32, b)
}

func Benchmark_uint64_fnva_64_l32(b *testing.B) {
	run64(fnv.New64a(), 32, b)
}

func Benchmark_uint64_fnv_64_l32(b *testing.B) {
	run64(fnv.New64(), 32, b)
}

func Benchmark_uint64_murmur3twmb_64_l512(b *testing.B) {
	run64(murmur3twmb.New64(), 512, b)
}

func Benchmark_uint64_murmur3spaolacci_64_l512(b *testing.B) {
	run64(murmur3spaghetti.New64(), 512, b)
}

func Benchmark_uint64_fnva_64_l512(b *testing.B) {
	run64(fnv.New64a(), 512, b)
}

func Benchmark_uint64_fnv_64_l512(b *testing.B) {
	run64(fnv.New64(), 512, b)
}
