package main

import "fmt"

func main (){
	var a int = 4
	var b int32
	var c float32
	var ptr *int
	ptr = &a
	fmt.Printf("The type of a is: %T\n",a)
	fmt.Printf("The type of b is: %T\n",b)
	fmt.Printf("The type of c is: %T\n",c)
	fmt.Printf("The type of ptr is: %T\n", ptr)
	fmt.Printf("The type of *ptr is: %T\n", *ptr)

	b = int32(a)
	fmt.Printf("The value of a is %d\n",a)
	fmt.Printf("The value of b is %d\n",b)
	fmt.Printf("The value of ptr is %d\n",*ptr)
	a = 114514
	fmt.Printf("Now the value of a is %d\n",a)
	fmt.Printf("Now the value of b is %d\n",b)
	fmt.Printf("Now the value of ptr is %d\n",*ptr)
}