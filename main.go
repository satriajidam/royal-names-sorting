package main

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

func main() {

}
