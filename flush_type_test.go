package main

import "testing"

func TestNewFlushTypeSuccess(t *testing.T) {
	testCase := []int{
		1,
		2,
	}
	for _, v := range testCase {
		_, err := NewFlushType(v)
		if err != nil {
			t.Fatalf("failed test %#v", err)
		}
	}
}
