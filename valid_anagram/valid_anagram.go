package main

import (
	"fmt"
	"sort"
	"strings"
)

// Approach 1: Using character frequency count (most efficient)
// Time: O(n), Space: O(1) - since we only have 26 lowercase letters
func isAnagram(s string, t string) bool {
	// If lengths are different, they can't be anagrams
	if len(s) != len(t) {
		return false
	}

	// Count frequency of each character
	charCount := make(map[rune]int)

	// Count characters in string s
	for _, char := range s {
		charCount[char]++
	}

	// Subtract characters from string t
	for _, char := range t {
		charCount[char]--
		// If count goes negative, t has more of this character than s
		if charCount[char] < 0 {
			return false
		}
	}

	// Check if all counts are zero
	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}

	return true
}

// Approach 2: Using sorting (less efficient but simpler)
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

// Approach 3: Using array instead of map for better performance (since only 26 lowercase letters)
// Time: O(n), Space: O(1)
func isAnagramArray(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Array to count frequency of each letter (a-z)
	count := [26]int{}

	// Count characters in both strings
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++ // Increment for s
		count[t[i]-'a']-- // Decrement for t
	}

	// Check if all counts are zero
	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
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
		result2 := isAnagramSort(tc.s, tc.t)
		result3 := isAnagramArray(tc.s, tc.t)

		fmt.Printf("Expected: %t\n", tc.expected)
		fmt.Printf("HashMap approach: %t\n", result1)
		fmt.Printf("Sorting approach: %t\n", result2)
		fmt.Printf("Array approach: %t\n", result3)

		if result1 == tc.expected && result2 == tc.expected && result3 == tc.expected {
			fmt.Println("✅ PASSED")
		} else {
			fmt.Println("❌ FAILED")
		}
		fmt.Println(strings.Repeat("-", 30))
	}

	// Performance explanation
	fmt.Println("\n=== Algorithm Explanation ===")
	fmt.Println("1. HashMap Approach (Recommended):")
	fmt.Println("   - Time: O(n), Space: O(1)")
	fmt.Println("   - Count character frequencies using a map")
	fmt.Println("   - Most intuitive and efficient")

	fmt.Println("\n2. Sorting Approach:")
	fmt.Println("   - Time: O(n log n), Space: O(n)")
	fmt.Println("   - Sort both strings and compare")
	fmt.Println("   - Simple but less efficient")

	fmt.Println("\n3. Array Approach (Most Efficient):")
	fmt.Println("   - Time: O(n), Space: O(1)")
	fmt.Println("   - Uses fixed-size array for 26 letters")
	fmt.Println("   - Best performance for this specific problem")
}
