package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	list := []int{7, 0, 6, 4, 2, 45, 21, 5, 6, 3, 2, 6, 9}
	fmt.Println(list)
	list = sort(list)
	fmt.Println(list)
}

func sort(arr []int) []int {
	rand.Seed(int64(time.Now().Second()))
	if (len(arr) > 0) {
		lesser, pivotList, greater := distributeElements(arr)
		greater = sort(greater)
		lesser = sort(lesser)
		return mergeSlices(lesser, pivotList, greater)

	}
	return arr
}

func mergeSlices(destination []int, source1 []int, source2 []int) []int {
	destination = append(destination, source1...)
	destination = append(destination, source2...)
	return destination
}

func distributeElements(arr []int) ([]int, []int, []int,) {
	pivot := rand.Intn(len(arr))
	var (
		lesser []int
		greater []int
		pivotList []int
	)

	for i := 0; i < len(arr); i++ {
		if (arr[i] == arr[pivot] ) {
			pivotList = append(pivotList, arr[i])
		} else {
			if (arr[i] > arr[pivot] ) {
				greater = append(greater, arr[i])
			} else {
				lesser = append(lesser, arr[i])
			}
		}
	}
	return lesser, pivotList, greater
}