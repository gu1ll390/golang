package mergesort

//MergeSort function"
func MergeSort(A []int64) {

	mergeSortRec(A, 0, len(A)-1, make([]int64, len(A)))
}

func mergeSortRec(A []int64, i, j int, B []int64) {
	if i < j {
		m := (i + j) / 2

		mergeSortRec(A, i, m, B)
		mergeSortRec(A, m+1, j, B)
	}

	merge(A, i, j, B)
}

func merge(A []int64, i, j int, B []int64) {
	m := (i + j) / 2

	k := i
	p := i
	q := m + 1

	for p <= m && q <= j {
		if A[p] < A[q] {
			B[k] = A[p]
			k++
			p++
		} else {
			B[k] = A[q]
			k++
			q++
		}
	}

	for p <= m {
		B[k] = A[p]
		k++
		p++
	}

	for q <= j {
		B[k] = A[q]
		k++
		q++
	}

	for k = i; k <= j; k++ {
		A[k] = B[k]
	}
}
