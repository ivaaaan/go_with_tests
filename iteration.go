package main

import "strings"

// Repeat returns string s repeated 5 times
func Repeat(s string, times int) string {
	var b strings.Builder
	for i := 0; i < times; i++ {
		b.WriteString(s)
	}
	return b.String()
}
