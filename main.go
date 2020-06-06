package main

import (
	"regexp"
	"strings"
)

func romanToInt(roman string) int {
	letters := []rune(roman)
	romanMap := map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	if len(letters) == 0 {
		return 0
	}

	result := 0
	temp := 0

	for i := len(letters) - 1; i >= 0; i-- {
		if temp <= romanMap[letters[i]] {
			result = result + romanMap[letters[i]]
		} else {
			result = result - romanMap[letters[i]]
		}
		temp = romanMap[letters[i]]
	}

	return result
}

func isRomanNumeral(letter string) bool {
	if letter == "" {
		return false
	}

	matched, _ := regexp.MatchString(
		`^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`,
		strings.ToUpper(letter),
	)

	return matched
}

func splitRoyalName(royalName string) (string, string) {
	if royalName == "" {
		return royalName, ""
	}

	items := strings.Split(royalName, " ")
	lastItem := items[len(items)-1]

	if !isRomanNumeral(lastItem) {
		return royalName, ""
	}

	return strings.Join(items[:len(items)-1], " "), lastItem
}

func main() {

}
