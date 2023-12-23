package strconvlen

import (
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestUint8(t *testing.T) {
	for i := 2; i < 36; i++ {
		vlen := Uint8(0, i)
		vstr := strconv.FormatUint(uint64(0), i)
		if len(vstr) != vlen {
			t.Errorf("expect Uint8(v: %d, base: %d) == len(%q) == %d but got %d",
				0, i, vstr, len(vstr), vlen)
		}

		for j := 1; j <= 3; j++ {
			v := randUint8(j)
			vlen := Uint8(v, i)
			vstr := strconv.FormatUint(uint64(v), i)
			if len(vstr) != vlen {
				t.Errorf("expect Uint8(v: %d, base: %d) == len(%q) == %d but got %d",
					v, i, vstr, len(vstr), vlen)
			}
		}
	}
}

func TestInt8(t *testing.T) {
	for i := 2; i < 36; i++ {
		vlen := Int8(0, i)
		vstr := strconv.FormatInt(int64(0), i)
		if len(vstr) != vlen {
			t.Errorf("expect Int8(v: %d, base: %d) == len(%q) == %d but got %d",
				0, i, vstr, len(vstr), vlen)
		}

		for j := 1; j <= 3; j++ {
			v := int8(randUint8(j))
			vlen := Int8(int8(v), i)
			vstr := strconv.FormatInt(int64(v), i)
			if len(vstr) != vlen {
				t.Errorf("expect Int8(v: %d, base: %d) == len(%q) == %d but got %d",
					v, i, vstr, len(vstr), vlen)
			}
		}
	}
}

func TestUint16(t *testing.T) {
	const maxPlaces = 5

	for i := 2; i < 36; i++ {
		// test 0 works
		vlen := Uint16(0, i)
		vstr := strconv.FormatUint(uint64(0), i)
		if len(vstr) != vlen {
			t.Errorf("expect Uint16(v: %d, base: %d) == len(%q) == %d but got %d",
				0, i, vstr, len(vstr), vlen)
		}

		// feed in progressively bigger numbers
		for j := 1; j <= maxPlaces; j++ {
			v := randUint16(j)
			vlen := Uint16(v, i)
			vstr := strconv.FormatUint(uint64(v), i)
			if len(vstr) != vlen {
				t.Errorf("expect Uint16(v: %d, base: %d) == len(%q) == %d but got %d",
					v, i, vstr, len(vstr), vlen)
			}
		}
	}
}

func TestUint32(t *testing.T) {
	const maxPlaces = 10

	for i := 2; i < 36; i++ {
		vlen := Uint32(0, i)
		vstr := strconv.FormatUint(uint64(0), i)
		if len(vstr) != vlen {
			t.Errorf("expect Uint32(v: %d, base: %d) == len(%q) == %d but got %d",
				0, i, vstr, len(vstr), vlen)
		}

		for j := 1; j <= 10; j++ {
			v := randUint32(j)
			vlen := Uint32(v, i)
			vstr := strconv.FormatUint(uint64(v), i)
			if len(vstr) != vlen {
				t.Errorf("expect Uint32(v: %d, base: %d) == len(%q) == %d but got %d",
					v, i, vstr, len(vstr), vlen)
			}
		}
	}
}

func TestUint64(t *testing.T) {
	const maxPlaces = 20

	for i := 2; i < 36; i++ {
		vlen := Uint64(0, i)
		vstr := strconv.FormatUint(uint64(0), i)
		if len(vstr) != vlen {
			t.Errorf("expect Uint64(v: %d, base: %d) == len(%q) == %d but got %d",
				0, i, vstr, len(vstr), vlen)
		}

		for j := 1; j <= maxPlaces; j++ {
			v := randUint64WithPlaces(j)
			vlen := Uint64(v, i)
			vstr := strconv.FormatUint(uint64(v), i)
			if len(vstr) != vlen {
				t.Errorf("expect Uint64(v: %d, base: %d) == len(%q) == %d but got %d",
					v, i, vstr, len(vstr), vlen)
			}
		}
	}
}

func TestInt64(t *testing.T) {
	const maxPlaces = 19

	for i := 2; i < 36; i++ {
		vlen := Int64(0, i)
		vstr := strconv.FormatInt(int64(0), i)
		if len(vstr) != vlen {
			t.Errorf("expect Int64(v: %d, base: %d) == len(%q) == %d but got %d",
				0, i, vstr, len(vstr), vlen)
		}

		for j := 1; j <= maxPlaces; j++ {
			v := int64(randIntWithPlaces(j, 0, 0))
			vlen := Int64(v, i)
			vstr := strconv.FormatInt(int64(v), i)
			if len(vstr) != vlen {
				t.Errorf("expect Int64(v: %d, base: %d) == len(%q) == %d but got %d",
					v, i, vstr, len(vstr), vlen)
			}
		}
	}
}

// helpers

// randUint64WithPlaces returns a random positive unsigned integer with the
// specified number of places, where places must be between 1-20.
func randUint64WithPlaces(places int) uint64 {
	if places < 1 || places > 20 {
		panic("places must be between 1-20")
	}
	if places == 1 {
		return uint64(rand.Intn(10))
	}

	if places < 19 {
		v := int(math.Pow10(places - 1))     // 1, 10, 100, ...
		return uint64(rand.Intn(v*10-v) + v) // 10-99, 100-999, etc...
	}

	lower := big.NewInt(0)
	upper := big.NewInt(0)
	if places == 19 {
		lower.SetUint64(1000000000000000000)
		upper.Mul(lower, big.NewInt(10))
	} else {
		lower.SetUint64(10000000000000000000)
		upper.SetUint64(math.MaxUint64)
	}

	delta := big.NewInt(0)
	delta.Sub(upper, lower)

	rndnum := big.NewInt(0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndnum.Rand(r, delta)

	nn := big.NewInt(0)
	nn.Add(rndnum, lower)
	return nn.Uint64()
}

// randIntWithPlaces returns a random positive integer with the specified number
// of places, e.g. 0-9 where places == 1, 10-99 where places == 2, etc.
func randIntWithPlaces(places, maxPlaces, maxValue int) int {
	if places < 1 || places > 19 {
		panic("places must be between 1-19")
	}
	if places == 1 {
		return rand.Intn(10)
	}
	v := int(math.Pow10(places - 1)) // 10, 100, ...

	if maxPlaces > 0 {
		if places == maxPlaces {
			return rand.Intn(maxValue-v) + v
		}
	} else {
		if places == 19 {
			return rand.Intn(math.MaxInt-v) + v
		}
	}
	return rand.Intn(v*10-v) + v // 10-99, 100-999, etc...
}

func randUint8(places int) uint8 {
	if places < 1 || places > 3 {
		panic("places must be between 1-3")
	}
	return uint8(randIntWithPlaces(places, 3, math.MaxUint8))
}

func randUint16(places int) uint16 {
	if places < 1 || places > 5 {
		panic("places must be between 1-5")
	}
	return uint16(randIntWithPlaces(places, 5, math.MaxUint16))
}

func randUint32(places int) uint32 {
	if places < 1 || places > 10 {
		panic("places must be between 1-5")
	}
	return uint32(randIntWithPlaces(places, 10, math.MaxUint32))
}
