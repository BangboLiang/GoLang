package main

import "fmt"

//global variables

var n1 = 100
var n2 = 200
var name1 = "BBL"

//Also a define of variables
var (
	n3 = 900
	n4 = 1600
	name2 = "LBB"
)

func main() {
	//variable's define, initial value is 0
	var i int
	var init int
	//type recgonize by the value.
	var a = 3.1415926
	//omit var := use only the first time.
	name := "Bangbo"
	i = 10
	fmt.Println("i's value is", i)
	fmt.Println("The variable's initial value is", init)
	fmt.Println("a=",a)
	fmt.Println("My name is", name)
	//name:= "111" .\var.go:18:6: no new variables on left side of :=

	// Define many variables at one time
	x1, x2, x3 := 12, 3.14, "Liang"
	fmt.Println(x1, x2, x3)

	fmt.Println("The global variables:", n1, n2, n3, n4, name1, name2)
	fmt.Println(name1+name2)

	var big uint64
	big = 0
	fmt.Println(big-1)

	var char byte //byte 0-255
	char = 64
	fmt.Println(char) 
}