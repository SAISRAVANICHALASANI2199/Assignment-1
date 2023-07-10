package main

import "fmt"


func reverseNumber(num int) int {
	revNum := 0
	for num > 0 {
		revNum = revNum*10 + num%10
		num = num / 10
	}
	return revNum
}

func main() {
	num := 2199
	fmt.Printf("Reverse of Number is %d", reverseNumber(num))
}
