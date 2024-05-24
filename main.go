package main

import "fmt"

func main() {
	fmt.Println("start at main")
	localBranchOne()
	mainFunc()
}

func mainFunc() {
	fmt.Println("created at main")
}

func localBranchOne() {
	fmt.Println("created new branch")
}
