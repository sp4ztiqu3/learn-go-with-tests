package iteration

import "strings"

func Repeat(char string, count int) string {
	var repeated strings.Builder
	for range count {
		repeated.WriteString(char)
	}
	return repeated.String()
}
