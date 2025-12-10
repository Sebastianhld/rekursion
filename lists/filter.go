package lists

// Liefert eine Liste mit allen Elementen aus list, die kleiner oder gleich key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterLess(list []int, key int) []int {
	if len(list) == 0 {
		return []int{}

	}
	head := list[0]
	rest := list[1:]

	if head <= key {
		return append([]int{head}, FilterLess(rest, key)...)
	}
	return FilterLess(rest, key)
}

// Liefert eine Liste mit allen Elementen aus list, die echt größer als key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterGreater(list []int, key int) []int {
	// Gehen Sie analog zu FilterLess vor.
	if Empty(list) {
		return []int{}
	}
	head := list[0]
	rest := list[1:]
	if head > key {
		return append([]int{head}, FilterGreater(rest, key)...)
	}
	return FilterGreater(rest, key)
}
