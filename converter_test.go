package numconverter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Small numbers
func TestSmall(t *testing.T) {
	assert := require.New(t)

	assert.Equal(Convert(0), "zero")
	assert.Equal(Convert(1), "one")
	assert.Equal(Convert(2), "two")
	assert.Equal(Convert(3), "three")
	assert.Equal(Convert(4), "four")
	assert.Equal(Convert(5), "five")
	assert.Equal(Convert(6), "six")
	assert.Equal(Convert(7), "seven")
	assert.Equal(Convert(8), "eight")
	assert.Equal(Convert(9), "nine")

	assert.Equal(Convert(10), "ten")
	assert.Equal(Convert(11), "eleven")
	assert.Equal(Convert(12), "twelve")
	assert.Equal(Convert(13), "thirteen")
	assert.Equal(Convert(17), "seventeen")
	assert.Equal(Convert(19), "nineteen")
}

func TestTens(t *testing.T) {
	assert := require.New(t)

	assert.Equal(Convert(20), "twenty")
	assert.Equal(Convert(30), "thirty")
	assert.Equal(Convert(40), "forty")
	assert.Equal(Convert(50), "fifty")
	assert.Equal(Convert(60), "sixty")
	assert.Equal(Convert(70), "seventy")
	assert.Equal(Convert(80), "eighty")
	assert.Equal(Convert(90), "ninety")

	assert.Equal(Convert(44), "forty four")
	assert.Equal(Convert(77), "seventy seven")
	assert.Equal(Convert(99), "ninety nine")

	assert.Equal(Convert(-45), "minus forty five")
}

func TestHundreds(t *testing.T) {
	assert := require.New(t)

	assert.Equal(Convert(100), "one hundred")
	assert.Equal(Convert(121), "one hundred twenty one")
	assert.Equal(ConvertAnd(121), "one hundred and twenty one")
	assert.Equal(Convert(200), "two hundred")
	assert.Equal(Convert(777), "seven hundred seventy seven")
	assert.Equal(Convert(990), "nine hundred ninety")
	assert.Equal(Convert(999), "nine hundred ninety nine")
}

func TestThousands(t *testing.T) {
	assert := require.New(t)

	assert.Equal(Convert(1000), "one thousand")
	assert.Equal(Convert(1002), "one thousand two")
	assert.Equal(Convert(5050), "five thousand fifty")
	assert.Equal(Convert(9999), "nine thousand nine hundred ninety nine")

	assert.Equal(Convert(10000), "ten thousand")
	assert.Equal(Convert(12053), "twelve thousand fifty three")
	assert.Equal(Convert(17482), "seventeen thousand four hundred eighty two")
	assert.Equal(Convert(19999), "nineteen thousand nine hundred ninety nine")
	assert.Equal(Convert(25012), "twenty five thousand twelve")
	assert.Equal(Convert(55897), "fifty five thousand eight hundred ninety seven")
	assert.Equal(Convert(82847), "eighty two thousand eight hundred forty seven")
	assert.Equal(Convert(99999), "ninety nine thousand nine hundred ninety nine")
}

func TestLakh(t *testing.T) {
	assert := require.New(t)

	assert.Equal(Convert(100000), "one lakh")
	assert.Equal(Convert(100005), "one lakh five")
	assert.Equal(Convert(101010), "one lakh one thousand ten")
	assert.Equal(Convert(150432), "one lakh fifty thousand four hundred thirty two")
	assert.Equal(Convert(999913), "nine lakh ninety nine thousand nine hundred thirteen")

	assert.Equal(Convert(1000000), "ten lakh")
	assert.Equal(Convert(1200000), "twelve lakh")
	assert.Equal(Convert(1993672), "nineteen lakh ninety three thousand six hundred seventy two")
	assert.Equal(Convert(2500000), "twenty five lakh")

	assert.Equal(Convert(5555655), "fifty five lakh fifty five thousand six hundred fifty five")
	assert.Equal(Convert(9999999), "ninety nine lakh ninety nine thousand nine hundred ninety nine")
}

func TestCrore(t *testing.T) {
	assert := require.New(t)

	assert.Equal(Convert(10000000), "one crore")
	assert.Equal(Convert(96273927), "nine crore sixty two lakh seventy three thousand nine hundred twenty seven")

	assert.Equal(Convert(100000000), "ten crore")
	assert.Equal(Convert(163527819), "sixteen crore thirty five lakh twenty seven thousand eight hundred nineteen")
	assert.Equal(Convert(767873521), "seventy six crore seventy eight lakh seventy three thousand five hundred twenty one")
	assert.Equal(Convert(999999999), "ninety nine crore ninety nine lakh ninety nine thousand nine hundred ninety nine")

	assert.Equal(Convert(1000000000), "one hundred crore")
	assert.Equal(Convert(9999999999), "nine hundred ninety nine crore ninety nine lakh ninety nine thousand nine hundred ninety nine")
}
