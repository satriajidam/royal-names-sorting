package main

import "testing"

func TestRomanToInt(t *testing.T) {
	cases := map[string]int{
		"": 0, "I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10,
		"XIV": 14, "XL": 40, "XLIX": 49, "L": 50, "XC": 90,
		"XCIX": 99, "C": 100, "CD": 400, "CDL": 450, "D": 500,
		"CM": 900, "CML": 950, "CMXC": 990, "M": 1000,
	}

	for l, e := range cases {
		result := romanToInt(l)
		if result != e {
			t.Errorf("romanToInt(\"%s\") got: %v, want: %v", l, result, e)
		}
	}
}

func TestIsRomanNumeral(t *testing.T) {
	notRomanNumeral := []string{
		"", "A", "B", "E", "F", "G", "H", "J", "K", "N",
		"O", "P", "Q", "R", "S", "T", "U", "W", "Y", "Z",
	}

	for _, l := range notRomanNumeral {
		notRoman := isRomanNumeral(l)
		if notRoman {
			t.Errorf("isRomanNumeral(\"%s\") got: %v, want: %v", l, notRoman, false)
		}
	}

	romanNumeral := []string{"I", "V", "X", "L", "C", "D", "M"}

	for _, l := range romanNumeral {
		isRoman := isRomanNumeral(l)
		if !isRoman {
			t.Errorf("isRomanNumeral(\"%s\") got: %v, want: %v", l, isRoman, true)
		}
	}
}

func TestSplitRoyalName(t *testing.T) {
	type expected struct {
		name  string
		roman string
	}

	cases := map[string]expected{
		"George VI":       {"George", "VI"},
		"William II":      {"William", "II"},
		"Marco Polo XVII": {"Marco Polo", "XVII"},
		"Edward D. Teach": {"Edward D. Teach", ""},
	}

	for n, e := range cases {
		name, roman := splitRoyalName(n)
		if e.name != name || e.roman != roman {
			t.Errorf("splitRoyalName(\"%s\") got: %v, want: %v", n, expected{name, roman}, e)
		}
	}
}
