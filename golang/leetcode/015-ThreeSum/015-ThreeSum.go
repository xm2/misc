package threesum

import (
	"sort"
)

/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/

func threesum(nums []int) [][3]int {
	sort.Ints(nums)
	ret := [][3]int{}

	for i := 0; i < len(nums)-2; {
		pivot := nums[i]
		for left, right := i+1, len(nums)-1; left < right; {
			sum := pivot + nums[left] + nums[right]
			leftV := nums[left]
			rightV := nums[right]

			if sum == 0 {
				ts := [3]int{pivot, nums[left], nums[right]}
				ret = append(ret, ts)

				for left < right && leftV == nums[left] {
					left++
				}

				for right > left && rightV == nums[right] {
					right--
				}

			} else {
				if sum < 0 {
					left++
				} else {
					right--
				}
			}
		}

		for i < len(nums)-2 && pivot == nums[i] {
			i++
		}
	}
	return ret
}
