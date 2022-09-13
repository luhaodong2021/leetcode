package findkthlargest

func MergeSort(src, dst []int, s, e int) {
	if s == e {
		return
	}
	mid := (s + e) / 2
	s1, e1 := s, mid
	s2, e2 := mid+1, e
	MergeSort(src, dst, s1, e1)
	MergeSort(src, dst, s2, e2)
	k := s1
	for s1 <= e1 && s2 <= e2 {
		if src[s1] < src[s2] {
			dst[k] = src[s1]
			s1++
		} else {
			dst[k] = src[s2]
			s2++
		}
		k++
	}
	for s1 <= e1 {
		dst[k] = src[s1]
		s1++
		k++
	}
	for s2 <= e2 {
		dst[k] = src[s2]
		s2++
		k++
	}
	for k = s; k <= e; k++ {
		src[k] = dst[k]
	}
}

func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dst := make([]int, len(nums))
	MergeSort(nums, dst, 0, len(nums)-1)
	return dst[len(nums)-k]
}

func Test() {
	findKthLargest([]int{1, 3, 2, 65, 8}, 3) // 3
}
