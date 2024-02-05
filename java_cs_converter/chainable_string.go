package main

import (
	"strings"
)

type ChainableString string

func (s ChainableString) Replace(old string, new string) ChainableString {
	return ChainableString(strings.ReplaceAll(string(s), old, new))
}
