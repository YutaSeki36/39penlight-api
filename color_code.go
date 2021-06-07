package main

import (
	"fmt"
	"regexp"
)

type ColorCode string

func NewColorCode(code string) (*ColorCode, error) {
	reg := regexp.MustCompile(`^#([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`)
	if !reg.MatchString(code) {
		return nil, fmt.Errorf("不正な値 code: %s", code)
	}

	resp := ColorCode(code)
	return &resp, nil
}

func (c *ColorCode) ToString() string {
	return string(*c)
}
