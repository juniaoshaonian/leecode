package d15

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	k := 0
	hashmap := make(map[int]bool)
	n := len(nums)
	ans := make([][]int, 0)
	for k < n && nums[k] <= 0 {
		if _, ok := hashmap[nums[k]]; ok {
			k++
			continue
		}
		i, j := k+1, n-1
		for i < j {
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[k], nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for j > i && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				for j > i && nums[i] == nums[i+1] {
					i++
				}
				i++
			} else {
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				j--
			}
		}

		hashmap[nums[k]] = true
		k++
	}
	return ans
}
