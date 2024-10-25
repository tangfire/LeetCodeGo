package main

import "fmt"

func main() {
	strs := []string{"flight", "fight", "friend"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	sh := strs[0]
	for j := range sh {
		for _, s := range strs {
			if j == len(s) || sh[j] != s[j] {
				return sh[:j]
			}
		}
	}
	return sh

}
