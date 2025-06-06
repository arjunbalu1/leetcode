package main

import "fmt"

// Two Sum - Brute Force Solution
// Time Complexity: O(n²)
// Space Complexity: O(1)
func twoSum(nums []int, target int) []int {
	// Check all pairs of numbers
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	// This should never be reached given the problem constraints
	return []int{}
}

func main() {
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
	fmt.Printf("Explanation: nums[%d] + nums[%d] = %d + %d = %d\n",
		result3[0], result3[1], nums3[result3[0]], nums3[result3[1]], target3)
}
