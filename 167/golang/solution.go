package problem167

// 167. Two Sum II
// Difficulty: Easy
// Pattern:    opposite-end
// Link:       https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
//
// TODO: implement solution

func twoSum(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)-1
	sum := 0
	for lo < hi {
		sum = numbers[lo] + numbers[hi]
		switch {
		case sum == target:
			return []int{lo + 1, hi + 1}
		case sum < target:
			lo++
		case sum > target:
			hi--
		}
	}
	return []int{lo + 1, hi + 1}
}
