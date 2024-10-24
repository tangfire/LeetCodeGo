package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(789987))
}

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}

	temp := 0
	for {
		if x > temp {
			temp = temp*10 + x%10
			x /= 10
		} else {
			break
		}
	}

	return x == temp || x == temp/10

}
