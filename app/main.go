package main

import (
	"fmt"
)

func outer() {
	fmt.Println("outer")
}
func main() {
	var i int = 100
	var s string = "aaaaa"
	fmt.Println(s)
	fmt.Println(i)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	i4 := 400
	fmt.Println(i4)

	outer()
	var s5 string = "Not Use"

}
