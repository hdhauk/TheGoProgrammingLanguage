// Package word provides utilities for word games
package word

import "unicode"

// IsPalindrome checks if word s reads the same backward as forward
func IsPalindrome(s string) bool {
	// Clean up string. Remove punctuation and convert to lowercase
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	// Check reverse
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
