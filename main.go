package main

import (
	"fmt"
)

func main() {

	a := new(int)

	*a = 20
	b := &a

	c := &b

	fmt.Println(***c)
}
