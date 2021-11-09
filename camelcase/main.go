//https://www.hackerrank.com/challenges/camelcase/problem

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Println("Input is", input)
	answer := 1 //string is one character

	for _, ch := range input {
		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
		}
	}
	//fmt.Println(ch)
	fmt.Println(answer)

}
