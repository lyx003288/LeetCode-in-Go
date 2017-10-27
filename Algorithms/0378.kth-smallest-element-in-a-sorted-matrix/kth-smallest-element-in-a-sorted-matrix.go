package Problem0378

func kthSmallest(mat [][]int, k int) int {
	// 题目中说了 1 <= n
	n := len(mat)

	lo, hi := mat[0][0], mat[n-1][n-1]

	var mid, count, i, j int
	for lo < hi {
		mid = lo + (hi-lo)>>1
		count = 0
		j = n - 1
		for i = 0; i < n; i++ {
			for j >= 0 && mat[i][j] > mid {
				j--
			}
			count += j + 1
		}
		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}
