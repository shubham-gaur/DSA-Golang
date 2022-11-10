package maze

func RunIntMaze() {
	m := InitIntMaze(3, 3)
	m.RightDown()
	m.RightDownSlant()
}

func RunBoolMaze() {
	bm := InitBoolMaze(3, 3)
	bm.RightDownSlant()
}
