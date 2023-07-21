package main

import (
	"fmt"
)

func main() {
	//float32 and float64, float64 is recommanded.
	var f1 float32
	f1 = 2147483647
	fmt.Println(f1)
	f1 = 3.14159265358
	//precision loss
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	fmt.Println("num1=", num1, "num2=", num2)
	fmt.Printf("You can use %%T to check var's type, %T\n", num2)
	fmt.Printf("printf:%f", num2)
}	