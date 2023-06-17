package mergeincreasingarray

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m+n)
	k := 0
	var i, j int
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			nums3[k] = nums1[i]
			i++
		} else {
			nums3[k] = nums2[j]
			j++
		}
		k++
	}
	if i < m {
		copy(nums3[k:], nums1[i:])
	}
	if j < n {
		copy(nums3[k:], nums2[j:])
	}
	copy(nums1, nums3)
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	j := 0
	for i := 0; i < n; i++ {
		j = 0
		for j < m+i {
			if nums1[j] > nums2[i] {
				copy(nums1[j+1:], nums1[j:])
				nums1[j] = nums2[i]
				break
			}
			j++
		}
		if j == m+i {
			nums1[j] = nums2[i]
		}
	}
}

func Test() {
	array := []int{1, 2, 3, 0, 0}
	merge(array, 3, []int{2, 31}, 2)
	merge1(array, 3, []int{2, 31}, 2)
	fmt.Println(array)
}
