package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	n := len(nums1)
	sort.Ints(nums1)

	m := nums1[(n/2)]
	if n % 2 == 0 {
		l := nums1[(n/2)-1]
		return float64(l+m)/2
	} else {
		return float64(m)
	}

	return 0
}
