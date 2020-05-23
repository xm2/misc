package containerwithmostwater

/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.


Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func minInt(i, j int) int {
	if i < j {
		return i
	}
	return j
}

/* time complexity O(n**n) */
func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			max = maxInt(max, minInt(height[i], height[j])*(j-i))
		}
	}

	return max
}

/* time complexity O(n) */
func maxAreaV2(height []int) int {
	max, left, right := 0, 0, len(height)-1
	for left < right {
		max = maxInt(max, minInt(height[left], height[right])*(right-left))
		if height[left] < len(height) {
			left++
		} else {
			right--
		}
	}

	return max
}
