package main

import (
	"fmt"
	"math"
)


	//function to run newthons method
	func z_next(x float64, z float64) float64 {
		return z - (((z*z) - x) / (2*z))
	}//z_next

	func main(){
		//number to calculate square root
		x := 20.0

		//initial guess
		z :=  float64(1)

		//loop through using newtons method until next guess == current guess
		for z = 1.0; z != z_next(x,z); z = z_next(x,z) {
			//display current guess
			fmt.Printf("The current Guess is %12.8f\n ", z)
		}//for

		fmt.Printf("The square of %.2f is %.2f\n", x, z)
		

		fmt.Printf("math.Sqrt gives the value as %.2f\n", math.Sqrt(x))
	}//main