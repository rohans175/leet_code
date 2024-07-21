package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFinalArray(t *testing.T) {
	type testData struct {
		testName string
		data
		expected []int
	}
	args := []testData{
		{
			testName: "basic 1",
			data: data{
				nums1: []int{1, 2, 3, 0, 0, 0},
				nums2: []int{4, 5, 6},
				m:     3,
				n:     3,
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			testName: "basic 2",
			data: data{
				nums1: []int{1},
				nums2: []int{},
				m:     1,
				n:     0,
			},
			expected: []int{1},
		},
		{
			testName: "basic 3",
			data: data{
				nums1: []int{0},
				nums2: []int{1},
				m:     0,
				n:     1,
			},
			expected: []int{1},
		},
		{
			testName: "twisted 1",
			data: data{
				nums1: []int{4, 0, 0, 0, 0, 0},
				nums2: []int{1, 2, 3, 5, 6},
				m:     1,
				n:     5,
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, eachArg := range args {
		fmt.Println("Test: ", eachArg.testName)
		eachArg.FinalArray()
		if !reflect.DeepEqual(eachArg.nums1, eachArg.expected) {
			fmt.Println("num1 ", eachArg.nums1)
			fmt.Println("expected: ", eachArg.expected)
			fmt.Println("num2 ", eachArg.nums2)
			fmt.Println("m ", eachArg.m)
			fmt.Println("n ", eachArg.n)
			t.Errorf("FAILED")
		} else {
			fmt.Println("PASSED")
		}
	}

}
