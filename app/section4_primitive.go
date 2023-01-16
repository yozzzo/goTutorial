package main

import "fmt"

func primitive() {
	var i int = 100
	// intは32bitのpcではint32になるし64ではint64になるのでcpu依存

	// var i2 int64 = 200

	fmt.Println(i + 50)

	// fmt.Println(i + i2)

	fmt.Printf("%T\n", i)

	// fmt.Printf("%T\n", int64(i2))
}
