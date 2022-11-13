package main

import (
	"dsa/custom-type/setarray"
	"dsa/problems/array"
	"dsa/recursion/maze"
	"dsa/searching/bst"

	"github.com/shubham-gaur/logit"
)

func main() {
	app := "DSA Tutorials"
	logit.Init(&app)
	logit.Info.Println(logit.GetCurrentFunctionName(), "Set In Array::")
	setarray.RunSetInArray()
	logit.Info.Println(logit.GetCurrentFunctionName(), "Searching BST::")
	bst.RunBST()
	logit.Info.Println(logit.GetCurrentFunctionName(), "Integer Maze::")
	maze.RunIntMaze()
	logit.Info.Println(logit.GetCurrentFunctionName(), "Boolean Maze::")
	maze.RunBoolMaze()
	logit.Info.Println(logit.GetCurrentFunctionName(), "Matching Brackets::")
	array.RunCheckMatchingBrackets()
}
