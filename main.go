package main

import (
	"dsa/recursion/maze"
	"dsa/searching/bst"

	"github.com/shubham-gaur/logit"
)

func main() {
	app := "DSA Tutorials"
	logit.Init(&app)
	bst.RunBST()
	maze.RunMaze()
}
