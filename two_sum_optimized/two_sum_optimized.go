package main

import "fmt"

// Two Sum - Optimized Hash Map Solution
// Time Complexity: O(n)
// Space Complexity: O(n)
func twoSum(nums []int, target int) []int {
	// Create a map to store number -> index
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		// Check if complement exists in map
		if index, exists := numMap[complement]; exists {
			return []int{index, i}
		}

		// Store current number and its index
		numMap[num] = i
	}

	return []int{} // This won't be reached given problem constraints
}

func main() {
	// Test the optimized solution
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("Input: nums = %v, target = %d\n", nums, target)
	fmt.Printf("Output: %v\n", result)
}
