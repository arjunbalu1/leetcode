package main

import "fmt"

// Valid Sudoku - Optimized Boolean Array Solution
// Time Complexity: O(1) - since board is always 9x9, we're doing constant work
// Space Complexity: O(1) - using fixed-size boolean arrays
func isValidSudoku(board [][]byte) bool {
	// Use 2D boolean arrays for tracking digits 1-9 in each constraint
	// Each array is [9][9] where first index = row/col/box, second index = digit (0-8 for digits 1-9)
	var rows, columns, squares [9][9]bool

	// Iterate through each cell in the 9x9 board
	for i, row := range board {
		for j, v := range row {
			// Skip empty cells
			if v != '.' {
				// Convert ASCII digit to array index: '1'->0, '2'->1, ..., '9'->8
				// ASCII value of '1' is 49, so subtract 49 to get 0-based index
				k := int(v) - 49

				// Check if digit already exists in current row, column, or 3x3 square
				// i/3*3 + j/3 calculates which 3x3 square this cell belongs to (0-8)
				if rows[i][k] || columns[j][k] || squares[i/3*3+j/3][k] {
					return false // Duplicate found
				}

				// Mark this digit as seen in the respective row, column, and square
				rows[i][k], columns[j][k], squares[i/3*3+j/3][k] = true, true, true
			}
		}
	}

	// No duplicates found, board is valid
	return true
}

// Alternative implementation using string concatenation for tracking
// This approach is more memory-efficient but slightly less readable
func isValidSudokuAlternative(board [][]byte) bool {
	seen := make(map[string]bool)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			cell := board[row][col]
			if cell == '.' {
				continue
			}

			// Create unique identifiers for row, column, and box
			rowKey := fmt.Sprintf("row%d-%c", row, cell)
			colKey := fmt.Sprintf("col%d-%c", col, cell)
			boxKey := fmt.Sprintf("box%d-%c", (row/3)*3+(col/3), cell)

			// Check if any identifier already exists
			if seen[rowKey] || seen[colKey] || seen[boxKey] {
				return false
			}

			// Mark as seen
			seen[rowKey] = true
			seen[colKey] = true
			seen[boxKey] = true
		}
	}

	return true
}

// Helper function to convert string board to byte board for testing
func stringToByte(board [][]string) [][]byte {
	result := make([][]byte, len(board))
	for i, row := range board {
		result[i] = make([]byte, len(row))
		for j, cell := range row {
			result[i][j] = cell[0]
		}
	}
	return result
}

// Helper function to print board nicely
func printBoard(board [][]string) {
	fmt.Println("Board:")
	for i, row := range board {
		if i%3 == 0 && i > 0 {
			fmt.Println("------+-------+------")
		}
		for j, cell := range row {
			if j%3 == 0 && j > 0 {
				fmt.Print("| ")
			}
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	fmt.Println("=== Valid Sudoku Solution ===")
	fmt.Println("Checking if a 9x9 Sudoku board is valid")
	fmt.Println("Rules: No duplicates in rows, columns, or 3x3 boxes")
	fmt.Println()

	// Test Case 1: Valid board
	board1 := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}

	fmt.Println("=== Test Case 1: Valid Board ===")
	printBoard(board1)
	result1 := isValidSudoku(stringToByte(board1))
	fmt.Printf("Result: %v (Expected: true)\n", result1)
	fmt.Println()

	// Test Case 2: Invalid board (duplicate 8 in top-left 3x3 box)
	board2 := [][]string{
		{"8", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}

	fmt.Println("=== Test Case 2: Invalid Board ===")
	printBoard(board2)
	result2 := isValidSudoku(stringToByte(board2))
	fmt.Printf("Result: %v (Expected: false)\n", result2)
	fmt.Printf("Explanation: Two 8's in the top-left 3x3 box\n")
	fmt.Println()

	// Test Case 3: Invalid board (duplicate in row)
	board3 := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "5"}, // Two 5's in first row
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}

	fmt.Println("=== Test Case 3: Invalid Board (Row Duplicate) ===")
	printBoard(board3)
	result3 := isValidSudoku(stringToByte(board3))
	fmt.Printf("Result: %v (Expected: false)\n", result3)
	fmt.Printf("Explanation: Two 5's in the first row\n")
	fmt.Println()

	// Test alternative implementation
	fmt.Println("=== Testing Alternative Implementation ===")
	alt1 := isValidSudokuAlternative(stringToByte(board1))
	alt2 := isValidSudokuAlternative(stringToByte(board2))
	fmt.Printf("Alternative method - Valid board: %v\n", alt1)
	fmt.Printf("Alternative method - Invalid board: %v\n", alt2)
	fmt.Println()

	// Educational explanation
	fmt.Println("=== Algorithm Explanation (Optimized Version) ===")
	fmt.Println("1. We use three 2D boolean arrays [9][9]: rows, columns, squares")
	fmt.Println("   - First index: row/column/square number (0-8)")
	fmt.Println("   - Second index: digit value (0-8 for digits 1-9)")
	fmt.Println("2. For each non-empty cell:")
	fmt.Println("   - Convert ASCII digit to index: int(v) - 49")
	fmt.Println("   - Check if digit exists in current row, column, or 3x3 square")
	fmt.Println("   - 3x3 square calculated as: i/3*3 + j/3")
	fmt.Println("3. If any duplicate found, return false")
	fmt.Println("4. Mark digit as seen using: rows[i][k] = true")
	fmt.Println()
	fmt.Println("Key Optimizations:")
	fmt.Println("• Boolean arrays instead of maps (faster access)")
	fmt.Println("• ASCII-to-index conversion: '1'-'9' → 0-8")
	fmt.Println("• Multiple assignment: a, b, c = true, true, true")
	fmt.Println()
	fmt.Println("Time Complexity: O(1) - always processing 81 cells")
	fmt.Println("Space Complexity: O(1) - fixed 3×9×9 boolean arrays")
}
