package lists

// Liefert zwei Listen:
// - Eine, die alle Elemente aus list enthält, die kleiner oder gleich key sind.
// - Eine, die alle übrigen Elemente aus list enthält.
func Partition(list []int, key int) ([]int, []int) {
	if len(list) == 0 {
		return []int{}, []int{}
	}
	// Verwende Kopien von list, damit die ursprüngliche Liste nicht verändert wird.

	first := list[0]
	rest := list[1:]

	smallerRest, biggerRest := Partition(rest, key)

	if first <= key {
		return append([]int{first}, smallerRest...), biggerRest
	}

	return smallerRest, append([]int{first}, biggerRest...)
}
