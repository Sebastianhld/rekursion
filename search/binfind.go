package search

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {

	var rec func(left, right int) int

	rec = func(left, right int) int {
		if left > right {
			return -1
		}

		mid := (left + right) / 2

		if list[mid] == x {
			return mid
		}

		if list[mid] < x {
			return rec(mid+1, right)
		}
		return rec(left, mid-1)
	}

	return rec(0, len(list)-1)
}
