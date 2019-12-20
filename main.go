package main

import "fmt"
// import "math"

func square(a int) int {
  return a * a
}

func something_and_two(w string) (string, int) {
  return w, 2
}

func main() {
  fmt.Println("Hello World")
  fmt.Println(square(100))
  fmt.Println(something_and_two("Go home today for Christmas"))
}