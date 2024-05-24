package main

import "fmt"

func main() {
	fmt.Println("main 1")
	fmt.Println("local 1")
	tryRebase()
}

func tryRebase() {
	fmt.Println("rebased")
}
