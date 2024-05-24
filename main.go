package main

import "fmt"

func main() {
	fmt.Println("start at main")
	fmt.Println("start at local")
	testRebase()
	fmt.Println("local 3")
	fmt.Println("main 3")
	fmt.Println("local 4")
	fmt.Println("main 4")
}

func testRebase() {
	fmt.Println("rebased")
}
