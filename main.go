package main

import "fmt"

func main() {
	fmt.Println("start at main")
	localBranchOne()
}

func localBranchOne() {
	fmt.Println("created new branch")
}
