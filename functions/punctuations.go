package functions

func Punctuations(s []string) []string {
	for i, word := range s {
		if i > 0 && (word == "," || word == "." || word == "!" || word == "?" || word == ":" || word == ";") {
			s[i-1] += word
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
