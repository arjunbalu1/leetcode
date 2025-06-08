package main

import (
	"fmt"
	"sort"
	"strings"
)

// Approach 1: Using array for character frequency count (FASTEST & MOST EFFICIENT)
// Time: O(n), Space: O(1) - using fixed array of size 26
func isAnagram(s string, t string) bool {
	// If lengths are different, they can't be anagrams
	if len(s) != len(t) {
		return false
	}

	// Array to count frequency of each letter (a-z)
	// Index 0 = 'a', Index 1 = 'b', ..., Index 25 = 'z'
	count := [26]int{}

	// Single pass through both strings simultaneously
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++ // Increment count for character in s
		count[t[i]-'a']-- // Decrement count for character in t
	}

	// If strings are anagrams, all counts should be zero
	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}

// Approach 2: Using HashMap (Alternative approach)
// Time: O(n), Space: O(1) - since we only have 26 lowercase letters
func isAnagramHashMap(s string, t string) bool {
	// If lengths are different, they can't be anagrams
	if len(s) != len(t) {
		return false
	}

	countS, countT := make(map[rune]int), make(map[rune]int)
	for i, ch := range s {
		countS[ch]++
		countT[rune(t[i])]++
	}

	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	return true
}

// Approach 3: Using sorting (less efficient but simpler)
// Time: O(n log n), Space: O(n)
func isAnagramSort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Convert strings to slices, sort them, and compare
	sRunes := []rune(s)
	tRunes := []rune(t)

	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] < sRunes[j]
	})

	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] < tRunes[j]
	})

	return string(sRunes) == string(tRunes)
}

func main() {
	// Test cases
	testCases := []struct {
		s, t     string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"listen", "silent", true},
		{"hello", "bello", false},
		{"", "", true},
		{"a", "a", true},
		{"ab", "ba", true},
		{"abc", "def", false},
	}

	fmt.Println("=== Valid Anagram Solutions ===")

	for i, tc := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("s = \"%s\", t = \"%s\"\n", tc.s, tc.t)

		result1 := isAnagram(tc.s, tc.t)
		result2 := isAnagramHashMap(tc.s, tc.t)
		result3 := isAnagramSort(tc.s, tc.t)

		fmt.Printf("Expected: %t\n", tc.expected)
		fmt.Printf("Array approach: %t\n", result1)
		fmt.Printf("HashMap approach: %t\n", result2)
		fmt.Printf("Sorting approach: %t\n", result3)

		if result1 == tc.expected && result2 == tc.expected && result3 == tc.expected {
			fmt.Println("✅ PASSED")
		} else {
			fmt.Println("❌ FAILED")
		}
		fmt.Println(strings.Repeat("-", 30))
	}

	// Performance explanation
	fmt.Println("\n=== Algorithm Explanation ===")
	fmt.Println("1. Array Approach (FASTEST - Recommended for LeetCode):")
	fmt.Println("   - Time: O(n), Space: O(1)")
	fmt.Println("   - Uses fixed-size array [26]int for 26 letters")
	fmt.Println("   - Best performance, direct array access")

	fmt.Println("\n2. HashMap Approach:")
	fmt.Println("   - Time: O(n), Space: O(1)")
	fmt.Println("   - Count character frequencies using maps")
	fmt.Println("   - More intuitive but slightly slower")

	fmt.Println("\n3. Sorting Approach:")
	fmt.Println("   - Time: O(n log n), Space: O(n)")
	fmt.Println("   - Sort both strings and compare")
	fmt.Println("   - Simple but least efficient")
}
