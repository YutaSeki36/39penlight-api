package main

import "testing"

func TestNewColorCode(t *testing.T) {
	testCases := []string{
		"#c71585",
		"#ffffff",
		"#008b8b",
	}

	for _, u := range testCases {
		_, err := NewColorCode(u)
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
	}
}

func TestNewColorCodeError(t *testing.T) {
	testCases := []string{
		"#ffff",
		"",
		"###",
		"12340",
	}

	for _, u := range testCases {
		_, err := NewColorCode(u)
		if err == nil {
			t.Fatalf("failed test %#v", "エラーが発生しませんでした．")
		}
	}
}
