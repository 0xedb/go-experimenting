package main

import "fmt"

var a, b, c, d int
var bc int8 = 127
var ui uint = 39

const FACTOR = 339.82

func square(a int) int {
	return a * a
}

func something_and_two(w string) (string, int) {
	return w, 2
}

func hundred() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func hundred_while() {
	for ui < 50 {
		ui = ui + 1
		fmt.Println(ui)
		if ui == 50 {
			fmt.Println("My job is done here")
		} else if ui == 40 {
			fmt.Println("just @ 40")
		} else {
			fmt.Println("***")
		}
	}
}

func switcher(a int) {
	switch a {
	case 0:
		fmt.Println("Nothing, absolutely nothing!!!")
	default:
		fmt.Println("something fishy!!!")
	}

}

// pointer
var number_of_apples *int

var talkatives [10]int

type Node struct {
	data int
	next *Node
}

var node *Node

func main() {
	// fmt.Println("Hello World")
	// fmt.Println(square(100))
	// fmt.Println(something_and_two("Go home today for Christmas"))
	// hundred()
	// defer hundred_while()
	// switcher(1)
	// number_of_apples = &b
	// *number_of_apples = 10
	// fmt.Println(*number_of_apples)
	// fmt.Println(b)
	// b = 200
	// fmt.Println(*number_of_apples, number_of_apples)
	// fmt.Println(b, &b)
	// fmt.Println(Node{10, nil})
	// talkatives[0] = 404
  // fmt.Println(talkatives[0])

  // names := [2]string{"bruno", "edoh"}
  // fmt.Println(names)

  // many := []bool{true, false, true}
  // fmt.Println(many)
  // var n []int
  // w := n == nil
  // fmt.Println(w)

}
