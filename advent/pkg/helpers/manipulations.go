package helpers

// Permutations generates all possible combinations from the input data
func Permutations(xs []int16) (permuts [][]int16) {
	// Taken from: https://www.golangprograms.com/golang-program-to-generate-slice-permutations-of-number-entered-by-user.html
	var rc func([]int16, int16)
	rc = func(a []int16, k int16) {
		if k == int16(len(a)) {
			permuts = append(permuts, append([]int16{}, a...))
		} else {
			for i := k; i < int16(len(xs)); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

// IsLocationValid returns if the provided x,y coordinates are within the range of the provided 2d array.
func IsLocationValid(arr [][]rune, x, y int) bool {
	//TODO change this to be a generic 2d array rather than rune specifically
	return (x >= 0) && (x < len(arr[0])) && (y >= 0) && (y < len(arr))
}

// Abs is simple function to return the absolute value of an integer. Absolute value being essentially an always positive number.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ManhattansDistance return thes Manhattan distance between two points
func ManhattansDistance(x1, y1, x2, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}
