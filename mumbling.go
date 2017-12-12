package main

import (
	"fmt"
	"strings"
)

// Accum codewars func
func Accum(s string) string {
	result := []string{}
	for i, c := range s {
		result = append(
			result,
			strings.Title(strings.Repeat(strings.ToLower(string(c)), i+1)),
		)
	}
	return strings.Join(result, "-")
}

func main() {
	fmt.Println(Accum("ZpglnRxqenU"))
}
