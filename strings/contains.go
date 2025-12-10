package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {
	if len(s) < len(seq) {
		return false
	}
	if s[:len(seq)] == seq {
		return true
	}
	return Contains(s[1:], seq)
}

//Prüft, ob seq ein Buchstabe von s ist
func ContainsOne(s, seq string) bool {
	if s == "" {
		return false
	}
	if string(s[0]) == seq {
		return true
	}
	return Contains(s[1:], seq)
}
