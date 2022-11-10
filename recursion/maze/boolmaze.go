package maze

import "fmt"

type boolmaze struct {
	rows int
	cols int
	val  [][]bool
}

type BoolMazeServicer interface {
	RightDownSlant()
}

func InitBoolMaze(r, c int) BoolMazeServicer {
	bm := boolmaze{rows: r, cols: c, val: make([][]bool, r, c)}
	for r := 0; r < bm.rows; r++ {
		bm.val[r] = make([]bool, 3)
		for c := 0; c < bm.cols; c++ {
			if r == 1 && c == 1 {
			} else {
				bm.val[r][c] = true
			}
		}
	}
	return &bm
}

func (bm *boolmaze) RightDownSlant() {
	fmt.Println(getRightDownSlantPathWithRestriction("", *bm, 0, 0))
	printRightDownSlantPathWithRestriction("", *bm, 0, 0)
}

func getRightDownSlantPathWithRestriction(path string, bm boolmaze, r, c int) []string {
	if r == bm.rows-1 && c == bm.cols-1 {
		var paths []string
		paths = append(paths, path)
		return paths
	}
	if !bm.val[r][c] {
		return []string{}
	}

	var paths []string
	if r < bm.rows-1 && c < bm.cols-1 {
		paths = append(paths, getRightDownSlantPathWithRestriction(path+"S", bm, r+1, c+1)...)
	}

	if r < bm.rows-1 {
		paths = append(paths, getRightDownSlantPathWithRestriction(path+"D", bm, r+1, c)...)
	}

	if c < bm.cols-1 {
		paths = append(paths, getRightDownSlantPathWithRestriction(path+"R", bm, r, c+1)...)
	}

	return paths
}

func printRightDownSlantPathWithRestriction(path string, bm boolmaze, r, c int) {
	if r == bm.rows-1 && c == bm.cols-1 {
		fmt.Println(path)
	}

	if !bm.val[r][c] {
		return
	}

	if r < bm.rows-1 && c < bm.cols-1 {
		printRightDownSlantPathWithRestriction(path+"S", bm, r+1, c+1)
	}

	if r < bm.rows-1 {
		printRightDownSlantPathWithRestriction(path+"D", bm, r+1, c)
	}

	if c < bm.cols-1 {
		printRightDownSlantPathWithRestriction(path+"R", bm, r, c+1)
	}
}
