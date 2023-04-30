package main

func romanToInt(s string) int {
	romanToNumberMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result, prevValue int

	for _, char := range s {
		value := romanToNumberMap[char]
		if value > prevValue {
			result += value - 2*prevValue
		} else {
			result += value
		}
		prevValue = value
	}
	return result
}
