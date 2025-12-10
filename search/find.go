package search

// Find sucht in einer Liste nach dem ersten Vorkommen von x
// und gibt dessen Index zurück. Falls x nicht gefunden wird,
// wird -1 zurückgegeben.
func Find(list []int, x int) int {
	if len(list) == 0 {
		return -1
	}

	if list[0] == x {
		return 0
	}
	rest := Find(list[1:], x)
	if rest == -1 {
		return -1
	}
	return rest + 1
}

// 	for i := 0; i < len(list); i++ {
// 		if list[i] == x {
// 			return i
// 		}
// 	}
// 	return -1
// }
