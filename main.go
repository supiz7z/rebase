package main

import "fmt"

func main() {
	fmt.Println("main 1")
	fmt.Println("local 1")
	tryRebase()
	tryRebaseToMain()
	victory()
}

func tryRebase() {
	fmt.Println("rebased")
}

func tryRebaseToMain() {
	fmt.Println("to main")
}

func victory() {
	fmt.Println("pidorasina")
}
