package main

import "fmt"

// Two Sum II - Input Array Is Sorted
// OPTIMAL SOLUTION: Two Pointers Approach
// Time Complexity: O(n) - single pass through the array
// Space Complexity: O(1) - only using two pointers, constant extra space
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		currentSum := numbers[left] + numbers[right]

		if currentSum == target {
			// Found the pair! Return 1-indexed positions
			return []int{left + 1, right + 1}
		} else if currentSum < target {
			// Sum is too small, move left pointer right to increase sum
			left++
		} else {
			// Sum is too large, move right pointer left to decrease sum
			right--
		}
	}

	// This should never be reached given the problem constraints
	return []int{}
}

// Alternative solution: Binary Search approach
// Time Complexity: O(n log n) - for each element, binary search for complement
// Space Complexity: O(1) - constant extra space
func twoSumBinarySearch(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		complement := target - numbers[i]
		// Binary search for complement in the remaining array
		left, right := i+1, len(numbers)-1

		for left <= right {
			mid := left + (right-left)/2
			if numbers[mid] == complement {
				return []int{i + 1, mid + 1} // 1-indexed
			} else if numbers[mid] < complement {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return []int{}
}

// Naive solution using HashMap (like original Two Sum)
// Time Complexity: O(n)
// Space Complexity: O(n) - violates the constant space requirement!
func twoSumHashMap(numbers []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range numbers {
		complement := target - num
		if index, exists := numMap[complement]; exists {
			return []int{index + 1, i + 1} // 1-indexed
		}
		numMap[num] = i
	}

	return []int{}
}

// Helper function to demonstrate how two pointers work step-by-step
func twoSumWithVisualization(numbers []int, target int) []int {
	fmt.Printf("Finding two numbers that sum to %d in array: %v\n", target, numbers)
	left, right := 0, len(numbers)-1
	step := 1

	for left < right {
		currentSum := numbers[left] + numbers[right]
		fmt.Printf("Step %d: left=%d (val=%d), right=%d (val=%d), sum=%d\n",
			step, left, numbers[left], right, numbers[right], currentSum)

		if currentSum == target {
			fmt.Printf("✅ Found target! Returning [%d, %d] (1-indexed)\n", left+1, right+1)
			return []int{left + 1, right + 1}
		} else if currentSum < target {
			fmt.Printf("   Sum %d < target %d, moving left pointer right\n", currentSum, target)
			left++
		} else {
			fmt.Printf("   Sum %d > target %d, moving right pointer left\n", currentSum, target)
			right--
		}
		step++
	}

	return []int{}
}

func main() {
	fmt.Println("=== Two Sum II - Input Array Is Sorted ===")
	fmt.Println("Given a 1-indexed sorted array, find two numbers that sum to target")
	fmt.Println("Must use constant extra space - optimal solution uses two pointers")
	fmt.Println()

	// Test cases from the problem
	testCases := []struct {
		numbers  []int
		target   int
		expected []int
		name     string
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}, "Example 1: Basic case"},
		{[]int{2, 3, 4}, 6, []int{1, 3}, "Example 2: Skip middle element"},
		{[]int{-1, 0}, -1, []int{1, 2}, "Example 3: Negative numbers"},
		{[]int{1, 2, 3, 4, 4, 9, 56, 90}, 8, []int{4, 5}, "Duplicate elements"},
		{[]int{1, 3, 4, 5, 7, 10, 11}, 9, []int{3, 4}, "Multiple valid pairs possible"},
		{[]int{-10, -5, -3, 0, 1, 3, 5, 12}, -8, []int{1, 3}, "Mixed negative/positive"},
		{[]int{1, 2}, 3, []int{1, 2}, "Minimum array size"},
		{[]int{0, 0, 3, 4}, 0, []int{1, 2}, "Zero sum target"},
	}

	fmt.Println("=== Testing Optimal Two Pointers Approach ===")
	for i, tc := range testCases {
		fmt.Printf("\n--- Test Case %d: %s ---\n", i+1, tc.name)
		fmt.Printf("Input: numbers = %v, target = %d\n", tc.numbers, tc.target)
		result := twoSum(tc.numbers, tc.target)
		fmt.Printf("Output: %v\n", result)

		// Verify the result
		if len(result) == 2 {
			idx1, idx2 := result[0]-1, result[1]-1 // Convert to 0-indexed for verification
			actualSum := tc.numbers[idx1] + tc.numbers[idx2]
			fmt.Printf("Verification: numbers[%d] + numbers[%d] = %d + %d = %d\n",
				result[0], result[1], tc.numbers[idx1], tc.numbers[idx2], actualSum)

			if actualSum == tc.target {
				fmt.Printf("✅ Correct! Expected: %v\n", tc.expected)
			} else {
				fmt.Printf("❌ Wrong sum! Expected target: %d\n", tc.target)
			}
		}
	}

	fmt.Println("\n=== Algorithm Visualization Example ===")
	fmt.Println("Let's trace through the two pointers algorithm step by step:")
	visualizeNumbers := []int{2, 7, 11, 15}
	visualizeTarget := 9
	twoSumWithVisualization(visualizeNumbers, visualizeTarget)

	fmt.Println("\n=== Comparing Different Approaches ===")
	testNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testTarget := 15

	fmt.Printf("Array: %v, Target: %d\n", testNumbers, testTarget)

	result1 := twoSum(testNumbers, testTarget)
	fmt.Printf("Two Pointers:   %v (O(n) time, O(1) space) ✅ Optimal\n", result1)

	result2 := twoSumBinarySearch(testNumbers, testTarget)
	fmt.Printf("Binary Search:  %v (O(n log n) time, O(1) space)\n", result2)

	result3 := twoSumHashMap(testNumbers, testTarget)
	fmt.Printf("HashMap:        %v (O(n) time, O(n) space) ❌ Violates constraint\n", result3)

	fmt.Println("\n=== Why Two Pointers Works for Sorted Arrays ===")
	fmt.Println("Key insights:")
	fmt.Println("1. Array is SORTED - we can use this property!")
	fmt.Println("2. If sum is too small → move left pointer right (increase sum)")
	fmt.Println("3. If sum is too large → move right pointer left (decrease sum)")
	fmt.Println("4. We'll never miss the answer because we systematically explore all possibilities")
	fmt.Println("5. Each element is considered at most once → O(n) time complexity")
	fmt.Println("6. Only using two pointer variables → O(1) space complexity")

	fmt.Println("\n=== Problem Differences from Original Two Sum ===")
	fmt.Println("Two Sum (Original)     vs     Two Sum II (This Problem)")
	fmt.Println("• Unsorted array              • SORTED array")
	fmt.Println("• 0-indexed return            • 1-indexed return")
	fmt.Println("• HashMap solution optimal    • Two pointers optimal")
	fmt.Println("• O(n) space acceptable       • O(1) space REQUIRED")
	fmt.Println("• One solution guaranteed     • One solution guaranteed")

	fmt.Println("\n=== Edge Cases Handled ===")
	fmt.Println("✅ Negative numbers")
	fmt.Println("✅ Zero as target or in array")
	fmt.Println("✅ Duplicate elements")
	fmt.Println("✅ Minimum array size (2 elements)")
	fmt.Println("✅ Large arrays")
	fmt.Println("✅ Mixed positive/negative numbers")

	fmt.Println("\n=== Time & Space Complexity Analysis ===")
	fmt.Println("OPTIMAL SOLUTION (Two Pointers):")
	fmt.Println("• Time: O(n) - each element visited at most once")
	fmt.Println("• Space: O(1) - only two pointer variables used")
	fmt.Println("• Why optimal: Can't do better than O(n) time (must examine elements)")
	fmt.Println("• Meets constraint: Constant extra space requirement")
	fmt.Println()
	fmt.Println("ALTERNATIVE SOLUTIONS:")
	fmt.Println("• Binary Search: O(n log n) time, O(1) space")
	fmt.Println("• HashMap: O(n) time, O(n) space (violates constraint)")
	fmt.Println("• Brute Force: O(n²) time, O(1) space (too slow)")
}
