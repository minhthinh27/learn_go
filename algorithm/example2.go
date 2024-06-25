package algorithm

import (
	"fmt"
	"math"
)

func CompareTriplets(a []int32, b []int32) []int32 {
	pointA := 0
	pointB := 0

	for i := range a {
		if a[i] > b[i] {
			pointA++
		} else if a[i] < b[i] {
			pointB++
		}
	}

	return []int32{int32(pointA), int32(pointB)}
}

func DiagonalDifference(arr [][]int32) int32 {
	var sum1, sum2 int32 = 0, 0
	point1, point2 := 0, len(arr)-1

	// fmt.Println(sum1, sum2, point1, point2)
	// //startDiagonals1, startDiagonals2 := 0, 0
	// // midDiagonals1, midDiagonals2 := 0, 0
	// // endDiagonals1, endDiagonals2 := 0, 0

	for _, v := range arr {
		sum1 += v[point1]
		sum2 += v[point2]
		point1++
		point2--
	}
	test := math.Abs(float64(sum1) - float64(sum2))

	return int32(test)
}

func Staircase(n int32) {
	var newN = int(n)
	for i := 0; i < newN; i++ {
		for j := 0; j < newN-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func PlusMinus(arr []int32) {
	// Write your code here
	positive, negative, zero := 0, 0, 0
	for _, v := range arr {
		if v > 0 {
			positive++
		} else if v < 0 {
			negative++
		} else {
			zero++
		}
	}
	fmt.Printf("%f\n", float32(positive)/float32(len(arr)))
	fmt.Printf("%f\n", float32(negative)/float32(len(arr)))
	fmt.Printf("%f\n", float32(zero)/float32(len(arr)))
}
