package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	list := []int{1, 5, 6, 3, 2, 6, 9}
	fmt.Println(list)
	list = sort(list)
	fmt.Println(list)
}

func sort(arr []int) []int {
	rand.Seed(int64(time.Now().Second()))
	var lesser []int
	if (len(arr) > 0) {
		pivot := rand.Intn(len(arr))
		var greater []int
		var pivotList []int

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

		greater = sort(greater)
		lesser = sort(lesser)
		for j := 0; j < len(pivotList); j++ {
			lesser = append(lesser, pivotList[j])
		}
		for j := 0; j < len(greater); j++ {
			lesser = append(lesser, greater[j])
		}
		return lesser

	}
	return lesser
}