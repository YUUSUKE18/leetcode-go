package determinewhethermatrixcanbeobtainedbyrotatio

func findRotation(mat [][]int, target [][]int) bool {
	// 0度、90度、180度、270度の回転をチェック
	for k := 0; k < 4; k++ {
		if equal(mat, target) {
			return true
		}
		mat = rotate(mat)
	}

	return false
}

func equal(mat1, mat2 [][]int) bool {
	n := len(mat1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat1[i][j] != mat2[i][j] {
				return false
			}
		}
	}
	return true
}

// 行列を90度時計回りに回転
func rotate(mat [][]int) [][]int {
	n := len(mat)
	rotated := make([][]int, n)
	for i := range rotated {
		rotated[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[j][n-1-i] = mat[i][j]
		}
	}

	return rotated
}
