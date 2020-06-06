package main

import (
	"fmt"
	"regexp"
	"sort"
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

func sortRoyalNames(royalNames []string) []string {
	if len(royalNames) == 0 {
		return []string{}
	}

	copiedNames := append([]string{}, royalNames...)

	sort.SliceStable(copiedNames, func(i, j int) bool {
		prevName, prevRoman := splitRoyalName(copiedNames[i])
		prevOrder := romanToInt(prevRoman)
		nextName, nextRoman := splitRoyalName(copiedNames[j])
		nextOrder := romanToInt(nextRoman)

		compareName := strings.Compare(prevName, nextName)

		if compareName < 0 {
			return true
		} else if compareName == 0 {
			if prevOrder < nextOrder {
				return true
			}
		}

		return false
	})

	return copiedNames
}

func main() {
	case1 := []string{"George VI", "William II", "George CMXC", "Elizabeth I", "William I"}
	fmt.Println("Before:", case1)
	result1 := sortRoyalNames(case1)
	fmt.Println("Before:", case1)
	fmt.Println("After:", result1)
}
