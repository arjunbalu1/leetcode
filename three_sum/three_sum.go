package main

import (
	"fmt"
	"sort"
	"strings"
)

// OPTIMAL SOLUTION: Sort + Two Pointers
// Time Complexity: O(n²) - one loop + two pointers for each element
// Space Complexity: O(1) - not counting the output array
func threeSum(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue //To prevent the repeat
		}
		target, left, right := -nums[i], i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
	return results
}

// USER'S ORIGINAL APPROACH: Brute Force with HashMap
// Time Complexity: O(n³) - three nested loops
// Space Complexity: O(n³) - potentially storing all triplets
func threeSumBruteForce(nums []int) [][]int {
	ans := [][]int{}
	seen := make(map[[3]int]bool)

	for i, num1 := range nums {
		for j, num2 := range nums {
			if j != i {
				for k, num3 := range nums {
					if k != j && k != i {
						if num1+num2+num3 == 0 {
							triplet := []int{num1, num2, num3}
							sort.Ints(triplet) // Sort to handle duplicates
							key := [3]int{triplet[0], triplet[1], triplet[2]}
							if !seen[key] {
								seen[key] = true
								ans = append(ans, triplet)
							}
						}
					}
				}
			}
		}
	}
	return ans
}

// ALTERNATIVE: Sort + HashMap (compromise between brute force and optimal)
// Time Complexity: O(n²) - for each pair, look up the third element
// Space Complexity: O(n) - for the hashmap
func threeSumHashMap(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates for first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		seen := make(map[int]int)
		target := -nums[i]

		for j := i + 1; j < len(nums); j++ {
			needed := target - nums[j]

			if idx, exists := seen[needed]; exists {
				result = append(result, []int{nums[i], nums[idx], nums[j]})

				// Skip duplicates for second and third elements
				for j+1 < len(nums) && nums[j] == nums[j+1] {
					j++
				}
			}

			// Add current element to seen map
			if _, exists := seen[nums[j]]; !exists {
				seen[nums[j]] = j
			}
		}
	}

	return result
}

// Helper function to demonstrate the optimal algorithm step by step
func threeSumWithVisualization(nums []int) [][]int {
	fmt.Printf("Finding all triplets that sum to 0 in: %v\n", nums)

	if len(nums) < 3 {
		fmt.Println("Array too small, need at least 3 elements")
		return [][]int{}
	}

	// Sort first
	sort.Ints(nums)
	fmt.Printf("After sorting: %v\n", nums)

	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates for first element
		if i > 0 && nums[i] == nums[i-1] {
			fmt.Printf("Skipping duplicate first element: %d\n", nums[i])
			continue
		}

		fmt.Printf("\n--- Iteration %d: first element = %d ---\n", i+1, nums[i])
		fmt.Printf("Looking for two numbers that sum to %d\n", -nums[i])

		left, right := i+1, len(nums)-1
		target := -nums[i]
		step := 1

		for left < right {
			currentSum := nums[left] + nums[right]
			fmt.Printf("Step %d: nums[%d]=%d + nums[%d]=%d = %d, target=%d\n",
				step, left, nums[left], right, nums[right], currentSum, target)

			if currentSum == target {
				triplet := []int{nums[i], nums[left], nums[right]}
				fmt.Printf("✅ Found triplet: %v\n", triplet)
				result = append(result, triplet)

				// Skip duplicates
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if currentSum < target {
				fmt.Printf("   Sum too small, moving left pointer right\n")
				left++
			} else {
				fmt.Printf("   Sum too large, moving right pointer left\n")
				right--
			}
			step++
		}
	}

	return result
}

func main() {
	fmt.Println("=== 3Sum Problem - Find All Triplets That Sum to Zero ===")
	fmt.Println("Given an array, find all unique triplets [a,b,c] where a+b+c=0")
	fmt.Println()

	// Test cases from the problem
	testCases := []struct {
		nums     []int
		expected [][]int
		name     string
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
			"Example 1: Mixed positive/negative",
		},
		{
			[]int{0, 1, 1},
			[][]int{},
			"Example 2: No valid triplets",
		},
		{
			[]int{0, 0, 0},
			[][]int{{0, 0, 0}},
			"Example 3: All zeros",
		},
		{
			[]int{-2, 0, 1, 1, 2},
			[][]int{{-2, 0, 2}, {-2, 1, 1}},
			"Multiple valid triplets",
		},
		{
			[]int{-1, 0, 1},
			[][]int{{-1, 0, 1}},
			"Simple case",
		},
		{
			[]int{1, 2, 3},
			[][]int{},
			"All positive numbers",
		},
		{
			[]int{-3, -2, -1},
			[][]int{},
			"All negative numbers",
		},
		{
			[]int{0, 0, 0, 0},
			[][]int{{0, 0, 0}},
			"Multiple zeros",
		},
	}

	fmt.Println("=== Testing All Approaches ===")
	for i, tc := range testCases {
		fmt.Printf("\n--- Test Case %d: %s ---\n", i+1, tc.name)
		fmt.Printf("Input: %v\n", tc.nums)

		result1 := threeSum(tc.nums)
		result2 := threeSumBruteForce(tc.nums)
		result3 := threeSumHashMap(tc.nums)

		fmt.Printf("Expected:     %v\n", tc.expected)
		fmt.Printf("Optimal:      %v\n", result1)
		fmt.Printf("Brute Force:  %v\n", result2)
		fmt.Printf("HashMap:      %v\n", result3)

		// Simple verification (order might differ)
		fmt.Printf("Results match: %v\n", len(result1) == len(result2) && len(result2) == len(result3))
	}

	fmt.Println("\n=== Algorithm Visualization ===")
	fmt.Println("Let's trace through the optimal algorithm:")
	visualNums := []int{-1, 0, 1, 2, -1, -4}
	threeSumWithVisualization(visualNums)

	fmt.Println("\n=== Your Original Approach Analysis ===")
	fmt.Println("✅ WHAT YOU DID RIGHT:")
	fmt.Println("• Correct logic with three nested loops")
	fmt.Println("• Smart duplicate handling with sorted keys")
	fmt.Println("• Proper index checking (i != j != k)")
	fmt.Println("• Correct sum validation")
	fmt.Println()
	fmt.Println("⚠️  AREAS FOR IMPROVEMENT:")
	fmt.Println("• Time complexity: O(n³) - can be optimized to O(n²)")
	fmt.Println("• Space complexity: O(n³) - can be optimized to O(1)")
	fmt.Println("• Doesn't leverage sorting for efficiency")

	fmt.Println("\n=== Algorithm Comparison ===")
	fmt.Printf("%-15s | %-12s | %-12s | %s\n", "Approach", "Time", "Space", "Notes")
	fmt.Println(strings.Repeat("-", 65))
	fmt.Printf("%-15s | %-12s | %-12s | %s\n", "Optimal (2-ptr)", "O(n²)", "O(1)", "✅ Best overall")
	fmt.Printf("%-15s | %-12s | %-12s | %s\n", "HashMap", "O(n²)", "O(n)", "Good compromise")
	fmt.Printf("%-15s | %-12s | %-12s | %s\n", "Your Brute Force", "O(n³)", "O(n³)", "Correct but slow")

	fmt.Println("\n=== Key Insights for 3Sum ===")
	fmt.Println("1. SORT FIRST: Enables two-pointer technique and easy duplicate skipping")
	fmt.Println("2. REDUCE TO 2SUM: Fix first element, find two-sum for remaining")
	fmt.Println("3. SKIP DUPLICATES: Crucial for avoiding duplicate triplets")
	fmt.Println("4. TWO POINTERS: Use sorted array property for O(n) inner loop")
	fmt.Println("5. TOTAL TIME: O(n log n) sort + O(n²) search = O(n²)")

	fmt.Println("\n=== How the Optimal Algorithm Works ===")
	fmt.Println("1. Sort the array: [-4, -1, -1, 0, 1, 2]")
	fmt.Println("2. For each element nums[i], find two elements that sum to -nums[i]")
	fmt.Println("3. Use two pointers (left=i+1, right=end) on remaining sorted array")
	fmt.Println("4. If sum too small → move left pointer right")
	fmt.Println("5. If sum too large → move right pointer left")
	fmt.Println("6. If sum equals target → found triplet, move both pointers")
	fmt.Println("7. Skip duplicates at each level to avoid duplicate triplets")

	fmt.Println("\n=== Practice Progression ===")
	fmt.Println("🎯 Your journey:")
	fmt.Println("1. ✅ Two Sum (HashMap) - Foundation")
	fmt.Println("2. ✅ Two Sum II (Two Pointers) - Sorted array optimization")
	fmt.Println("3. ✅ 3Sum (Your approach) - Brute force understanding")
	fmt.Println("4. 🎯 3Sum (Optimal) - Combine sorting + two pointers")
	fmt.Println("5. 🔜 Next: 4Sum, 3Sum Closest, etc.")
}
