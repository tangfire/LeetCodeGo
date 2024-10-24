package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	n := twoSum(nums, target)
	for v := range n {
		fmt.Println(v)
	}
}

func twoSum(nums []int, target int) []int {
	mapNum := map[int]int{}
	for i, v := range nums {
		tempNum := target - v
		numIndex, ok := mapNum[tempNum]
		if ok {
			return []int{i, numIndex}
		} else {
			mapNum[v] = i
		}
	}
	return []int{}
}
