package main

import (
	"fmt"
	"math"
)

func getThreeBiggestNum(nums []int) (int, int, int) {

	first, second, third := math.MinInt32, math.MinInt32, math.MinInt32

	for _, num := range nums {
		switch {
		case num > first:
			third = second
			second = first
			first = num
		case num > second:
			third = second
			second = num
		case num > third:
			third = num
		}
	}

	return first, second, third
}

func getTwoSmallestNum(nums []int) (int, int) {
	first, second := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		switch {
		case num < first:
			second = first
			first = num
		case num < second:
			second = num
		}
	}

	return first, second
}

func maximumProduct(nums []int) int {

	firstBiggest, secondBiggest, thirdBiggest := getThreeBiggestNum(nums)
	fmt.Println(firstBiggest, secondBiggest, thirdBiggest)
	firstSmallest, secondSmallest := getTwoSmallestNum(nums)
	fmt.Println(firstSmallest, secondSmallest)
	a := firstBiggest * secondBiggest * thirdBiggest
	b := firstBiggest * firstSmallest * secondSmallest

	if firstSmallest == math.MinInt32 {
		return a
	}
	// fmt.Println(a, b)
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maximumProduct([]int{-100, -98, -1, 2, 3, 4}))
}
