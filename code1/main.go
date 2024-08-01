package main

import (
	"fmt"
)

// 189. Rotate Array
// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

// Example 1:

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
// Example 2:

// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]

// Constraints:

// 1 <= nums.length <= 105
// -231 <= nums[i] <= 231 - 1
// 0 <= k <= 105

type data struct {
	k     int
	input []int
}

func main() {
	k := 2
	input := []int{-1}
	data1 := data{
		k:     k,
		input: input,
	}
	data1.Rotate()
	fmt.Println(input)

	k = 3
	input = []int{-1, -100, 3, 99}
	data1 = data{
		k:     k,
		input: input,
	}
	data1.Rotate()
	fmt.Println(input)
}

func (d *data) Rotate() {
	nums := d.input
	k := d.k

	input_size := len(nums)

	if k > input_size {
		k = k % input_size
	}
	fmt.Println(k % input_size)
	if k == 0 {
		return
	}
	numsk := nums[k:input_size]
	nums0 := nums[0 : input_size-k+1]
	fmt.Println("initnums", nums)
	fmt.Println("numsk", numsk)
	fmt.Println("nums0", nums0)
	nums = append(numsk, nums0...)
	fmt.Println("nums", nums)

	d.input = nums

}
