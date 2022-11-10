package maze

import (
	"github.com/shubham-gaur/logit"
)

func RunMaze() {
	logit.Info.Println(logit.GetCurrentFunctionName(), "Integer Maze::")
	m := InitIntMaze(3, 3)
	m.RightDown()
	m.RightDownSlant()

	logit.Info.Println(logit.GetCurrentFunctionName(), "Boolean Maze::")
	bm := InitBoolMaze(3, 3)
	bm.RightDownSlant()
}
