package maze

import "fmt"

type MazeServicer interface {
	RightDown()
	RightDownSlant()
}

type maze struct {
	rows int
	cols int
}

func InitIntMaze(r, c int) MazeServicer {
	return &maze{
		rows: r,
		cols: c,
	}
}

func (m *maze) RightDown() {
	fmt.Println("Count: ", countRightDown(m.rows, m.cols))
	printRightDownPath("", m.rows, m.cols)
	fmt.Println(saveRightDownPath("", m.rows, m.cols))
}

func (m *maze) RightDownSlant() {
	fmt.Println(saveRightDownSlantPath("", m.rows, m.cols))
}

func countRightDown(r, c int) int {
	if r == 1 || c == 1 {
		return 1
	}
	left := countRightDown(r-1, c)
	right := countRightDown(r, c-1)
	return left + right
}

func printRightDownPath(path string, r, c int) {
	if r == 1 && c == 1 {
		fmt.Println(path)
		return
	}
	if r > 1 {
		printRightDownPath(path+"D", r-1, c)
	}

	if c > 1 {
		printRightDownPath(path+"R", r, c-1)
	}
}

func saveRightDownPath(path string, r, c int) []string {
	if r == 1 && c == 1 {
		var paths []string
		paths = append(paths, path)
		return paths
	}

	var paths []string

	if r > 1 {
		paths = append(paths, saveRightDownPath(path+"D", r-1, c)...)
	}

	if c > 1 {
		paths = append(paths, saveRightDownPath(path+"R", r, c-1)...)
	}
	return paths
}

func saveRightDownSlantPath(path string, r, c int) []string {
	if r == 1 && c == 1 {
		var paths []string
		paths = append(paths, path)
		return paths
	}

	var paths []string

	if r > 1 && c > 1 {
		paths = append(paths, saveRightDownSlantPath(path+"S", r-1, c-1)...)
	}

	if r > 1 {
		paths = append(paths, saveRightDownSlantPath(path+"D", r-1, c)...)
	}

	if c > 1 {
		paths = append(paths, saveRightDownSlantPath(path+"R", r, c-1)...)
	}
	return paths
}
