package main

import "fmt"

func main(){
	//
	var a int
	a = 114514
	if a % 2 == 0 {
		fmt.Printf("%d is an even number\n", a)
	}

	var b float32 = 3.1415926535
	if (b >= 4) {
		fmt.Printf("%f is more than 4\n", b)
	} else {
		fmt.Printf("%f is less than 4\n", b)
	}

	var aa int = 100
	var bb int = 200
	if aa == 100 {
		fmt.Printf("aa is equal to 100\n")
		if bb == 200 {
			fmt.Printf("bb is equal to 200\n")
		}
	}
}

