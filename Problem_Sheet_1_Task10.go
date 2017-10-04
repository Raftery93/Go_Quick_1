package main

import "fmt"

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}


func main(){

var input string

	fmt.Println("Please Enter a String:")
	fmt.Scanf("%s", &input)

	fmt.Println(Reverse(input))

}