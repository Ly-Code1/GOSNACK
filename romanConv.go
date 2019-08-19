package main

import "fmt"

func romannum(inpnum int) string {
	outroman := ""
	digit := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(digit); i++ {
		for inpnum >= digit[i] {
			outroman = outroman + roman[i]
			inpnum = inpnum - digit[i]
		}
	}
	return outroman
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println("Number ", i, ": ", romannum(i))
	}

}
