package main

import "fmt"

func main() {
	var rows int
	fmt.Scanln(&rows)
	ans := Pascal(rows)
	fmt.Println(ans)

}
func Pascal(rows int) [][]int {
	ans := make([][]int, 0)

	for i := 0; i < rows; i++ {
		arr := make([]int, 0)
		for j := 0; j <= i; j++ {
			value := 1
			if i > 1 && j > 0 && j < i {
				value = ans[i-1][j-1] + ans[i-1][j]
			}
			arr = append(arr, value)
		}
		ans = append(ans, arr)
	}

	return ans
}