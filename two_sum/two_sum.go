package main

import "fmt"

// Two Sum - Optimized HashMap Solution
// Time Complexity: O(n) - single pass through the array
// Space Complexity: O(n) - for the HashMap storage
func twoSum(nums []int, target int) []int {
	// Create a map to store number -> index mapping
	numMap := make(map[int]int)

	// Single pass through the array
	for i, num := range nums {
		// Calculate what number we need to reach the target
		complement := target - num

		// Check if the complement exists in our map
		if index, exists := numMap[complement]; exists {
			// Found the pair! Return indices of complement and current number
			return []int{index, i}
		}

		// Store current number and its index for future lookups
		numMap[num] = i
	}

	// This should never be reached given the problem constraints
	return []int{}
}

// Alternative implementation: Two-pass HashMap approach
// This is slightly less efficient but more readable for learning
func twoSumTwoPass(nums []int, target int) []int {
	// First pass: build the hash map
	numMap := make(map[int]int)
	for i, num := range nums {
		numMap[num] = i
	}

	// Second pass: find the complement
	for i, num := range nums {
		complement := target - num
		if index, exists := numMap[complement]; exists && index != i {
			return []int{i, index}
		}
	}

	return []int{}
}

func main() {
	fmt.Println("=== Optimized Two Sum Solution ===")
	fmt.Println("Using HashMap approach: O(n) time, O(n) space")
	fmt.Println()

	// Test cases from the problem

	// Example 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := twoSum(nums1, target1)
	fmt.Printf("Input: nums = %v, target = %d\n", nums1, target1)
	fmt.Printf("Output: %v\n", result1)
	fmt.Printf("Explanation: nums[%d] + nums[%d] = %d + %d = %d\n\n",
		result1[0], result1[1], nums1[result1[0]], nums1[result1[1]], target1)

	// Example 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	fmt.Printf("Input: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("Output: %v\n", result2)
	fmt.Printf("Explanation: nums[%d] + nums[%d] = %d + %d = %d\n\n",
		result2[0], result2[1], nums2[result2[0]], nums2[result2[1]], target2)

	// Example 3
	nums3 := []int{3, 3}
	target3 := 6
	result3 := twoSum(nums3, target3)
	fmt.Printf("Input: nums = %v, target = %d\n", nums3, target3)
	fmt.Printf("Output: %v\n", result3)
	fmt.Printf("Explanation: nums[%d] + nums[%d] = %d + %d = %d\n\n",
		result3[0], result3[1], nums3[result3[0]], nums3[result3[1]], target3)

	// Demonstrate the alternative two-pass approach
	fmt.Println("=== Testing Two-Pass Approach ===")
	result1Alt := twoSumTwoPass(nums1, target1)
	fmt.Printf("Two-pass result for [2,7,11,15], target 9: %v\n", result1Alt)
}
