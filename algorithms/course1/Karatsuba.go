package algorithms

import (
	"math"
	"strconv"
	"strings"
)

// Formula:
//   (10^(n/2) * a + b) * (10^(n/2) * c + d)
// = 10^n * ac + 10^(n/2) * (ad + bc) + bd

// With n = the number of digits of the bigger operand

// Compute recursively: ac, bd, (a + b) * (c + d)

func Karatsuba(x string, y string) string {
	n := int(math.Max(float64(len(x)), float64(len(y))))

	if n%2 != 0 {
		n++
	}

	// Padd insufficiently long operand with zeros
	x = paddFront(x, n-len(x))
	y = paddFront(y, n-len(y))

	if n <= 8 {
		num1, _ := strconv.Atoi(x)
		num2, _ := strconv.Atoi(y)

		mult := num1 * num2
		return strconv.Itoa(mult)
	}

	a := x[:n/2]
	b := x[n/2:]
	c := y[:n/2]
	d := y[n/2:]

	ac := Karatsuba(a, c)
	bd := Karatsuba(b, d)
	abcd := Karatsuba(addXY(a, b), addXY(c, d))

	return addXY(addXY(paddBack(ac, n), paddBack(subXYZ(abcd, ac, bd), n/2)), bd)
}

func addXY(x string, y string) string {
	n := int(math.Max(float64(len(x)), float64(len(y))))

	x = paddFront(x, n-len(x))
	y = paddFront(y, n-len(y))

	var carry int
	var res string

	for i := n - 1; i >= 0; i-- {
		dig1, _ := strconv.Atoi(string(x[i]))
		dig2, _ := strconv.Atoi(string(y[i]))

		tot := dig1 + dig2 + carry
		prod := tot % 10
		carry = tot / 10

		res = strconv.Itoa(prod) + res
	}

	if carry > 0 {
		res = "1" + res
	}

	return res
}

func subXYZ(x string, y string, z string) string {
	// Assume that x >= y + z holds
	sumYZ := addXY(y, z)
	n := len(x)
	sumYZ = paddFront(sumYZ, n-len(sumYZ))

	var carry int
	var res string

	for i := n - 1; i >= 0; i-- {
		dig1, _ := strconv.Atoi(string(x[i]))
		dig2, _ := strconv.Atoi(string(sumYZ[i]))

		sub := 10 + dig1 - dig2 - carry
		prod := sub % 10
		if sub < 10 {
			carry = 1
		} else {
			carry = 0
		}

		res = strconv.Itoa(prod) + res
	}

	return res
}

func paddFront(num string, times int) string {
	zeros := strings.Repeat("0", times)
	return zeros + num
}

func paddBack(num string, times int) string {
	zeros := strings.Repeat("0", times)
	return num + zeros
}
