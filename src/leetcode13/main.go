package main

import "fmt"

var num = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func main() {
	fmt.Println(romanToInt("III"))
}

func romanToInt(s string) (res int) {

	for i := range s {
		if i < len(s)-1 && num[s[i]] < num[s[i+1]] {
			res -= num[s[i]]
		} else {
			res += num[s[i]]
		}

	}
	return
}
