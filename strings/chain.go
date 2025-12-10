package strings

// Chain erwartet einen String und hängt ihn n mal hintereinander.
// Liefert das Ergebnis.
func Chain(s string, n int) string {
	if n <= 0 {
		return ""
	}
	return s + Chain(s, n-1)
}
