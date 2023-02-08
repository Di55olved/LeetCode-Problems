func merge(nums1 []int, nums2 []int) []int {
	arr := []int{}
	if len(nums1) == len(nums2) {
		check := 0
		for i := 0; i < len(nums1); i++ {
			for j := check; j < len(nums2); j++ {
				if check == len(nums2) {
					return arr
				}
				if nums1[i] == nums2[j] {
					arr = append(arr, nums1[i])
					arr = append(arr, nums2[j])
					check++
					break
				} else if nums1[i] > nums2[j] {
					arr = append(arr, nums2[j])
					check++
				} else {
					arr = append(arr, nums1[i])
					break
				}
			}
		}
		for i := check; i < len(nums2); i++ {
			arr = append(arr, nums2[i])
		}
		return arr
	} else if len(nums1) > len(nums2) {
		check := 0
		for i := 0; i < len(nums1); i++ {
			for j := check; j < len(nums2); j++ {
				if check == len(nums2) {
					return arr
				}
				if nums1[i] == nums2[j] {
					arr = append(arr, nums1[i])
					arr = append(arr, nums2[j])
					check++
					break
				} else if nums1[i] > nums2[j] {
					arr = append(arr, nums2[j])
					check++
				} else {
					arr = append(arr, nums1[i])
					break
				}
			}
		}
		for i := check; i < len(nums2); i++ {
			arr = append(arr, nums2[i])
		}
		return arr
	} else {
		check := 0
		for i := 0; i < len(nums2); i++ {
			for j := check; j < len(nums1); j++ {
				if check == len(nums1) {
					return arr
				}
				if nums2[i] == nums1[j] {
					arr = append(arr, nums2[i])
					arr = append(arr, nums1[j])
					check++
					break
				} else if nums2[i] > nums1[j] {
					arr = append(arr, nums1[j])
					check++
				} else {
					arr = append(arr, nums2[i])
					break
				}
			}
		}
		for i := check; i < len(nums1); i++ {
			arr = append(arr, nums1[i])
		}
		return arr
	}
}
