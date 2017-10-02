package main

import (
	"fmt"
	"sort"
)

func DoAppendSlices() {

	slice1 := []int{1, 4, 6}
	slice2 := []int{2, 3, 5}

	//var finishedSlice[6] int

	slice1 = append(slice1, slice2...)

	//myIntSlice := []int{1, 4, 6, 9, 22, 55, 28, -10000000000}
	//myIntSlice = append(myIntSlice, []int{2, 3, 5, 11234}...)

	sort.Ints(slice1)
	fmt.Println(slice1)
}
func main() {
	DoAppendSlices()
}