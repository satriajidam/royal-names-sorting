package main

import "testing"

func TestRomanNumToInt(t *testing.T) {
	tests := map[string]int{
		"": 0, "I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10,
		"XIV": 14, "XL": 40, "XLIX": 49, "L": 50, "XC": 90,
		"XCIX": 99, "C": 100, "CD": 400, "CDL": 450, "D": 500,
		"CM": 900, "CML": 950, "CMXC": 990, "M": 1000,
	}

	for letter, want := range tests {
		got := romanNumToInt(letter)
		if got != want {
			t.Errorf("romanNumToInt(\"%s\") got: %v, want: %v", letter, got, want)
		}
	}
}

func TestIsRomanNum(t *testing.T) {
	falseTests := []string{
		"", "A", "B", "E", "F", "G", "H", "J", "K", "N",
		"O", "P", "Q", "R", "S", "T", "U", "W", "Y", "Z",
	}

	for _, letter := range falseTests {
		got := isRomanNum(letter)
		if got {
			t.Errorf("isRomanNum(\"%s\") got: %v, want: %v", letter, got, false)
		}
	}

	trueTests := []string{"I", "V", "X", "L", "C", "D", "M"}

	for _, letter := range trueTests {
		got := isRomanNum(letter)
		if !got {
			t.Errorf("isRomanNum(\"%s\") got: %v, want: %v", letter, got, true)
		}
	}
}

func TestSplitRoyalName(t *testing.T) {
	tests := map[string]struct {
		name     string
		romanNum string
	}{
		"George VI":       {"George", "VI"},
		"William II":      {"William", "II"},
		"Marco Polo XVII": {"Marco Polo", "XVII"},
		"Edward Teach":    {"Edward Teach", ""},
	}

	for name, want := range tests {
		gotName, gotRomanNum := splitRoyalName(name)
		if want.name != gotName {
			t.Errorf("splitRoyalName(\"%s\") got: %v, want: %v", name, gotName, want.name)
		}
		if want.romanNum != gotRomanNum {
			t.Errorf("splitRoyalName(\"%s\") got: %v, want: %v", name, gotRomanNum, want.romanNum)
		}
	}
}

func TestSortRoyalNames(t *testing.T) {
	tests := []struct {
		name       string
		royalNames RoyalNames
		want       RoyalNames
	}{
		{
			name: "case1",
			royalNames: RoyalNames{
				"Louis IX", "Louis VIII", "David II",
			},
			want: RoyalNames{
				"David II", "Louis VIII", "Louis IX",
			},
		},
		{
			name: "case2",
			royalNames: RoyalNames{
				"George VI", "William II", "George CMXC", "Elizabeth I", "William I",
			},
			want: RoyalNames{
				"Elizabeth I", "George VI", "George CMXC", "William I", "William II",
			},
		},
		{
			name: "case3",
			royalNames: RoyalNames{
				"Marco Polo XVII", "Edward Teach", "Indiana Jones XXXI", "Hamengkubuwono IX", "Hamengkubuwono X",
			},
			want: RoyalNames{
				"Edward Teach", "Hamengkubuwono IX", "Hamengkubuwono X", "Indiana Jones XXXI", "Marco Polo XVII",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := append(RoyalNames{}, tt.royalNames...)
			got.Sort()
			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("%s.Sort(%v) got: %v, want: %v", tt.name, tt.royalNames, got, tt.want)
				}
			}
		})
	}
}
