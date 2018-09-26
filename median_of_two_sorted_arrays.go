package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 1 {
		return findKth(nums1, nums2, n/2+1)
	} else {
		smaller := findKth(nums1, nums2, n/2)
		bigger := findKth(nums1, nums2, n/2+1)
		fmt.Printf("smaller:%v, bigger:%v\n", smaller, bigger)
		return (float64(smaller) + float64(bigger)) / 2
	}

}

func findKth(nums1 []int, nums2 []int, k int) float64 {
	if len(nums1) == 0 {
		return float64(nums2[k-1])
	} else if len(nums2) == 0 {
		return float64(nums1[k-1])
	}

	if k == 1 {
		return math.Min(float64(nums1[0]), float64(nums2[0]))
	}

	m := k / 2

	if len(nums2) < m || (len(nums1) >= m && nums1[m-1] < nums2[m-1]) {
		return findKth(nums1[k/2:], nums2, k-k/2)
	}

	return findKth(nums1, nums2[k/2:], k-k/2)
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
