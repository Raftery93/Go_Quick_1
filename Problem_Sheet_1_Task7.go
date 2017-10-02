package main

import "fmt"

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {

	var input string

	fmt.Println("Please Enter a String:")
	fmt.Scanf("%s", &input)

	if isPalindrome(input) == true{
		fmt.Printf("%s is a palindrome", input)
	}else{
		fmt.Printf("%s is not a palindrome", input)
	}

}