package qsort

func Qsort(parr *[]int, min, max int) {
	if min >= max {
		return
	}
	pivot := partition(parr, min, max)

	Qsort(parr, min, pivot-1)
	Qsort(parr, pivot+1, max)
}

func partition(parr *[]int, min, max int) int {
	p := (*parr)[max]

	i := min
	for j := min; j < max; j++ {
		if (*parr)[j] < p {
			// swap
			(*parr)[i], (*parr)[j] = (*parr)[j], (*parr)[i]
			i++
		}
	}

	(*parr)[i], (*parr)[max] = (*parr)[max], (*parr)[i]

	return i
}

func Qsort2(parr *[]int, min, max int) {
	if min >= max {
		return
	}
	pivot := partition2(parr, min, max)

	Qsort2(parr, min, pivot-1)
	Qsort2(parr, pivot+1, max)
}

func partition2(parr *[]int, min, max int) int {
	p := (*parr)[max]

	i := min
	for j := max - 1; j >= i; {
		if (*parr)[i] > p {
			// swap
			(*parr)[i], (*parr)[j] = (*parr)[j], (*parr)[i]
			j--
		} else {
			i++
		}
	}

	(*parr)[i], (*parr)[max] = (*parr)[max], (*parr)[i]

	return i
}
