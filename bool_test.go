package strconvlen

import (
	"testing"
)

func TestBool(t *testing.T) {
	if l := Bool(true); l != 4 {
		t.Fatalf("expect 4 but got %d", l)
	}
	if l := Bool(false); l != 5 {
		t.Fatalf("expect 5 but got %d", l)
	}
}

// helper

/*
func randUintWithDigits(places, maxValue int) uint {
	if places < 1 || places > 20 {
		panic("places must be between 1-20")
	}
	if places == 1 {
		return uint(rand.Intn(10))
	}
	//  9223372036854775807
	// 18446744073709551615
	// 12345678901234567890
	if places >= 19 {
		v := randIntWithDigits(19) * -1
		return uint(v)
	}
	return uint(randIntWithDigits(places))
}
*/
