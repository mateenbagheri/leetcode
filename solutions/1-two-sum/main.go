package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumsOptimal(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for index, value := range nums {
		if requiredIndex, isPresent := indexMap[target-value]; isPresent {
			return []int{requiredIndex, index}
		}
		indexMap[value] = index
	}
	return []int{}
}
