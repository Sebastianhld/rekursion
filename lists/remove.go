package lists

// Liefert eine Liste, die alle Elemente aus list enthält,
// außer dem an Stelle pos.
// Wenn pos außerhalb des Bereichs der Liste liegt, wird die
// ursprüngliche Liste zurückgegeben.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func RemoveElement(list []int, pos int) []int {
	if Empty(list) {
		return []int{}
	}
	first := list[0]
	rest := list[1:]

	new := RemoveElement(rest, pos)

	if first != pos {
		return append([]int{first}, new...)
	}
	return new

}
