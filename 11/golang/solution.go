package problem11

// 11. Container With Most Water
// Difficulty: Medium
// Pattern:    opposite-end
// Link:       https://leetcode.com/problems/container-with-most-water/
//
// TODO: implement solution

func Min[T ~int | ~float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T ~int | ~float64](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func maxArea(height []int) int {
	lo, hi := 0, len(height)-1
	result := -1000
	area := 0
	for lo < hi {
		area = Min(height[lo], height[hi]) * (hi - lo)
		result = Max(result, area)
		if height[lo] < height[hi] {
			lo++
		} else {
			hi--
		}
	}
	return result
}
