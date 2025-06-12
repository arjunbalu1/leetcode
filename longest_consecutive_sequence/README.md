# Longest Consecutive Sequence

## Problem Description
Given an unsorted array of integers `nums`, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in **O(n) time**.

## Solutions Provided

### 1. HashSet Approach (Optimal) - O(n)
- Uses a HashSet for O(1) lookups
- Only starts counting from sequence beginnings
- Each number is visited at most twice

### 2. Sorting Approach - O(n log n)
- Sort the array first
- Count consecutive sequences while handling duplicates
- Easier to understand but not optimal

### 3. Union-Find Approach - O(n)
- Advanced data structure approach
- Union consecutive numbers into components
- Find the largest component size

## Key Insight
The HashSet approach achieves O(n) time complexity by only starting to count from the beginning of each sequence (when `num-1` doesn't exist in the set).

## Usage
```bash
go run longest_consecutive_sequence.go
``` 