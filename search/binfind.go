package search

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {
	if len(list) == 0 {
		return -1
	}

	head := list[0]
	rest := list[1:]
	if head == x {
		return 0
	}

	pos := FindSorted(rest, x)

	if pos == -1 {
		return -1
	}
	return pos + 1
}
