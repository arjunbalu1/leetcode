package main

import "fmt"

// Contains Duplicate - Hash Map Solution
// Time Complexity: O(n)
// Space Complexity: O(n)
func containsDuplicate(nums []int) bool {
	// Create a map to track seen numbers
	seen := make(map[int]bool)

	for _, num := range nums {
		// If we've seen this number before, it's a duplicate
		if seen[num] {
			return true
		}
		// Mark this number as seen
		seen[num] = true
	}

	// No duplicates found
	return false
}

func main() {
	// Test cases
	test1 := []int{1, 2, 3, 1}
	test2 := []int{1, 2, 3, 4}
	test3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	fmt.Printf("nums = %v, contains duplicate: %v\n", test1, containsDuplicate(test1))
	fmt.Printf("nums = %v, contains duplicate: %v\n", test2, containsDuplicate(test2))
	fmt.Printf("nums = %v, contains duplicate: %v\n", test3, containsDuplicate(test3))
}
