package functions

func ChangeA(s []string) {
	for i, word := range s {
		if i < len(s)-1 {
			nextLetter := string(s[i+1][0])
			if shouldReplaceWithAn(word, nextLetter) {
				s[i] = "an"
			}
		}
	}
}

func shouldReplaceWithAn(word, nextLetter string) bool {
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	for _, letter := range vowels {
		if (word == "a" || word == "A") && string(nextLetter[0]) == letter {
			return true
		}
	}
	return false
}
