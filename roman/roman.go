package roman

import "strings"

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for range arabic {
		result.WriteString("I")
	}

	return result.String()
}
