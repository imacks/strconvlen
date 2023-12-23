package strconvlen

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func BenchmarkUint8(b *testing.B) {
	const places = 3

	if envv, ok := os.LookupEnv("INT_BENCH"); ok {
		if envv != "uint8" {
			b.Skipf("skip benchmark with env INT_BENCH=%s", envv)
		}
	}

	samples := make([]uint8, places)
	for i := 0; i < len(samples); i++ {
		samples[i] = randUint8(i + 1)
	}

	for _, i := range getBaseRange() {
		for j, v := range samples {
			b.Run(fmt.Sprintf("base%d_%dp", i, j+1), func(b2 *testing.B) {
				b2.ReportAllocs()
				for k := 0; k < b2.N; k++ {
					Uint8(v, i)
				}
			})
		}
	}
}

func BenchmarkInt8(b *testing.B) {
	const places = 3

	if envv, ok := os.LookupEnv("INT_BENCH"); ok {
		if envv != "int8" {
			b.Skipf("skip benchmark with env INT_BENCH=%s", envv)
		}
	}

	samples := make([]int8, places)
	for i := 0; i < len(samples); i++ {
		samples[i] = int8(randUint8(i + 1))
	}

	for _, i := range getBaseRange() {
		for j, v := range samples {
			b.Run(fmt.Sprintf("base%d_%dp", i, j+1), func(b2 *testing.B) {
				b2.ReportAllocs()
				for k := 0; k < b2.N; k++ {
					Int8(v, i)
				}
			})
		}
	}
}

func BenchmarkUint16(b *testing.B) {
	const places = 5

	if envv, ok := os.LookupEnv("INT_BENCH"); ok {
		if envv != "uint16" {
			b.Skipf("skip benchmark with env INT_BENCH=%s", envv)
		}
	}

	samples := make([]uint16, places)
	for i := 0; i < len(samples); i++ {
		samples[i] = randUint16(i + 1)
	}

	for _, i := range getBaseRange() {
		for j, v := range samples {
			b.Run(fmt.Sprintf("base%d_%dp", i, j+1), func(b2 *testing.B) {
				b2.ReportAllocs()
				for k := 0; k < b2.N; k++ {
					Uint16(v, i)
				}
			})
		}
	}
}

func BenchmarkUint32(b *testing.B) {
	const places = 10

	if envv, ok := os.LookupEnv("INT_BENCH"); ok {
		if envv != "uint32" {
			b.Skipf("skip benchmark with env INT_BENCH=%s", envv)
		}
	}

	samples := make([]uint32, places)
	for i := 0; i < len(samples); i++ {
		samples[i] = randUint32(i + 1)
	}

	for _, i := range getBaseRange() {
		for j, v := range samples {
			b.Run(fmt.Sprintf("base%d_%dp", i, j+1), func(b2 *testing.B) {
				b2.ReportAllocs()
				for k := 0; k < b2.N; k++ {
					Uint32(v, i)
				}
			})
		}
	}
}

func BenchmarkUint64(b *testing.B) {
	const places = 20

	if envv, ok := os.LookupEnv("INT_BENCH"); ok {
		if envv != "uint64" {
			b.Skipf("skip benchmark with env INT_BENCH=%s", envv)
		}
	}

	samples := make([]uint64, places)
	for i := 0; i < len(samples); i++ {
		samples[i] = randUint64WithPlaces(i + 1)
	}

	for _, i := range getBaseRange() {
		for j, v := range samples {
			b.Run(fmt.Sprintf("base%d_%dp", i, j+1), func(b2 *testing.B) {
				b2.ReportAllocs()
				for k := 0; k < b2.N; k++ {
					Uint64(v, i)
				}
			})
		}
	}
}

func BenchmarkInt64(b *testing.B) {
	const places = 19

	if envv, ok := os.LookupEnv("INT_BENCH"); ok {
		if envv != "int64" {
			b.Skipf("skip benchmark with env INT_BENCH=%s", envv)
		}
	}

	samples := make([]int64, places)
	for i := 0; i < len(samples); i++ {
		samples[i] = int64(randIntWithPlaces(i+1, 0, 0))
	}

	for _, i := range getBaseRange() {
		for j, v := range samples {
			b.Run(fmt.Sprintf("base%d_%dp", i, j+1), func(b2 *testing.B) {
				b2.ReportAllocs()
				for k := 0; k < b2.N; k++ {
					Int64(v, i)
				}
			})
		}
	}
}

func BenchmarkFormatInt(b *testing.B) {
	const places = 19

	if envv, ok := os.LookupEnv("INT_BENCH"); ok {
		if envv != "stdlib" {
			b.Skipf("skip benchmark with env INT_BENCH=%s", envv)
		}
	}

	samples := make([]int64, places)
	for i := 0; i < len(samples); i++ {
		samples[i] = int64(randIntWithPlaces(i+1, 0, 0))
	}

	for _, i := range getBaseRange() {
		for j, v := range samples {
			b.Run(fmt.Sprintf("base%d_%dp", i, j+1), func(b2 *testing.B) {
				b2.ReportAllocs()
				for k := 0; k < b2.N; k++ {
					strconv.FormatInt(int64(v), i)
				}
			})
		}
	}
}

func getBaseRange() []int {
	_, ok := os.LookupEnv("POW2_BENCH")
	if ok {
		return []int{2, 4, 8, 16, 32}
	}

	_, ok = os.LookupEnv("POW10_BENCH")
	if ok {
		return []int{10}
	}

	_, ok = os.LookupEnv("UNUSUAL_BASE_BENCH")
	if ok {
		return []int{3, 5, 6, 7, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36}
	}

	bb := make([]int, 36-2)
	for i := 2; i < 37; i++ {
		bb[i-2] = i
	}
	return bb
}
