package main

import "fmt"

func main() {
	// Creating a hash map (called "map" in Go)
	// map[KeyType]ValueType
	ages := make(map[string]int)

	// Adding key-value pairs
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35

	fmt.Println("=== Basic Hash Map Example ===")
	fmt.Println("ages map:", ages)

	// Looking up a value (very fast - O(1))
	fmt.Println("Alice's age:", ages["Alice"])

	// Checking if a key exists
	age, exists := ages["David"]
	if exists {
		fmt.Println("David's age:", age)
	} else {
		fmt.Println("David not found in map")
	}

	fmt.Println("\n=== Why Hash Maps are Amazing ===")

	// Compare: Finding a number in a slice vs hash map
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("numbers slice:", numbers)

	// To find if 30 exists in slice, you need to check each element
	for i, num := range numbers {
		if num == 30 {
			fmt.Printf("Found 30 at index %d (checked %d elements)\n", i, i+1)
			break
		}
	}

	// With hash map, it's instant!
	numberMap := map[int]bool{10: true, 20: true, 30: true, 40: true, 50: true}
	if numberMap[30] {
		fmt.Println("Found 30 in hash map instantly!")
	}

	fmt.Println("\n=== Two Sum Example ===")
	twoSumExample()
}

func twoSumExample() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Printf("Find two numbers in %v that add to %d\n", nums, target)

	// Hash map approach
	numMap := make(map[int]int) // number -> index

	for i, num := range nums {
		complement := target - num
		fmt.Printf("At index %d: num=%d, need complement=%d\n", i, num, complement)

		if index, exists := numMap[complement]; exists {
			fmt.Printf("Found it! %d (at index %d) + %d (at index %d) = %d\n",
				complement, index, num, i, target)
			return
		}

		numMap[num] = i
		fmt.Printf("Stored: numMap[%d] = %d\n", num, i)
	}
}