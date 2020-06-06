package main

import "testing"

func TestRomanToInt(t *testing.T) {
	result := romanToInt("")
	if result != 0 {
		t.Errorf("romanToInt(\"\") got: %v, want: %v", result, 0)
	}

	result = romanToInt("I")
	if result != 1 {
		t.Errorf("romanToInt(\"I\") got: %v, want: %v", result, 1)
	}

	result = romanToInt("IV")
	if result != 4 {
		t.Errorf("romanToInt(\"IV\") got: %v, want: %v", result, 4)
	}

	result = romanToInt("V")
	if result != 5 {
		t.Errorf("romanToInt(\"V\") got: %v, want: %v", result, 5)
	}

	result = romanToInt("IX")
	if result != 9 {
		t.Errorf("romanToInt(\"IX\") got: %v, want: %v", result, 9)
	}

	result = romanToInt("X")
	if result != 10 {
		t.Errorf("romanToInt(\"X\") got: %v, want: %v", result, 10)
	}

	result = romanToInt("XIV")
	if result != 14 {
		t.Errorf("romanToInt(\"XIV\") got: %v, want: %v", result, 14)
	}

	result = romanToInt("XL")
	if result != 40 {
		t.Errorf("romanToInt(\"XL\") got: %v, want: %v", result, 40)
	}

	result = romanToInt("XLIX")
	if result != 49 {
		t.Errorf("romanToInt(\"XLIX\") got: %v, want: %v", result, 49)
	}

	result = romanToInt("L")
	if result != 50 {
		t.Errorf("romanToInt(\"L\") got: %v, want: %v", result, 50)
	}

	result = romanToInt("XC")
	if result != 90 {
		t.Errorf("romanToInt(\"XC\") got: %v, want: %v", result, 90)
	}

	result = romanToInt("XCIX")
	if result != 99 {
		t.Errorf("romanToInt(\"XCIX\") got: %v, want: %v", result, 99)
	}

	result = romanToInt("C")
	if result != 100 {
		t.Errorf("romanToInt(\"C\") got: %v, want: %v", result, 100)
	}

	result = romanToInt("CD")
	if result != 400 {
		t.Errorf("romanToInt(\"CD\") got: %v, want: %v", result, 400)
	}

	result = romanToInt("CDL")
	if result != 450 {
		t.Errorf("romanToInt(\"CDL\") got: %v, want: %v", result, 450)
	}

	result = romanToInt("D")
	if result != 500 {
		t.Errorf("romanToInt(\"D\") got: %v, want: %v", result, 500)
	}

	result = romanToInt("CM")
	if result != 900 {
		t.Errorf("romanToInt(\"CM\") got: %v, want: %v", result, 900)
	}

	result = romanToInt("CML")
	if result != 950 {
		t.Errorf("romanToInt(\"CML\") got: %v, want: %v", result, 950)
	}

	result = romanToInt("CMXC")
	if result != 990 {
		t.Errorf("romanToInt(\"CMXC\") got: %v, want: %v", result, 990)
	}

	result = romanToInt("M")
	if result != 1000 {
		t.Errorf("romanToInt(\"M\") got: %v, want: %v", result, 1000)
	}
}
