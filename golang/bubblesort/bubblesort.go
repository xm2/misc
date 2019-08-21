package bubblesort

func Bubblesort(parr *[]int) {
	max := len(*parr)
	for max > 0 {
		swap := false
		for i := 0; i < max-1; i++ {
			if (*parr)[i] > (*parr)[i+1] {
				(*parr)[i], (*parr)[i+1] = (*parr)[i+1], (*parr)[i]
				swap = true
			}
		}
		if !swap {
			break
		}
		max--
	}
}
