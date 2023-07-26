package main

import "fmt"

func main () {
	var a = 1
	for ;a <= 10; a++ {
		fmt.Printf("a = %d\n", a)
	}
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("This is a loop.\n")
	}

	for i <= 20 {	//Like while in C/C++
		fmt.Printf("This is also a loop.\n")
		i = i + 1;
	}
	var sum int = 0
	for i = 1; i <= 100; i++ {
		sum += i 
	}
	fmt.Println("The sum is", sum)
	
}