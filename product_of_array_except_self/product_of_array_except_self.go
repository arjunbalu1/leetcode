package main

import "fmt"

// Product of Array Except Self - Optimized O(1) Extra Space Solution
// Time Complexity: O(n) - two passes through the array
// Space Complexity: O(1) - only using the output array (which doesn't count as extra space)
func productExceptSelf(nums []int) []int {
	count := len(nums)
	product := make([]int, count)

	// Pass 1: Fill with prefix products (left products)
	prefix := 1
	for i := 0; i < count; i++ {
		product[i] = prefix
		prefix *= nums[i]
	}

	// Pass 2: Multiply by suffix products (right products)
	suffix := 1
	for i := count - 1; i >= 0; i-- {
		product[i] *= suffix
		suffix *= nums[i]
	}

	return product
}

// Alternative implementation: Using extra space for clarity (O(n) space)
// This is easier to understand but uses O(n) extra space
func productExceptSelfWithExtraSpace(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	leftProducts := make([]int, n)
	rightProducts := make([]int, n)

	// Calculate left products
	leftProducts[0] = 1
	for i := 1; i < n; i++ {
		leftProducts[i] = leftProducts[i-1] * nums[i-1]
	}

	// Calculate right products
	rightProducts[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rightProducts[i] = rightProducts[i+1] * nums[i+1]
	}

	// Combine left and right products
	for i := 0; i < n; i++ {
		result[i] = leftProducts[i] * rightProducts[i]
	}

	return result
}

// Educational helper function to demonstrate the concept step by step
func explainProductExceptSelf(nums []int) {
	n := len(nums)
	fmt.Printf("Input array: %v\n", nums)
	fmt.Println("\nStep-by-step explanation:")

	// Show left products
	fmt.Println("\n1. Left products (product of all elements to the left):")
	leftProducts := make([]int, n)
	leftProducts[0] = 1
	fmt.Printf("   Index 0: 1 (no elements to the left)\n")

	for i := 1; i < n; i++ {
		leftProducts[i] = leftProducts[i-1] * nums[i-1]
		fmt.Printf("   Index %d: %d (product of elements at indices 0 to %d)\n",
			i, leftProducts[i], i-1)
	}
	fmt.Printf("   Left products array: %v\n", leftProducts)

	// Show right products
	fmt.Println("\n2. Right products (product of all elements to the right):")
	rightProducts := make([]int, n)
	rightProducts[n-1] = 1
	fmt.Printf("   Index %d: 1 (no elements to the right)\n", n-1)

	for i := n - 2; i >= 0; i-- {
		rightProducts[i] = rightProducts[i+1] * nums[i+1]
		fmt.Printf("   Index %d: %d (product of elements at indices %d to %d)\n",
			i, rightProducts[i], i+1, n-1)
	}
	fmt.Printf("   Right products array: %v\n", rightProducts)

	// Show final result
	fmt.Println("\n3. Final result (left[i] * right[i]):")
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = leftProducts[i] * rightProducts[i]
		fmt.Printf("   Index %d: %d * %d = %d\n",
			i, leftProducts[i], rightProducts[i], result[i])
	}
	fmt.Printf("   Final result: %v\n", result)
}

func main() {
	fmt.Println("=== Product of Array Except Self Solution ===")
	fmt.Println("Problem: Return array where each element is product of all other elements")
	fmt.Println("Constraints: O(n) time, no division operation, O(1) extra space")
	fmt.Println()

	// Test cases from the problem

	// Example 1
	fmt.Println("=== Example 1 ===")
	nums1 := []int{1, 2, 3, 4}
	fmt.Printf("Input: nums = %v\n", nums1)
	result1 := productExceptSelf(nums1)
	fmt.Printf("Output: %v\n", result1)
	fmt.Println("Expected: [24, 12, 8, 6]")
	fmt.Println("Explanation:")
	fmt.Println("  Index 0: 2*3*4 = 24")
	fmt.Println("  Index 1: 1*3*4 = 12")
	fmt.Println("  Index 2: 1*2*4 = 8")
	fmt.Println("  Index 3: 1*2*3 = 6")
	fmt.Println()

	// Example 2
	fmt.Println("=== Example 2 ===")
	nums2 := []int{-1, 1, 0, -3, 3}
	fmt.Printf("Input: nums = %v\n", nums2)
	result2 := productExceptSelf(nums2)
	fmt.Printf("Output: %v\n", result2)
	fmt.Println("Expected: [0, 0, 9, 0, 0]")
	fmt.Println("Explanation: Since there's a 0 in the array, all products")
	fmt.Println("except the one at index 2 will be 0")
	fmt.Println()

	// Additional test case with all positive numbers
	fmt.Println("=== Additional Test Case ===")
	nums3 := []int{2, 3, 4, 5}
	fmt.Printf("Input: nums = %v\n", nums3)
	result3 := productExceptSelf(nums3)
	fmt.Printf("Output: %v\n", result3)
	fmt.Println()

	// Demonstrate the step-by-step explanation
	fmt.Println("=== Detailed Step-by-Step Explanation ===")
	explainProductExceptSelf([]int{1, 2, 3, 4})

	// Compare both implementations
	fmt.Println("\n=== Comparing Both Implementations ===")
	testNums := []int{2, 3, 4, 5}
	optimized := productExceptSelf(testNums)
	withExtraSpace := productExceptSelfWithExtraSpace(testNums)

	fmt.Printf("Input: %v\n", testNums)
	fmt.Printf("Optimized (O(1) space): %v\n", optimized)
	fmt.Printf("With extra space: %v\n", withExtraSpace)
	fmt.Printf("Results match: %v\n", fmt.Sprintf("%v", optimized) == fmt.Sprintf("%v", withExtraSpace))

	fmt.Println("\n=== Algorithm Analysis ===")
	fmt.Println("Time Complexity: O(n) - we make exactly 2 passes through the array")
	fmt.Println("Space Complexity: O(1) - only using the output array (doesn't count as extra space)")
	fmt.Println("Key Insight: result[i] = (product of all elements left of i) × (product of all elements right of i)")
	fmt.Println("Optimization: Use the result array to store left products, then multiply with right products in-place")
}
