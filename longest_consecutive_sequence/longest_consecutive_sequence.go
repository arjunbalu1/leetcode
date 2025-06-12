package main

import (
	"fmt"
	"sort"
	"strings"
)

// Approach 1: Using HashSet (OPTIMAL - O(n) time)
// Time: O(n), Space: O(n)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Create a set for O(1) lookup
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 0

	for num := range numSet {
		// Only start counting if this is the beginning of a sequence
		// (i.e., num-1 is not in the set)
		if !numSet[num-1] {
			currentNum := num
			currentLength := 1

			// Keep extending the sequence
			for numSet[currentNum+1] {
				currentNum++
				currentLength++
			}

			// Update max length if current sequence is longer
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}

// Approach 2: Using Sorting (Not optimal but easier to understand)
// Time: O(n log n), Space: O(1)
func longestConsecutiveSort(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Sort the array
	sort.Ints(nums)

	maxLength := 1
	currentLength := 1

	for i := 1; i < len(nums); i++ {
		// Skip duplicates
		if nums[i] == nums[i-1] {
			continue
		}
		// If consecutive, extend current sequence
		if nums[i] == nums[i-1]+1 {
			currentLength++
		} else {
			// Sequence broken, update max and reset current
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 1
		}
	}

	// Don't forget to check the last sequence
	if currentLength > maxLength {
		maxLength = currentLength
	}

	return maxLength
}

// Approach 3: Using Union-Find (Advanced approach)
// Time: O(n), Space: O(n)
type UnionFind struct {
	parent map[int]int
	size   map[int]int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[int]int),
		size:   make(map[int]int),
	}
}

func (uf *UnionFind) Add(x int) {
	if _, exists := uf.parent[x]; !exists {
		uf.parent[x] = x
		uf.size[x] = 1
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX != rootY {
		// Union by size
		if uf.size[rootX] < uf.size[rootY] {
			rootX, rootY = rootY, rootX
		}
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	}
}

func (uf *UnionFind) GetSize(x int) int {
	return uf.size[uf.Find(x)]
}

func longestConsecutiveUnionFind(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	uf := NewUnionFind()

	// Add all numbers and union consecutive ones
	for _, num := range nums {
		uf.Add(num)
	}

	for _, num := range nums {
		if _, exists := uf.parent[num+1]; exists {
			uf.Union(num, num+1)
		}
	}

	// Find the maximum component size
	maxLength := 0
	for _, num := range nums {
		size := uf.GetSize(num)
		if size > maxLength {
			maxLength = size
		}
	}

	return maxLength
}

func main() {
	// Test cases
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},              // [1,2,3,4]
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},      // [0,1,2,3,4,5,6,7,8]
		{[]int{1, 0, 1, 2}, 3},                        // [0,1,2]
		{[]int{}, 0},                                  // Empty array
		{[]int{1}, 1},                                 // Single element
		{[]int{1, 2, 0, 1}, 3},                        // Duplicates: [0,1,2]
		{[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7}, // [-1,0,1,3,4,5,6] and others
	}

	fmt.Println("=== Longest Consecutive Sequence Solutions ===")

	for i, tc := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("nums = %v\n", tc.nums)

		result1 := longestConsecutive(tc.nums)
		result2 := longestConsecutiveSort(tc.nums)
		result3 := longestConsecutiveUnionFind(tc.nums)

		fmt.Printf("Expected: %d\n", tc.expected)
		fmt.Printf("HashSet approach: %d\n", result1)
		fmt.Printf("Sorting approach: %d\n", result2)
		fmt.Printf("Union-Find approach: %d\n", result3)

		if result1 == tc.expected && result2 == tc.expected && result3 == tc.expected {
			fmt.Println("✅ PASSED")
		} else {
			fmt.Println("❌ FAILED")
		}
		fmt.Println(strings.Repeat("-", 40))
	}

	// Algorithm explanation
	fmt.Println("\n=== Algorithm Explanation ===")
	fmt.Println("1. HashSet Approach (OPTIMAL - Recommended for LeetCode):")
	fmt.Println("   - Time: O(n), Space: O(n)")
	fmt.Println("   - Key insight: Only start counting from sequence beginnings")
	fmt.Println("   - Check if num-1 exists; if not, start counting from num")
	fmt.Println("   - Each number is visited at most twice")

	fmt.Println("\n2. Sorting Approach:")
	fmt.Println("   - Time: O(n log n), Space: O(1)")
	fmt.Println("   - Sort array and count consecutive sequences")
	fmt.Println("   - Handle duplicates by skipping them")
	fmt.Println("   - Easier to understand but not O(n)")

	fmt.Println("\n3. Union-Find Approach (Advanced):")
	fmt.Println("   - Time: O(n), Space: O(n)")
	fmt.Println("   - Union consecutive numbers into components")
	fmt.Println("   - Find the largest component size")
	fmt.Println("   - Demonstrates advanced data structure usage")

	fmt.Println("\n=== Key Insights ===")
	fmt.Println("• The HashSet approach is optimal because we only start")
	fmt.Println("  counting from the beginning of each sequence")
	fmt.Println("• This ensures each number is processed at most twice")
	fmt.Println("• Time complexity: O(n) despite nested loops!")
}
