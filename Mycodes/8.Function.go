package main

import "fmt"

func mysum (a int, b int) int {
	var sum int = a + b
	return sum
}

func weirdfunc (a int, b float64) float64 {
	var _res float64
	_res = float64(a) / b
	return _res
}

func swap (a string, b string) (string, string) {
	return b, a
}

func main () {
	var a int = 40086
	var b int = 15958
	fmt.Printf("The sum is %d\n", mysum(a, b))
	fmt.Printf("The divede is %f\n", weirdfunc(32, 5))
	var str1 string = "10086"
	var str2 string = "75571"
	str1, str2 = swap(str1, str2)
	fmt.Println(str1,str2)
}