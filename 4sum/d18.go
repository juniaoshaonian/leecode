package _sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	ans := make([][]int, 0, 16)
	a := 0
	n := len(nums)
	for a < n {
		b := a + 1
		for b < n {
			c, d := b+1, n-1
			for c < d {
				if nums[a]+nums[b]+nums[c]+nums[d] == target {
					ans = append(ans, []int{nums[a], nums[b], nums[c], nums[d]})
					for c < d && nums[c] == nums[c+1] {
						c++
					}
					for c < d && nums[d] == nums[d-1] {
						d--
					}
					c++
					d--
				} else if nums[a]+nums[b]+nums[c]+nums[d] < target {
					for c < d && nums[c] == nums[c+1] {
						c++
					}
					c++
				} else {
					for c < d && nums[d] == nums[d-1] {
						d--
					}
					d--
				}

			}
			for b+1 < n && nums[b] == nums[b+1] {
				b++
			}
			b++
		}
		for a+1 < n && nums[a] == nums[a+1] {
			a++
		}
		a++
	}
	return ans

}
