package main

import "fmt"
import "unsafe"
import "strconv"

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

	//function strconv.ParseXXX()
	var strb string = "true"
	var b bool
	//Function: strconv.ParseBool(str) -> (value bool, err error)
	//Use _ to ignore the second return value.
	b, _ = strconv.ParseBool(strb)
	fmt.Printf("b's type is %T, b=%v\n", b, b)

	//Function float
	var strf string = "3.1415926535"
	var f float64
	f, _ = strconv.ParseFloat(strf,64)
	fmt.Printf("f's type is %T, f=%v\n", f, f)

	//Function int
	var stri string = "114514"
	var inti int64
	inti, _ = strconv.ParseInt(stri, 10, 64)
	fmt.Printf("inti's type is %T, inti=%v\n", inti, inti)
	var int2 = int(inti)
	fmt.Printf("int2's type is %T, int2=%v\n", int2, int2)
}