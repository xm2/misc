package medianoftwosortedarrays

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/

// refer to leetcode solution
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if n < m {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	} //  m =< n to simplify the logic

	if m == 0 {
		if n%2 == 0 {
			return float64(nums2[n/2-1]+nums2[n/2]) / 2.0
		}
		return float64(nums2[n/2])
	}

	imin, imax := 0, m
	maxleft, minright := 0, 0
	for {
		if imin > imax {
			break
		}

		i := (imin + imax) / 2
		j := (m+n+1)/2 - i
		if i > 0 && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else if i < m && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else {
			if i == 0 {
				maxleft = nums2[j-1]
			} else if j == 0 {
				maxleft = nums1[i-1]
			} else {
				if nums1[i-1] > nums2[j-1] {
					maxleft = nums1[i-1]
				} else {
					maxleft = nums2[j-1]
				}
			}

			if (m+n)%2 == 1 {
				return float64(maxleft)
			}

			if i == m {
				minright = nums2[j]
			} else if j == n {
				minright = nums1[i]
			} else {
				if nums1[i] > nums2[j] {
					minright = nums2[j]
				} else {
					minright = nums1[i]
				}
			}

			return float64(minright+maxleft) / (2.0)

		}

	}

	return 0.0
}
