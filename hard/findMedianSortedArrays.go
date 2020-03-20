package hard

import "math"

/*
	给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

	请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

	你可以假设 nums1 和 nums2 不会同时为空。
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	left, right := (l1+l2+1)/2, (l1+l2+2)/2

	minFunc := func(a, b int) int {
		return int(math.Min(float64(a), float64(b)))
	}

	var getKthFunc func([]int, int, int, []int, int, int, int) float64
	getKthFunc = func(nums1 []int, start1, end1 int, nums2 []int, start2, end2, k int) float64 {
		len1 := end1 - start1 + 1
		len2 := end2 - start2 + 1
		//让 len1 的长度小于 len2，这样就能保证如果有数组空了，一定是 len1
		if len1 > len2 {
			return getKthFunc(nums2, start2, end2, nums1, start1, end1, k)
		}
		if len1 == 0 {
			return float64(nums2[start2+k-1])
		}

		if k == 1 {
			return float64(minFunc(nums1[start1], nums2[start2]))
		}

		i := start1 + minFunc(len1, k/2) - 1
		j := start2 + minFunc(len2, k/2) - 1

		if nums1[i] > nums2[j] {
			return getKthFunc(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
		} else {
			return getKthFunc(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
		}
	}

	return (getKthFunc(nums1, 0, l1-1, nums2, 0, l2-1, left) + getKthFunc(nums1, 0, l1-1, nums2, 0, l2-1, right)) * 0.5
}
