package main

import "fmt"
import "unsafe"

func main() {
	var flag = false
	fmt.Println("Flag is",flag)
	//use unsafe.Sizeof() to check the sapce a variable uses.
	fmt.Println("The space that flag uses =", unsafe.Sizeof(flag))

	//String in Go is not changable
	var address string = "China Mainland"
	fmt.Println(address)
	fmt.Println("-------------------------------------------------------------")
	//use ` to let the string print directly, just like this:
	str1 := 
`package main
import "fmt"
import "unsafe"
var flag = false
fmt.Println("Flag is",flag)`
	fmt.Println(str1)
	str2 := "Hello,"
	str3 := " World!"
	fmt.Println(str2+str3)

	//Change Types
	var i1 int = 100
	var n1 = int64(i1)
	fmt.Printf("%T, %T\n",i1, n1)
	//use fmt.Sprintf() to change number to string
	var number int = 10086
	var string1 string = fmt.Sprintf("%d", number)
	fmt.Println(string1)
}