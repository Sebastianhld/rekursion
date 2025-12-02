package calc

// Liefert das Produkt von x und y.
func Mult(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}
	if x == 1 {
		return y
	}
	if y == 1 {
		return x
	}

	return x + (Mult(x, y-1))
}
