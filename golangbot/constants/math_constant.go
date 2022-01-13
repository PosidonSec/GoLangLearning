package main

import (
	"math"
)

func main() {
	var a = math.Sqrt(2)   // allowed
	const b = math.Sqrt(2) //not allowed  this is because value is not known at compilation of the code
}
