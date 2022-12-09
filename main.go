package main

import (
	"github.com/shubham-gaur/dsa-golang/custom-type/setarray"
	"github.com/shubham-gaur/dsa-golang/datastructures"
	"github.com/shubham-gaur/dsa-golang/recursion/maze"
	"github.com/shubham-gaur/dsa-golang/searching/bst"

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
	logit.Info.Println(logit.GetCurrentFunctionName(), "Graph::")
	datastructures.PrintEgDirectedGraph()
}
