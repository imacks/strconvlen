package strconvlen

import (
	"math/bits"
)

const host32bit = ^uint(0)>>32 == 0

// Uint8 returns the same result as len(strconv.FormatUint(uint64(n), base)).
func Uint8(n byte, base int) int {
	if base < 2 || base > 36 {
		panic("strconvlen: illegal Uint8 base")
	}

	// For powers of 2, the general formula is:
	//
	//   bb := bits.Len8(b) // length in base 2
	//   strlen(b, 2^p) := floor((bb+p-1)/p)
	//
	// For example:
	//
	//   strlen(b, 2^2) = floor((bb+1)/2)
	//   strlen(b, 2^3) = floor((bb+2)/3)
	//   ...
	//
	// This can be optimized a little further in Go:
	//
	//   bb := 63 - bits.LeadingZeros(uint(base))
	//   return (bits.Len(n) + (bb - 1)) / bb
	//
	// In practice, it turns out that division takes a toll. Everything cancels 
	// out in base 2, and you can get around the division in case of base 4 by 
	// bitshift right 1 place. The decision to go with formula or using a binary 
	// tree for other bases is a matter of emperical benchmarking.
	//
	// Remarks: xxluv for contribution by TheMeow =(^.^)=

	switch base {
	case 10:
		if n < 100 {
			if n < 10 {
				return 1
			}
			return 2
		}
		return 3
	case 2:
		if n == 0 {
			return 1
		}
		return bits.Len8(n)
	case 4:
		if n == 0 {
			return 1
		}
		return (bits.Len8(n) + 1) >> 1
	case 16:
		if n < 16 {
			return 1
		}
		return 2
	case 32:
		if n < 32 {
			return 1
		}
		return 2
	case 3:
		if n < 27 {
			if n < 9 {
				if n < 3 {
					return 1
				}
				return 2
			}
			return 3
		}
		if n < 243 {
			if n < 81 {
				return 4
			}
			return 5
		}
		return 6
	case 5, 6:
		bb := byte(base)
		if n < bb {
			return 1
		}
		if n < bb*bb {
			return 2
		}
		if n < bb*bb*bb {
			return 3
		}
		return 4
	case 7, 8, 9, 11, 12, 13, 14, 15:
		bb := byte(base)
		if n < bb {
			return 1
		}
		if n < bb*bb {
			return 2
		}
		return 3
	default:
		// base 17-36 (excl 32)
		if n < byte(base) {
			return 1
		}
		return 2
	}
}

// Int8 returns the same result as len(strconv.FormatInt(int64(n), base)).
func Int8(n int8, base int) int {
	if base < 2 || base > 36 {
		panic("strconvlen: illegal Int8 base")
	}

	if base == 10 {
		if n < 0 {
			if n > -100 {
				if n > -10 {
					return 2
				}
				return 3
			}
			return 4
		}
		if n < 100 {
			if n < 10 {
				return 1
			}
			return 2
		}
		return 3
	}

	if n < 0 {
		n = n * -1
		return Uint8(uint8(n), base) + 1
	}
	return Uint8(uint8(n), base)
}

// Uint16 returns the same result as len(strconv.FormatUint(uint64(n), base)).
func Uint16(n uint16, base int) int {
	if base < 2 || base > 36 {
		panic("strconvlen: illegal Uint16 base")
	}

	switch base {
	case 10:
		if n < 1000 {
			if n < 100 {
				if n < 10 {
					return 1
				}
				return 2
			}
			return 3
		}
		if n < 10000 {
			return 4
		}
		return 5
	case 2:
		if n == 0 {
			return 1
		}
		return bits.Len16(n)
	case 4:
		if n == 0 {
			return 1
		}
		return (bits.Len16(n) + 1) >> 1
	case 8:
		if n < 8*8*8 {
			if n < 8*8 {
				if n < 8 {
					return 1
				}
				return 2
			}
			return 3
		}
		if n < 8*8*8*8*8 {
			if n < 8*8*8*8 {
				return 4
			}
			return 5
		}
		return 6
	case 16:
		if n < 16*16 {
			if n < 16 {
				return 1
			}
			return 2
		}
		if n < 16*16*16 {
			return 3
		}
		return 4
	case 32:
		if n < 32*32 {
			if n < 32 {
				return 1
			}
			return 2
		}
		if n < 32*32*32 {
			return 3
		}
		return 4
	case 3:
		if n < 3*3*3*3*3 {
			if n < 3*3*3 {
				if n < 3*3 {
					if n < 3 {
						return 1
					}
					return 2
				}
				return 3
			}
			if n < 3*3*3*3 {
				return 4
			}
			return 5
		}
		if n < 3*3*3*3*3*3*3*3 {
			if n < 3*3*3*3*3*3*3 {
				if n < 3*3*3*3*3*3 {
					return 6
				}
				return 7
			}
			return 8
		}
		if n < 3*3*3*3*3*3*3*3*3*3 {
			if n < 3*3*3*3*3*3*3*3*3 {
				return 9
			}
			return 10
		}
		return 11
	case 5:
		if n < 5*5*5 {
			if n < 5*5 {
				if n < 5 {
					return 1
				}
				return 2
			}
			return 3
		}
		if n < 5*5*5*5*5*5 {
			if n < 5*5*5*5*5 {
				if n < 5*5*5*5 {
					return 4
				}
				return 5
			}
			return 6
		}
		return 7
	case 6:
		if n < 6*6*6 {
			if n < 6*6 {
				if n < 6 {
					return 1
				}
				return 2
			}
			return 3
		}
		if n < 6*6*6*6*6*6 {
			if n < 6*6*6*6*6 {
				if n < 6*6*6*6 {
					return 4
				}
				return 5
			}
			return 6
		}
		return 7
	case 7:
		if n < 7*7*7 {
			if n < 7*7 {
				if n < 7 {
					return 1
				}
				return 2
			}
			return 3
		}
		if n < 7*7*7*7*7 {
			if n < 7*7*7*7 {
				return 4
			}
			return 5
		}
		return 6
	case 9:
		if n < 9*9*9 {
			if n < 9*9 {
				if n < 9 {
					return 1
				}
				return 2
			}
			return 3
		}
		if n < 9*9*9*9*9 {
			if n < 9*9*9*9 {
				return 4
			}
			return 5
		}
		return 6
	case 11, 12, 13, 14, 15:
		bb := uint16(base)
		if n < bb {
			return 1
		}
		if n < bb*bb {
			return 2
		}
		if n < bb*bb*bb {
			return 3
		}
		if n < bb*bb*bb*bb {
			return 4
		}
		return 5
	default:
		// base 17-36 (excl 32)
		bb := uint16(base)
		if n < bb {
			return 1
		}
		if n < bb*bb {
			return 2
		}
		if n < bb*bb*bb {
			return 3
		}
		return 4
	}
}

// Int16 returns the same result as len(strconv.FormatInt(int64(n), base)).
func Int16(n int16, base int) int {
	if base < 2 || base > 36 {
		panic("strconvlen: illegal Int16 base")
	}

	if base == 10 {
		if n < 0 {
			if n > -1000 {
				if n > -100 {
					if n > -10 {
						return 2
					}
					return 3
				}
				return 4
			}
			if n > -10000 {
				return 5
			}
			return 6
		}
		if n < 1000 {
			if n < 100 {
				if n < 10 {
					return 1
				}
				return 2
			}
			return 3
		}
		if n < 10000 {
			return 4
		}
		return 5
	}

	if n < 0 {
		n = n * -1
		return Uint16(uint16(n), base) + 1
	}
	return Uint16(uint16(n), base)
}

// Uint32 returns the same result as len(strconv.FormatUint(uint64(n), base)).
func Uint32(n uint32, base int) int {
	if base < 2 || base > 36 {
		panic("strconvlen: illegal Uint32 base")
	}

	switch base {
	case 10:
		if n < 100000 {
			if n < 1000 {
				if n < 100 {
					if n < 10 {
						return 1
					}
					return 2
				}
				return 3
			}
			if n < 10000 {
				return 4
			}
			return 5
		}
		if n < 100000000 {
			if n < 10000000 {
				if n < 1000000 {
					return 6
				}
				return 7
			}
			return 8
		}
		if n < 1000000000 {
			return 9
		}
		return 10
	case 2:
		if n == 0 {
			return 1
		}
		return bits.Len32(n)
	case 4:
		if n == 0 {
			return 1
		}
		return (bits.Len32(n) + 1) >> 1
	case 8:
		if n < 8*8*8*8*8*8 {
			if n < 8*8*8 {
				if n < 8*8 {
					if n < 8 {
						return 1
					}
					return 2
				}
				return 3
			}
			if n < 8*8*8*8*8 {
				if n < 8*8*8*8 {
					return 4
				}
				return 5
			}
			return 6
		}
		if n < 8*8*8*8*8*8*8*8 {
			if n < 8*8*8*8*8*8*8 {
				return 7
			}
			return 8
		}
		if n < 8*8*8*8*8*8*8*8*8*8 {
			if n < 8*8*8*8*8*8*8*8*8 {
				return 9
			}
			return 10
		}
		return 11
	case 16:
		if n < 16*16*16*16 {
			if n < 16*16 {
				if n < 16 {
					return 1
				}
				return 2
			}
			if n < 16*16*16 {
				return 3
			}
			return 4
		}
		if n < 16*16*16*16*16*16 {
			if n < 16*16*16*16*16 {
				return 5
			}
			return 6
		}
		if n < 16*16*16*16*16*16*16 {
			return 7
		}
		return 8
	case 32:
		if n < 32*32*32*32 {
			if n < 32*32 {
				if n < 32 {
					return 1
				}
				return 2
			}
			if n < 32*32*32 {
				return 3
			}
			return 4
		}
		if n < 32*32*32*32*32*32 {
			if n < 32*32*32*32*32 {
				return 5
			}
			return 6
		}
		return 7
	case 3:
		if n < 177147 {
			if n < 729 {
				if n < 27 {
					if n < 9 {
						if n < 3 {
							return 1
						}
						return 2
					}
					return 3
				}
				if n < 243 {
					if n < 81 {
						return 4
					}
					return 5
				}
				return 6
			}
			if n < 19683 {
				if n < 6561 {
					if n < 2187 {
						return 7
					}
					return 8
				}
				return 9
			}
			if n < 59049 {
				return 10
			}
			return 11
		}
		if n < 43046721 {
			if n < 4782969 {
				if n < 1594323 {
					if n < 531441 {
						return 12
					}
					return 13
				}
				return 14
			}
			if n < 14348907 {
				return 15
			}
			return 16
		}
		if n < 1162261467 {
			if n < 387420489 {
				if n < 129140163 {
					return 17
				}
				return 18
			}
			return 19
		}
		if n < 3486784401 {
			return 20
		}
		return 21
	default:
		bbase := uint32(base)

		if n < bbase {
			return 1
		}

		var maxpow int
		if base >= 24 {
			maxpow = 6
		} else if base >= 17 {
			maxpow = 7
		} else if base >= 12 {
			maxpow = 8
		} else {
			switch base {
			case 5:
				maxpow = 13
			case 6:
				maxpow = 12
			case 7:
				maxpow = 11
			case 9:
				maxpow = 10
			case 11:
				maxpow = 9
			}
		}

		mbase := uint32(base)
		for i := 2; i < (maxpow + 1); i++ {
			mbase *= bbase
			if n < mbase {
				return i
			}
		}

		return maxpow + 1
	}
}

// Int32 returns the same result as len(strconv.FormatInt(int64(n), base)).
func Int32(n int32, base int) int {
	if base < 2 || base > 36 {
		panic("strconvlen: illegal Int32 base")
	}

	if base == 10 {
		if n < 0 {
			if n > -100000 {
				if n > -1000 {
					if n > -100 {
						if n > -10 {
							return 3
						}
						return 4
					}
					return 4
				}
				if n > -10000 {
					return 5
				}
				return 6
			}
			if n > -100000000 {
				if n > -10000000 {
					if n > -1000000 {
						return 7
					}
					return 8
				}
				return 9
			}
			if n > -1000000000 {
				return 10
			}
			return 11
		}
		if n < 100000 {
			if n < 1000 {
				if n < 100 {
					if n < 10 {
						return 1
					}
					return 2
				}
				return 3
			}
			if n < 10000 {
				return 4
			}
			return 5
		}
		if n < 100000000 {
			if n < 10000000 {
				if n < 1000000 {
					return 6
				}
				return 7
			}
			return 8
		}
		if n < 1000000000 {
			return 9
		}
		return 10
	}

	if n < 0 {
		n = n * -1
		return Uint32(uint32(n), base) + 1
	}
	return Uint32(uint32(n), base)
}

// Uint64 returns the same result as len(strconv.FormatUint(n, base)).
func Uint64(n uint64, base int) int {
	if base < 2 || base > 36 {
		panic("strconvlen: illegal Uint64 base")
	}

	switch base {
	case 10:
		if n < 10000000000 {
			if n < 100000 {
				if n < 1000 {
					if n < 100 {
						if n < 10 {
							return 1
						}
						return 2
					}
					return 3
				}
				if n < 10000 {
					return 4
				}
				return 5
			}
			if n < 100000000 {
				if n < 10000000 {
					if n < 1000000 {
						return 6
					}
					return 7
				}
				return 8
			}
			if n < 1000000000 {
				return 9
			}
			return 10
		}
		if n < 1000000000000000 {
			if n < 10000000000000 {
				if n < 1000000000000 {
					if n < 100000000000 {
						return 11
					}
					return 12
				}
				return 13
			}
			if n < 100000000000000 {
				return 14
			}
			return 15
		}
		if n < 1000000000000000000 {
			if n < 100000000000000000 {
				if n < 10000000000000000 {
					return 16
				}
				return 17
			}
			return 18
		}
		if n < 10000000000000000000 {
			return 19
		}
		return 20
	case 2:
		if n == 0 {
			return 1
		}
		return bits.Len64(n)
	case 4:
		if n == 0 {
			return 1
		}
		return (bits.Len64(n) + 1) >> 1
	case 8:
		if n < 8*8*8*8*8*8*8*8*8*8*8 {
			if n < 8*8*8*8*8*8 {
				if n < 8*8*8 {
					if n < 8*8 {
						if n < 8 {
							return 1
						}
						return 2
					}
					return 3
				}
				if n < 8*8*8*8*8 {
					if n < 8*8*8*8 {
						return 4
					}
					return 5
				}
				return 6
			}
			if n < 8*8*8*8*8*8*8*8*8*8 {
				if n < 8*8*8*8*8*8*8*8 {
					if n < 8*8*8*8*8*8*8 {
						return 7
					}
					return 8
				}
				if n < 8*8*8*8*8*8*8*8*8 {
					return 9
				}
				return 10
			}
			return 11
		}
		if n < 8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8 {
			if n < 8*8*8*8*8*8*8*8*8*8*8*8*8 {
				if n < 8*8*8*8*8*8*8*8*8*8*8*8 {
					return 12
				}
				return 13
			}
			if n < 8*8*8*8*8*8*8*8*8*8*8*8*8*8*8 {
				if n < 8*8*8*8*8*8*8*8*8*8*8*8*8*8 {
					return 14
				}
				return 15
			}
			return 16
		}
		if n < 8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8 {
			if n < 8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8 {
				if n < 8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8 {
					return 17
				}
				return 18
			}
			return 19
		}
		if n < 8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8 {
			if n < 8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8*8 {
				return 20
			}
			return 21
		}
		return 22
	case 16:
		if n < 16*16*16*16*16*16*16*16 {
			if n < 16*16*16*16 {
				if n < 16*16 {
					if n < 16 {
						return 1
					}
					return 2
				}
				if n < 16*16*16 {
					return 3
				}
				return 4
			}
			if n < 16*16*16*16*16*16 {
				if n < 16*16*16*16*16 {
					return 5
				}
				return 6
			}
			if n < 16*16*16*16*16*16*16 {
				return 7
			}
			return 8
		}
		if n < 16*16*16*16*16*16*16*16*16*16*16*16 {
			if n < 16*16*16*16*16*16*16*16*16*16 {
				if n < 16*16*16*16*16*16*16*16*16 {
					return 9
				}
				return 10
			}
			if n < 16*16*16*16*16*16*16*16*16*16*16 {
				return 11
			}
			return 12
		}
		if n < 16*16*16*16*16*16*16*16*16*16*16*16*16*16 {
			if n < 16*16*16*16*16*16*16*16*16*16*16*16*16 {
				return 13
			}
			return 14
		}
		if n < 16*16*16*16*16*16*16*16*16*16*16*16*16*16*16 {
			return 15
		}
		return 16
	case 32:
		if n < 32*32*32*32*32*32 {
			if n < 32*32*32 {
				if n < 32*32 {
					if n < 32 {
						return 1
					}
					return 2
				}
				return 3
			}
			if n < 32*32*32*32*32 {
				if n < 32*32*32*32 {
					return 4
				}
				return 5
			}
			return 6
		}
		if n < 32*32*32*32*32*32*32*32*32 {
			if n < 32*32*32*32*32*32*32*32 {
				if n < 32*32*32*32*32*32*32 {
					return 7
				}
				return 8
			}
			return 9
		}
		if n < 32*32*32*32*32*32*32*32*32*32*32 {
			if n < 32*32*32*32*32*32*32*32*32*32 {
				return 10
			}
			return 11
		}
		if n < 32*32*32*32*32*32*32*32*32*32*32*32 {
			return 12
		}
		return 13
	default:
		bbase := uint64(base)

		if n < bbase {
			return 1
		}

		var maxpow int
		if base >= 31 {
			maxpow = 12
		} else if base >= 24 {
			maxpow = 13
		} else if base >= 20 {
			maxpow = 14
		} else if base >= 16 {
			maxpow = 15
		} else if base >= 14 {
			maxpow = 16
		} else if base >= 12 {
			maxpow = 17
		} else {
			switch base {
			case 3:
				maxpow = 40
			case 5:
				maxpow = 27
			case 6:
				maxpow = 24
			case 7:
				maxpow = 22
			case 9:
				maxpow = 20
			case 11:
				maxpow = 18
			}
		}

		mbase := uint64(base)
		for i := 2; i < (maxpow + 1); i++ {
			mbase *= bbase
			if n < mbase {
				return i
			}
		}

		return maxpow + 1
	}
}

// Int64 returns the same result as len(strconv.FormatInt(n, base)).
func Int64(n int64, base int) int {
	if base < 2 || base > 36 {
		panic("strconvlen: illegal Int64 base")
	}

	switch base {
	case 10:
		if n < 0 {
			if n > -10000000000 {
				if n > -100000 {
					if n > -1000 {
						if n > -100 {
							if n > -10 {
								return 2
							}
							return 3
						}
						return 4
					}
					if n > -10000 {
						return 5
					}
					return 6
				}
				if n > -100000000 {
					if n > -10000000 {
						if n > -1000000 {
							return 7
						}
						return 8
					}
					return 9
				}
				if n > -1000000000 {
					return 10
				}
				return 11
			}
			if n > -1000000000000000 {
				if n > -10000000000000 {
					if n > -1000000000000 {
						if n > -100000000000 {
							return 12
						}
						return 13
					}
					return 14
				}
				if n > -100000000000000 {
					return 15
				}
				return 16
			}
			if n > -1000000000000000000 {
				if n > -100000000000000000 {
					if n > -10000000000000000 {
						return 17
					}
					return 18
				}
				return 19
			}
			return 20
		}
		if n < 10000000000 {
			if n < 100000 {
				if n < 1000 {
					if n < 100 {
						if n < 10 {
							return 1
						}
						return 2
					}
					return 3
				}
				if n < 10000 {
					return 4
				}
				return 5
			}
			if n < 100000000 {
				if n < 10000000 {
					if n < 1000000 {
						return 6
					}
					return 7
				}
				return 8
			}
			if n < 1000000000 {
				return 9
			}
			return 10
		}
		if n < 1000000000000000 {
			if n < 10000000000000 {
				if n < 1000000000000 {
					if n < 100000000000 {
						return 11
					}
					return 12
				}
				return 13
			}
			if n < 100000000000000 {
				return 14
			}
			return 15
		}
		if n < 1000000000000000000 {
			if n < 100000000000000000 {
				if n < 10000000000000000 {
					return 16
				}
				return 17
			}
			return 18
		}
		return 19
	default:
		if n < 0 {
			n = n * -1
			return Uint64(uint64(n), base) + 1
		}
		return Uint64(uint64(n), base)
	}
}

// Int is the same as Int64 or Int32, depending on CPU architecture.
func Int(n, base int) int {
	if !host32bit {
		return Int64(int64(n), base)
	}
	return Int32(int32(n), base)
}

// Uint is the same as Uint64 or Uint32, depending on CPU architecture.
func Uint(n uint, base int) int {
	if !host32bit {
		return Uint64(uint64(n), base)
	}
	return Uint32(uint32(n), base)
}
