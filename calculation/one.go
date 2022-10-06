package main

import (
  "fmt"
  "math"
)

func main() {
  var x, y float64
  var a, b, c, d, f, i float64

  x = 2
  y = 2

  a = math.Pow(x, 2)
  b = 3*a
  c = 10*y
  d = 1/b+10
  f = d + c + 7
  fmt.Println(f)
}