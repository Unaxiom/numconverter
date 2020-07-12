package numconverter

import (
	"math"
)

var _smallNumbers = []string{
	"zero", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen",
	"fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}
var _tens = []string{
	"", "", "twenty", "thirty", "forty", "fifty",
	"sixty", "seventy", "eighty", "ninety",
}
var _scaleNum = []string{
	"", "thousand", "lakh", "crore",
}

// how many digit groups to process
const groupsCount int = 4

// Convert converts number into the words representation.
func Convert(number int) string {
	return convert(number, false)
}

// ConvertAnd converts number into the words representation
// with " and " added between number groups.
func ConvertAnd(number int) string {
	return convert(number, true)
}

func convert(number int, useAnd bool) string {
	// Zero rule
	if number == 0 {
		return _smallNumbers[0]
	}

	// Divide into groups
	var groups [groupsCount]int64
	positiveNum := int64(math.Abs(float64(number)))

	// Form three-digit groups
	for i := 0; i < groupsCount; i++ {
		var divisor int64
		if i == 0 {
			divisor = 1000
		} else {
			divisor = 100
		}

		if i == (groupsCount - 1) {
			groups[i] = positiveNum
		} else {
			groups[i] = positiveNum % divisor
			positiveNum /= divisor
		}
	}

	var textGroup [groupsCount]string
	for i := 0; i < groupsCount; i++ {
		textGroup[i] = digitGroup2Text(groups[i], useAnd)
	}
	combinedText := textGroup[0]
	and := useAnd && (groups[0] > 0 && groups[0] < 100)

	for i := 1; i < groupsCount; i++ {
		if groups[i] != 0 {
			prefix := textGroup[i] + " " + _scaleNum[i]

			if len(combinedText) != 0 {
				prefix += separator(and)
			}

			and = false

			combinedText = prefix + combinedText
		}
	}

	if number < 0 {
		combinedText = "minus " + combinedText
	}

	return combinedText
}

func intMod(x, y int) int {
	return int(math.Mod(float64(x), float64(y)))
}

func digitGroup2Text(group int64, useAnd bool) (ret string) {
	hundreds := group / 100
	tensUnits := intMod(int(group), 100)

	if hundreds != 0 {
		ret += _smallNumbers[hundreds] + " hundred"

		if tensUnits != 0 {
			ret += separator(useAnd)
		}
	}

	tens := tensUnits / 10
	units := intMod(tensUnits, 10)

	if tens >= 2 {
		ret += _tens[tens]

		if units != 0 {
			ret += " " + _smallNumbers[units]
		}
	} else if tensUnits != 0 {
		ret += _smallNumbers[tensUnits]
	}

	return
}

// separator returns proper separator string between
// number groups.
func separator(useAnd bool) string {
	if useAnd {
		return " and "
	}
	return " "
}
