package main

import "fmt"

func reverseString(s string) string {
	l := len(s)
	srt := ""
	for i := l - 1; i >= 0; i-- {
		srt += fmt.Sprintf("%c", s[i])
	}
	return srt
}
