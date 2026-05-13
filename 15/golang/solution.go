package problem15

import "slices"

// 15. 3Sum
// Difficulty: Medium
// Pattern:    fix+scan
// Link:       https://leetcode.com/problems/3sum/
//
// TODO: implement solution

// num[i] + num[l] + num[r] == 0 TARGET

// num[i] + num[l] + num[r] > 0 => num[l] + num[r] > sum[] -num[i]

// num[i] + num[l] + num[r] < 0 => num[l] + num[r] <  - num[i]
func threeSum(nums []int) [][]int {
	results := [][]int{}
	slices.Sort(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			switch {
			case sum == 0:
				results = append(results, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			case sum > 0:
				r--
			case sum < 0:
				l++
			}
		}
	}
	return results
}
