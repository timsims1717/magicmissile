package util

import (
	"fmt"
	"strings"
)

func RomanNumeral(i int) string {
	var sb strings.Builder
	for i >= 1000 {
		sb.WriteString("M")
		i -= 1000
	}
	for i >= 900 {
		sb.WriteString("CM")
		i -= 900
	}
	for i >= 500 {
		sb.WriteString("D")
		i -= 500
	}
	for i >= 400 {
		sb.WriteString("CD")
		i -= 400
	}
	for i >= 100 {
		sb.WriteString("C")
		i -= 100
	}
	for i >= 90 {
		sb.WriteString("XC")
		i -= 90
	}
	for i >= 50 {
		sb.WriteString("L")
		i -= 50
	}
	for i >= 40 {
		sb.WriteString("XL")
		i -= 40
	}
	for i >= 10 {
		sb.WriteString("X")
		i -= 10
	}
	for i >= 9 {
		sb.WriteString("IX")
		i -= 9
	}
	for i >= 5 {
		sb.WriteString("V")
		i -= 5
	}
	for i >= 4 {
		sb.WriteString("IV")
		i -= 4
	}
	for i >= 1 {
		sb.WriteString("I")
		i--
	}
	return sb.String()
}

func testRomanInner(i int) {
	fmt.Printf("Test %d: %s\n", i, RomanNumeral(i))
}

func TestRomanNumerals() {
	testRomanInner(1)
	testRomanInner(3)
	testRomanInner(4)
	testRomanInner(8)
	testRomanInner(9)
	testRomanInner(10)
	testRomanInner(16)
	testRomanInner(22)
	testRomanInner(43)
	testRomanInner(56)
	testRomanInner(78)
	testRomanInner(99)
	testRomanInner(101)
	testRomanInner(333)
	testRomanInner(412)
	testRomanInner(556)
	testRomanInner(692)
	testRomanInner(844)
	testRomanInner(919)
	testRomanInner(1000)
	testRomanInner(2023)
	testRomanInner(4601)
	testRomanInner(5000)

}
