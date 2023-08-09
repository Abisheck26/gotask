package main

import (
	
	"fmt"
	
)

func main() {
	var input string
	fmt.Scanln(&input)
	ans := Reverse(input)
	fmt.Println(ans)
}
func Reverse(input string) string {
	val := 0
	var str = []byte(input)
	for i := 0; i < len(str)-1; i++ {
		if val == 2 {
			temp := str[i]
			str[i] = str[i+1]
			str[i+1] = temp
			i++
			val = 0
		} else {
			val++
		}
	}
	ans := string(str)
	return ans
}
