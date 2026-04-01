package main

import "fmt"

func ShiftArray(array []int, i int) []int {
	if len(array) == 0 {
		return []int{}
	}

	n := 1
	for n*n < len(array) {
		n++
	}

	matrix := make([][]int, n)
	for r := 0; r < n; r++ {
		matrix[r] = make([]int, n)
		for c := 0; c < n; c++ {
			matrix[r][c] = array[r*n+c]
		}
	}

	layers := n / 2
	for layer := 0; layer < layers; layer++ {
		ring := []int{}
		top, left := layer, layer
		bottom, right := n-1-layer, n-1-layer

		for c := left; c <= right; c++ {
			ring = append(ring, matrix[top][c])
		}
		for r := top + 1; r <= bottom; r++ {
			ring = append(ring, matrix[r][right])
		}
		for c := right - 1; c >= left; c-- {
			ring = append(ring, matrix[bottom][c])
		}
		for r := bottom - 1; r >= top+1; r-- {
			ring = append(ring, matrix[r][left])
		}

		ringLen := len(ring)
		shift := i % ringLen
		rotated := append(ring[ringLen-shift:], ring[:ringLen-shift]...)

		idx := 0
		for c := left; c <= right; c++ {
			matrix[top][c] = rotated[idx]
			idx++
		}
		for r := top + 1; r <= bottom; r++ {
			matrix[r][right] = rotated[idx]
			idx++
		}
		for c := right - 1; c >= left; c-- {
			matrix[bottom][c] = rotated[idx]
			idx++
		}
		for r := bottom - 1; r >= top+1; r-- {
			matrix[r][left] = rotated[idx]
			idx++
		}
	}

	result := make([]int, len(array))
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			result[r*n+c] = matrix[r][c]
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(ShiftArray(arr, 1))
}