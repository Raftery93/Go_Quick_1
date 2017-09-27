// Author: Conor Raftery
// Date: 2017-09-27

package main

import "fmt"

func main(){

	var lowest = 0
	var highest = 0

	nums := [9]int{15, 8, 1, 34, 55, 99, 7, 94, 5}

	lowest = nums[0]
	highest = nums[0]

	for i := range nums {
		if nums[i] < lowest {
			lowest = nums[i]
		}
		if nums[i] > highest {
			highest = nums[i]
		}
	}


fmt.Println(lowest, highest)
}