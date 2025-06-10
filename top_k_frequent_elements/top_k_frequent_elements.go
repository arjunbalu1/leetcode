package main

import (
	"fmt"
	"sort"
)

// Top K Frequent Elements - Multiple Solution Approaches

// Approach 1: Hash Map + Bucket Sort (Optimal Solution) - User's Optimized Version
// Time Complexity: O(n) where n is array size
// Space Complexity: O(n) for the hash map and buckets
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	bucket := make([][]int, len(nums)+1)
	for num, frequency := range freq {
		bucket[frequency] = append(bucket[frequency], num)
	}

	ans := []int{}
	for i := len(bucket) - 1; k > 0; i-- {
		for _, num := range bucket[i] {
			ans = append(ans, num)
			k--
		}
	}

	return ans
}

// Approach 2: Hash Map + Sorting (Basic approach - O(n log n))
// This doesn't meet the follow-up requirement but good for understanding
func topKFrequentSorting(nums []int, k int) []int {
	// Step 1: Count frequencies
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: Create slice of (number, frequency) pairs
	type NumFreq struct {
		num  int
		freq int
	}

	pairs := make([]NumFreq, 0, len(freqMap))
	for num, freq := range freqMap {
		pairs = append(pairs, NumFreq{num, freq})
	}

	// Step 3: Sort by frequency in descending order
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	// Step 4: Extract top k elements
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = pairs[i].num
	}

	return result
}

// Approach 3: User's Cleaner Implementation - Sort Keys by Frequency
// Time Complexity: O(n log n) - sorting dominates
// Space Complexity: O(n) for frequency map and keys slice
func topKFrequentUserVersion(nums []int, k int) []int {
	// Step 1: Count frequencies
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: Extract all unique numbers (keys from the map)
	keys := []int{}
	for key := range freqMap {
		keys = append(keys, key)
	}

	// Step 3: Sort keys by their frequency in descending order
	sort.Slice(keys, func(i, j int) bool {
		return freqMap[keys[i]] > freqMap[keys[j]]
	})

	// Step 4: Return the first k elements
	return keys[:k]
}

func main() {
	fmt.Println("=== Top K Frequent Elements - Multiple Solutions ===")
	fmt.Println()

	// Test cases from the problem
	testCases := []struct {
		nums []int
		k    int
		desc string
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, "Example 1: [1,1,1,2,2,3], k=2"},
		{[]int{1}, 1, "Example 2: [1], k=1"},
		{[]int{4, 1, -1, 2, -1, 2, 3}, 2, "Custom: [4,1,-1,2,-1,2,3], k=2"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, "All unique elements, k=3"},
	}

	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.desc)
		fmt.Printf("Input: nums = %v, k = %d\n", tc.nums, tc.k)

		// Test both approaches
		result1 := topKFrequent(tc.nums, tc.k)
		result2 := topKFrequentSorting(tc.nums, tc.k)
		result3 := topKFrequentUserVersion(tc.nums, tc.k)

		fmt.Printf("Bucket Sort (O(n)):            %v\n", result1)
		fmt.Printf("Sorting Solution (O(n log n)): %v\n", result2)
		fmt.Printf("User's Corrected Implementation (O(n log n)): %v\n", result3)
		fmt.Println()
	}

	// Educational section about algorithm complexities
	fmt.Println("=== Algorithm Analysis ===")
	fmt.Println("1. Bucket Sort Solution (Optimal!):")
	fmt.Println("   - Time:  O(n) - meets follow-up requirement!")
	fmt.Println("   - Space: O(n) for hash map and buckets")
	fmt.Println("   - Best overall approach, most efficient")
	fmt.Println()

	fmt.Println("2. Sorting Solution:")
	fmt.Println("   - Time:  O(n log n) - doesn't meet follow-up requirement")
	fmt.Println("   - Space: O(n) for storing frequency pairs")
	fmt.Println("   - Simple but not optimal for large inputs")
}
