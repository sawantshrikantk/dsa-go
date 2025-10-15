package pairsum

import "sort"

func PairSum(arr []int, target int) []int {
	// sort the int array
	sort.Ints(arr)
	// if first element is greater than target, return nil
	// if target is less than the first number return nil
	if arr[0] > target || target < arr[0] {
		return nil
	}
	index := 0
	return GetPairSum(arr, target, index)
}

func GetPairSum(arr []int, target int, index int) []int {
	if index >= len(arr) {
		return nil
	}
	for i, j := index, index+1; j < len(arr); {
		if arr[i]+arr[j] == target {
			return []int{arr[i], arr[j]}
		}
		j++
	}
	return GetPairSum(arr, target, index+1)
}
