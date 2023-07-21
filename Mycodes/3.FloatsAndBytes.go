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
	fmt.Printf("printf:%f\n", num2)
	fmt.Println("--------------------------------------------------------------")
	//About Byte
	var char1 byte = 'a'
	//var char2 byte = '0'
	fmt.Printf("%T\n",char1)
	fmt.Println(char1)//This will print alphabet a's ASCII 97
	fmt.Printf("Char1=%c\n", char1) //This will print a
	//The character in Go use UTF-8 : http://www.mytju.com/classcode/tools/encode_utf8.asp
	var c4 int = 22269 //22269-> '国'
	fmt.Printf("c4=%c\n", c4)
}	