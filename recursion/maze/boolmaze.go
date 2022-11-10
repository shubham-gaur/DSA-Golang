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
	fmt.Println("Boolean Matrix")
	bm := boolmaze{rows: r, cols: c, val: make([][]bool, r, c)}
	for r := 0; r < bm.rows; r++ {
		fmt.Printf("[")
		bm.val[r] = make([]bool, c)
		for c := 0; c < bm.cols; c++ {
			if r == 1 && c == 1 {
			} else {
				bm.val[r][c] = true
			}
			fmt.Printf("%-5v ", bm.val[r][c])
		}
		fmt.Println("]")
	}
	fmt.Println()
	return &bm
}

func (bm *boolmaze) RightDownSlant() {
	fmt.Println("Path List Right Down Slant:\n\t", getRightDownSlantPathWithRestriction("", *bm, 0, 0))
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
