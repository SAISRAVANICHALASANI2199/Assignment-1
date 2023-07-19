package main

import (
	"fmt"
	"math/rand"
)

var i int
var k string
var l string
var m int
var j int
var o int

func one() int {
	fmt.Println("Enter the amount you want to load")
	fmt.Scanf("%d\n", &i)
	return i
}

func main() {
	one()
	for i != 0 {
		fmt.Printf("******\n")
		fmt.Printf("Choose from the following :\n 1.red or black \n 2.high or low (1 -- 18(low), 19 -- 36(high)) \n 3. specific number \n")
		fmt.Scanf("%d\n", &j)
		fmt.Println("Enter the amount you want to bet/enter '0' if no bet")
		fmt.Scanf("%d\n", &o)
		if o <= i {
			fmt.Println("The bet amount is :", o)
			u := rand.Intn(37)
			if j == 1 {
				fmt.Println("black and red")
				fmt.Scanf("%s\n", &k)
				//checks for even and odd
				if k == "red" || k == "black" {
					if u%2 == 0 && k == "red" {
						fmt.Println("red! You won :", o)
						i = i + o
						fmt.Println("The updated income is: ", i)
					} else if u%2 != 0 && k == "black" {
						fmt.Println("black!  won ", o)
						i = i + o
						fmt.Println("The updated income is: ", i)
					} else {
						fmt.Println("opposite color! Lost ", o)
						i = i - o
						fmt.Println("The updated income is: ", i)
					}
				}
			} else if j == 2 {
				fmt.Println("high or low (1 -- 18(low), 19 -- 36(high))")
				fmt.Scanf("%s\n", &l)
				fmt.Println(" The number is ", u)
				if l == "high" || l == "low" {
					if (0 < u) && (u < 19) && l == "low" {
						fmt.Println("low ! You won :", o)
						i = i + o
						fmt.Println("The updated income is: ", i)
					} else if (20 < u) && (u < 37) && l == "high" {
						fmt.Println("high range!  You won :", o)
						i = i + o
						fmt.Println("The updated income is: ", i)
					} else {
						fmt.Println("Not in Range!! you lost:", o)
						i = i - o
						fmt.Println("The updated income is: ", i)
					}

				}

			} else if j == 3 {
				fmt.Println(" Select a specific number(1-36)")
				fmt.Scanf("%d\n", &m)
				fmt.Println(" The number is ", u)
				if u == m {
					fmt.Println(" Won by matching numbers! Amount won is", 36*o)
					i = i + 35*o
					fmt.Println("The updated income is: ", i)
				} else {
					fmt.Println("Number is not matching! You have Lost ", o)
					i = i - o
					fmt.Println("The updated income is: ", i)
				}

			}
		} else if o > i {
			fmt.Println("Bet is more than range, select another one")
		} else if o == 0 {
			fmt.Println("You didn't participate in the bet")
		}
	}
}