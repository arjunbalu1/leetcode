package main

import "fmt"

// Valid Palindrome - Pure Go Solution (No External Libraries)
// Time Complexity: O(n) - single pass through the string
// Space Complexity: O(1) - only using two pointers, no extra space
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}
		if toLowerCase(s[left]) != toLowerCase(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLowerCase(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

// Brute force approach: clean then check
func isPalindromeBruteForce(s string) bool {
	cleaned := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		char := s[i]
		if isAlphanumeric(char) {
			cleaned = append(cleaned, toLowerCase(char))
		}
	}
	n := len(cleaned)
	for i := 0; i < n/2; i++ {
		if cleaned[i] != cleaned[n-1-i] {
			return false
		}
	}
	return true
}

// For visualization in main
func cleanAndShow(s string) string {
	fmt.Printf("Original: %q\n", s)
	fmt.Print("Processing: ")
	cleaned := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		char := s[i]
		if isAlphanumeric(char) {
			lowerChar := toLowerCase(char)
			cleaned = append(cleaned, lowerChar)
			fmt.Printf("%c", lowerChar)
		} else {
			fmt.Printf("[%c-skip]", char)
		}
	}
	result := string(cleaned)
	fmt.Printf("\nCleaned: %q\n", result)
	return result
}

func main() {
	fmt.Println("=== Valid Palindrome Solution (No External Libraries) ===")
	fmt.Println("A palindrome reads the same forward and backward after:")
	fmt.Println("1. Converting to lowercase")
	fmt.Println("2. Removing non-alphanumeric characters")
	fmt.Println()

	testCases := []struct {
		input    string
		expected bool
		name     string
	}{
		{"A man, a plan, a canal: Panama", true, "Classic palindrome"},
		{"race a car", false, "Not a palindrome"},
		{" ", true, "Empty after cleaning"},
		{"", true, "Empty string"},
		{"Madam", true, "Simple palindrome"},
		{"No 'x' in Nixon", true, "Complex palindrome"},
		{"Mr. Owl ate my metal worm", true, "Long palindrome"},
		{"12321", true, "Numeric palindrome"},
		{"A Santa at NASA", true, "Mixed case palindrome"},
		{"Was it a car or a cat I saw?", true, "Question palindrome"},
		{"Nope", false, "Simple non-palindrome"},
		{"12345", false, "Numeric non-palindrome"},
		{"a", true, "Single character"},
		{"Aa", true, "Two same characters different case"},
		{"Ab", false, "Two different characters"},
	}

	fmt.Println("=== Testing Both Approaches ===")
	for i, tc := range testCases {
		fmt.Printf("\n--- Test Case %d: %s ---\n", i+1, tc.name)
		cleanAndShow(tc.input)
		result1 := isPalindrome(tc.input)
		fmt.Printf("Optimized result: %v\n", result1)
		result2 := isPalindromeBruteForce(tc.input)
		fmt.Printf("Brute force result: %v\n", result2)
		if result1 == result2 && result1 == tc.expected {
			fmt.Printf("✅ Both approaches agree: %v (Expected: %v)\n", result1, tc.expected)
		} else {
			fmt.Printf("❌ Results don't match! Opt:%v Brute:%v Expected:%v\n", result1, result2, tc.expected)
		}
	}

	fmt.Println("\n=== Manual Helper Function Tests ===")
	fmt.Println("Testing our custom helper functions:")
	testChars := []byte{'a', 'Z', '5', ' ', ',', '!', '0', '9'}
	fmt.Print("isAlphanumeric tests: ")
	for _, char := range testChars {
		fmt.Printf("%c:%v ", char, isAlphanumeric(char))
	}
	fmt.Println()
	testUppers := []byte{'A', 'B', 'Z', 'a', '5', ' '}
	fmt.Print("toLowerCase tests: ")
	for _, char := range testUppers {
		fmt.Printf("%c→%c ", char, toLowerCase(char))
	}
	fmt.Println()

	fmt.Println("\n=== Algorithm Explanation ===")
	fmt.Println("1. OPTIMIZED TWO-POINTER (Recommended):")
	fmt.Println("   • Time: O(n) - single pass")
	fmt.Println("   • Space: O(1) - only two pointer variables")
	fmt.Println("   • No string preprocessing required")
	fmt.Println("   • Handles case conversion on-the-fly")
	fmt.Println()
	fmt.Println("2. BRUTE FORCE APPROACH:")
	fmt.Println("   • Time: O(n) - two passes (clean + check)")
	fmt.Println("   • Space: O(n) - stores cleaned byte slice")
	fmt.Println("   • More readable step-by-step process")
	fmt.Println("   • Good for understanding the problem")
	fmt.Println()
	fmt.Println("=== Key Implementation Details ===")
	fmt.Println("• No external libraries used (only fmt for output)")
	fmt.Println("• Manual ASCII case conversion: 'A' + 32 = 'a'")
	fmt.Println("• Character range checks: 'a' <= c <= 'z'")
	fmt.Println("• Byte slice operations instead of string building")
	fmt.Println("• Two-pointer technique for O(1) space complexity")
	fmt.Println()
	fmt.Println("ASCII Values Used:")
	fmt.Println("• 'A' = 65, 'Z' = 90")
	fmt.Println("• 'a' = 97, 'z' = 122")
	fmt.Println("• '0' = 48, '9' = 57")
	fmt.Println("• Uppercase to lowercase: add 32")
}
