package main

import (
	"fmt"
	"sort"
)

func DoAppendSlices() {

	slice1 := []int{1, 4, 6}
	slice2 := []int{2, 3, 5}

	slice1 = append(slice1, slice2...)

	sort.Ints(slice1)
	fmt.Println(slice1)
}
func main() {
	DoAppendSlices()
}