package main

import (
	"fmt"
	"testing"
)

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

func TestNewFlushTypeFailed(t *testing.T) {
	testCase := []int{
		0,
		-1,
		3,
		100,
	}
	for _, v := range testCase {
		_, err := NewFlushType(v)
		if err == nil {
			t.Fatalf("failed test %#v", "エラーが発生しませんでした．")
		}
		if err.Error() != fmt.Sprintf("不明なフラッシュタイプ. 値: %d", v) {
			t.Fatalf("failed test %#v", "想定するエラーではありません．")
		}
	}
}
