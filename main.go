package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type RoyalNames []string

func (r RoyalNames) romanNumToInt(romanNum string) int {
	letters := []rune(romanNum)
	romanNumMap := map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	if len(letters) == 0 {
		return 0
	}

	result := 0
	temp := 0

	for i := len(letters) - 1; i >= 0; i-- {
		if temp <= romanNumMap[letters[i]] {
			result = result + romanNumMap[letters[i]]
		} else {
			result = result - romanNumMap[letters[i]]
		}
		temp = romanNumMap[letters[i]]
	}

	return result
}

func (r RoyalNames) isRomanNum(letter string) bool {
	if letter == "" {
		return false
	}

	matched, _ := regexp.MatchString(
		`^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`,
		strings.ToUpper(letter),
	)

	return matched
}

func (r RoyalNames) getRomanNum(royalName string) string {
	if royalName == "" {
		return ""
	}

	items := strings.Split(royalName, " ")
	lastItem := items[len(items)-1]

	if !r.isRomanNum(lastItem) {
		return ""
	}

	return lastItem
}

func (r RoyalNames) Less(i, j int) bool {
	compareName := strings.Compare(r[i], r[j])

	if compareName < 0 {
		return true
	} else if compareName == 0 {
		if r.romanNumToInt(r.getRomanNum(r[i])) <
			r.romanNumToInt(r.getRomanNum(r[j])) {
			return true
		}
	}

	return false
}

func (r RoyalNames) Len() int {
	return len(r)
}

func (r RoyalNames) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r RoyalNames) Sort() {
	if len(r) == 0 {
		return
	}
	sort.Sort(r)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Error: no input & output file path provided")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: no input file path provided")
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Error: no provide output file provided")
		os.Exit(1)
	}

	inputFilepath := os.Args[1]
	inputFile, err := os.Open(inputFilepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed reading input file: %s\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	names := []string{}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed reading input file: %s\n", err)
		os.Exit(1)
	}

	royalNames := RoyalNames(names)
	royalNames.Sort()

	outputFilepath := os.Args[2]
	if _, err := os.Stat(outputFilepath); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: failed writing output file: create %s: file already exists\n", outputFilepath)
		os.Exit(1)
	}

	outputFile, err := os.OpenFile(outputFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed writing output file: %s\n", err)
		os.Exit(1)
	}

	writer := bufio.NewWriter(outputFile)
	for _, name := range royalNames {
		_, err := writer.WriteString(name + "\n")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed writing to output file: %s\n", err)
			os.Exit(1)
		}
	}

	writer.Flush()
	outputFile.Close()
}
