package internal

import (
	"math"
)

const (
	kBufferSize = 2200
	// 0x7FF0'0000'0000'0000
	kExponentMask            = 9218868437227405312
	kPhysicalSignificandSize = 52
	kExponentBias            = 0x3FF + kPhysicalSignificandSize
	chars                    = "0123456789abcdefghijklmnopqrstuvwxyz"
)

// FloatToStringRadix Translated from https://github.com/v8/v8/blob/master/src/numbers/conversions.cc#L1269
// and https://github.com/v8/v8/blob/master/src/numbers/double.h
func FloatToStringRadix(value float64, radix int, buf []byte) int {
	buffer := make([]byte, kBufferSize)
	integerCursor := kBufferSize / 2
	fractionCursor := integerCursor

	negative := value < 0
	if negative {
		value = -value
	}

	integer := math.Floor(value)
	fraction := value - integer
	delta := 0.5 * (double(value).NextDouble() - value)
	delta = math.Max(double(0.0).NextDouble(), delta)

	if fraction >= delta {
		buffer[fractionCursor] = 46
		fractionCursor++

		for {
			fraction *= float64(radix)
			delta *= float64(radix)

			digit := int(fraction)
			buffer[fractionCursor] = chars[digit]
			fractionCursor++

			fraction -= float64(digit)

			if fraction > 0.5 || (fraction == 0.5 && ((digit & 1) != 0)) {
				if fraction+delta > 1 {
					for {
						fractionCursor--
						if fractionCursor == kBufferSize/2 {
							integer += 1
							break
						}

						c := buffer[fractionCursor]
						var digit byte
						if c > 57 {
							digit = c - 97 + 10
						} else {
							digit = c - 48
						}

						if int(digit+1) < radix {
							buffer[fractionCursor] = chars[digit+1]
							fractionCursor++
							break
						}
					}
					break
				}
			}

			if !(fraction >= delta) {
				break
			}
		}
	}

	for double(integer/float64(radix)).Exponent() > 0 {
		integer /= float64(radix)
		integerCursor--
		buffer[integerCursor] = 48
	}

	for {
		remainder := math.Remainder(integer, float64(radix))
		integerCursor--
		buffer[integerCursor] = chars[int(remainder)]
		integer = (integer - remainder) / float64(radix)

		if !(integer > 0) {
			break
		}
	}

	if negative {
		integerCursor--
		buffer[integerCursor] = 45
	}

	//copy(buf[offset:], buffer[integerCursor:fractionCursor])
	copy(buf, buffer[integerCursor:fractionCursor])
	return fractionCursor - integerCursor
}

type double float64

func (d double) AsUint64() uint64 {
	return math.Float64bits(float64(d))
}

func (d double) NextDouble() float64 {
	d64 := d.AsUint64()

	// When the original float is negative, you must decrement
	// the uint to get the next greater double when converting back
	if d < 0 {
		return math.Float64frombits(d64 - 1)
	}

	return math.Float64frombits(d64 + 1)
}

func (d double) Exponent() int {
	d64 := d.AsUint64()

	biasedE := math.Float64bits(float64((d64 & kExponentMask) >> kPhysicalSignificandSize))
	return int(biasedE) - kExponentBias
}
