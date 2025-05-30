package roman

import "testing"

func TestRomanNumerals(t *testing.T) {
	t.Run("1 to I", func(t *testing.T) {
		got := ConvertToRoman(1)
		want := "I"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
