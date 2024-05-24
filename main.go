package main

import "fmt"

func main() {
	fmt.Println("start at main")
	fmt.Println("start at local")
	testRebase()
	fmt.Println("local 3")
	fmt.Println("main 3")
}

func testRebase() {
	fmt.Println("rebased")
}
