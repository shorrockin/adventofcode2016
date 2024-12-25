package crt

import "adventofcode2016/pkg/utils/assert"

// implementation of chineese remainder theorem which given a remainder
// and modulus will return the first common occurance of the two numbers
// where they intersect.
func CRT[T int | int64](remainders, moduli []T) T {
	if len(remainders) != len(moduli) {
		assert.Fail("expected remainders and moduli to be the same length")
		return -1
	}

	product := T(1)
	for _, modulus := range moduli {
		product *= modulus
	}

	sum := T(0)
	for i := range remainders {
		partialProduct := product / moduli[i]
		inverse := modInverse(partialProduct, moduli[i])
		sum += remainders[i] * inverse * partialProduct
	}

	return sum % product
}

func extendedGCD[T int | int64](a, b T) (T, T, T) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := extendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}

func modInverse[T int | int64](a, m T) T {
	gcd, x, _ := extendedGCD(a, m)
	if gcd != 1 {
		return 0 // Inverse doesn't exist if a and m are not coprime
	}
	return (x%m + m) % m
}
