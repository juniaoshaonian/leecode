package sort

import "math/rand"

func sortArray(nums []int) []int {
	var quicksort func(int, int)
	quicksort = func(l, r int) {
		if l >= r {
			return
		}
		randam := FindRandom(l, r, nums)
		quicksort(l, randam)
		quicksort(randam+1, r)
	}
	return nums
}

func FindRandom(l, r int, nums []int) int {
	randIndex := rand.Intn(r-l) + l
	randVal := nums[randIndex]
	i, j := l, r
	for i <= j {
		for nums[i] <= randVal {
			i++
		}

		for nums[j] > randVal {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return i

}
