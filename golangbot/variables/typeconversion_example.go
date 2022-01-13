package main

import (
	"fmt"
)

func main() {
	i := 55   //int
	j := 67.8 // float
	//sum := i + j //int + float not allowed
	sum := i + int(j) // converts j to an int
	fmt.Println(sum)
}
