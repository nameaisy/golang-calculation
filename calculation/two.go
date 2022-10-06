package main

import (
  "fmt"
  "math"
)

func main() {
	var x float64
	var a, b, c, d, e, f, h float64
	//fmt.Scanln(&x)
	x = 1.6
	a = math.Pow(x, 3)
	b = math.Pow(x, 4)
	c = math.Pow(x, 2)
	e = 3.0 * x
	f = -3.0 * c
	h = 4.0
	d = (a+e) / (b+f+h)
	fmt.Println(d)
}