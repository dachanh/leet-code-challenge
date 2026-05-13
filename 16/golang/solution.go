package problem16

import "slices"

// 16. 3Sum Closest
// Difficulty: Medium
// Pattern:    fix+scan (two pointers)
// Link:       https://leetcode.com/problems/3sum-closest/
//
// TODO: implement solution

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if abs(target-sum) < abs(target-closest) {
				closest = sum
			}
			switch {
			case sum < target:
				l++
			case sum > target:
				r--
			case sum == target:
				return target
			}
		}
	}
	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
