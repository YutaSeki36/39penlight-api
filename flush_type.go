package main

import "fmt"

type FlushType int

const (
	NORMAL FlushType = iota + 1 // NORMAL 通常フラッシュ
	WAVE                        // WAVE ウェーブ
	BOUND                       // BOUND バウンド
)

func NewFlushType(flushType int) (FlushType, error) {
	switch flushType {
	case 1:
		return NORMAL, nil
	case 2:
		return WAVE, nil
	case 3:
		return BOUND, nil
	default:
		return 0, fmt.Errorf("不明なフラッシュタイプ. 値: %d", flushType)
	}
}

func (a FlushType) ToInt() int {
	return int(a)
}
