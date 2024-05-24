package main

import "fmt"

func main() {
	fmt.Println("start at main")
	fmt.Println("start at local")
	testRebase()
}

func testRebase() {
	fmt.Println("rebased")
}
