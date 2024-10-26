package main

import "fmt"

func main() {
	str := letterCombinations("23")
	for i := range str {
		fmt.Println(str[i])
	}
}

var mapping = [...]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) (ans []string) {
	var n = len(digits)
	if n == 0 {
		return
	}
	var dfs func(int)

	path := make([]byte, n)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path))
			return
		}

		for _, s := range mapping[digits[i]-'0'] {
			path[i] = byte(s)
			dfs(i + 1)
		}
	}

	dfs(0)
	return

}
