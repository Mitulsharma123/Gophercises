//Write a function with one variadic parameter that finds the greates number in a list of numbers.

package main

import "fmt"

func fun(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	nums := []int{5, 1, 4, 3, 7}
	fmt.Println(fun(nums...))
}
