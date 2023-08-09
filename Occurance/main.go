package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	input, _ := scanner.ReadString('\n')
	ans := Occurance(input)
	fmt.Println("Number of occurance of go=",ans)
}
func Occurance(input string) int {
	ans := strings.Count(input, "go")
	return ans
}
