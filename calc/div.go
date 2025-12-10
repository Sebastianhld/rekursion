package calc

func Div(x, y float64) float64 {

	if y == 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return (x / y)
}
func DivWithRest(x, y int) (int, int) {
	if y == 0 {
		return -1, -1 // Fehlerfall
	}

	if x < y {
		return 0, x // Quotient 0, Rest x
	}

	q, r := DivWithRest(x-y, y) // Rekursion

	return q + 1, r // bei jedem Schritt 1 zum Quotienten addieren
}
