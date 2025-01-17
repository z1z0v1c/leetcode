/**
 * 74 - Medium
 * 
 * You are given an m x n integer matrix matrix with the following two properties:
 * 
 *   - Each row is sorted in non-decreasing order.
 *   - The first integer of each row is greater than the last integer of the previous row.
 * 
 * Given an integer target, return true if target is in matrix or false otherwise.
 * You must write a solution in O(log(m * n)) time complexity.
 * 
 * Example 1:
 * 
 * 	Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * 	Output: true
 * 
 * Example 2:
 * 
 * 	Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * 	Output: false
 * 
 * Constraints:
 * 
 *   - m == matrix.length
 *   - n == matrix[i].length
 *   - 1 <= m, n <= 100
 *   - -10^4 <= matrix[i][j], target <= 10^4
*/
package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	start := 0
	end := len(matrix) - 1

	if matrix[start][0] > target {
		return false
	}
	if matrix[end][len(matrix[end])-1] < target {
		return false
	}

	for start <= end {
		mid := (start + end) / 2

		if matrix[mid][len(matrix[end])-1] < target {
			start = mid + 1
		} else if matrix[mid][0] > target {
			end = mid - 1
		} else {
			return searchRow(matrix[mid], target)
		}
	}

	return false
}

func searchRow(row []int, target int) bool {
	start := 0
	end := len(row) - 1

	for start <= end {
		mid := (start + end) / 2

		if (row[mid] < target) {
			start = mid + 1
		} else if (row[mid] > target) {
			end = mid - 1
		} else {
			return true
		}
	}

	return false
}