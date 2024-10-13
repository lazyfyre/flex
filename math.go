package flex

import "math"

// from https://github.com/rkusa/gm/blob/master/math32/bits.go

var (
	NAN = math.NaN()
)

// NaN returns an IEEE 754 “not-a-number” value.
func NaN() float64 { return math.NaN() }

// IsNaN reports whether f is an IEEE 754 “not-a-number” value.
func IsNaN(f float64) (is bool) {
	return math.IsNaN(f)
}

func feq(a, b float64) bool {
	if IsNaN(a) && IsNaN(b) {
		return true
	}
	return a == b
}

// https://github.com/evanphx/ulysses-libc/blob/master/src/math/fmaxf.c
func fmaxf(a float64, b float64) float64 {
	if IsNaN(a) {
		return b
	}
	if IsNaN(b) {
		return a
	}
	// TODO: signed zeros
	if a > b {
		return a
	}
	return b
}

// https://github.com/evanphx/ulysses-libc/blob/master/src/math/fminf.c
func fminf(a float64, b float64) float64 {
	if IsNaN(a) {
		return b
	}
	if IsNaN(b) {
		return a
	}
	// TODO: signed zeros
	if a < b {
		return a
	}
	return b
}

func fabs(x float64) float64 {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0 // return correctly abs(-0)
	}
	return x
}

func fmodf(x, y float64) float64 {
	res := math.Mod(float64(x), float64(y))
	return float64(res)
}
