package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	// Map to group strings by their character frequency signature
	// Using [26]byte instead of [26]int for better memory efficiency
	group := make(map[[26]byte][]string)

	for _, str := range strs {
		// Create frequency count array for current string
		freq := [26]byte{}
		for _, char := range str {
			freq[char-'a']++ // Count frequency of each character
		}

		// Group strings with same frequency signature
		group[freq] = append(group[freq], str)
	}

	// Convert map values to result slice
	result := [][]string{}
	for _, anagramGroup := range group {
		result = append(result, anagramGroup)
	}

	return result
}

// Alternative version with overflow protection (if needed)
func groupAnagramsSafe(strs []string) [][]string {
	group := make(map[[26]byte][]string)

	for _, str := range strs {
		freq := [26]byte{}
		for _, char := range str {
			if freq[char-'a'] < 255 { // Prevent overflow
				freq[char-'a']++
			}
		}
		group[freq] = append(group[freq], str)
	}

	result := [][]string{}
	for _, anagramGroup := range group {
		result = append(result, anagramGroup)
	}

	return result
}

func main() {
	// Test cases
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("Input: %v\n", strs1)
	fmt.Printf("Output: %v\n\n", groupAnagrams(strs1))

	strs2 := []string{""}
	fmt.Printf("Input: %v\n", strs2)
	fmt.Printf("Output: %v\n\n", groupAnagrams(strs2))

	strs3 := []string{"a"}
	fmt.Printf("Input: %v\n", strs3)
	fmt.Printf("Output: %v\n", groupAnagrams(strs3))

	// Test with longer string to show byte efficiency
	strs4 := []string{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"}
	fmt.Printf("Input: %v\n", strs4)
	fmt.Printf("Output: %v\n", groupAnagrams(strs4))
}
