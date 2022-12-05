package main

import "fmt"


/*

A very hungry rabbit is placed in the center of a garden, represented by a
rectangular N x M 2D matrix. The values of the matrix will represent numbers of
carrots available to the rabbit in each square of the garden. If the garden
does not have an exact center, the rabbit should start in the square closest to
the center with the highest carrot count.

On a given turn, the rabbit will eat the carrots available on the square that
it is on, and then move up, down, left, or right, choosing the square that has
the most carrots. If there are no carrots left on any of the adjacent squares,
the rabbit will go to sleep. You may assume that the rabbit will never have to
choose between two squares with the same number of carrots.

Write a function which takes a garden matrix and returns the number of carrots
the rabbit eats. You may assume the matrix is rectangular with at least 1 row
and 1 column, and that it is populated with non-negative integers.

For example,
[[5, 7, 8, 6, 3],
[0, 0, 7, 0, 4],
[4, 6, 3, 4, 9],
[3, 1, 0, 5, 8]]

Should return
27

*/


func main () {

    g := [][]int {
        {5, 7, 8, 6, 3},
        {0, 0, 7, 0, 4},
        {4, 6, 3, 4, 9},
        {3, 1, 0, 5, 8},
    }

    fmt.Println(max_carrots_consumed(g))
}

func max_carrots_consumed(g [][]int) int {

    start_row, start_col := start_position(g)
    return explore_grid(g, start_row, start_col)
}

func start_position(g [][]int) (int, int) {
    rows := len(g)
    cols := len(g[0])
    // starting position of rabbit
    start_row := -1
    start_col := -1
    if rows%2 == 1 && cols%2 == 1 {
        // center present, so use that as the start position
        start_row = rows/2
        start_col = cols/2
    } else {
        r := []int{}
        if rows%2 == 0 {
            r = append(r, rows/2 - 1)
            r = append(r, rows/2)
            r = append(r, rows/2 + 1)
        } else {
            r = append(r, rows/2)
        }
        c := []int{}
        if cols%2 == 0 {
            c = append(c, cols/2 - 1)
            c = append(c, cols/2)
            c = append(c, cols/2 + 1)
        } else {
           c  = append(c, cols/2)
        }

        // now searh in the grid for max carrot square
        start_row = r[0]
        start_col = c[0]
        for i := 0 ; i < len(r); i++ {
            for j := 0 ; j < len(c); j++ {
                tr := r[i]
                tc := c[j]
                if tr < 0 || tr > rows {
                    continue
                }
                if tc < 0 || tc > cols {
                    continue
                }
                if g[tr][tc] > g[start_row][start_col] {
                    start_row = tr
                    start_col = tc
                }
            }
        }
    }

    return start_row, start_col

}

func explore_grid (g [][]int, start_row int, start_col int) int {

    nrows := len(g)
    ncols := len(g[0])
    //fmt.Println(start_row, start_col)

    visited := make([][]bool, nrows)
    for i := 0 ; i < nrows; i++ {
        visited[i] = make([]bool, ncols)
    }

    current_row := start_row
    current_col  := start_col

    // consume current square and move to adjacent square with max carrots
    carrots_consumed := 0
    for ;; {
        carrots_consumed += g[current_row][current_col]
        //fmt.Println("C", current_row, current_col, g[current_row][current_col], carrots_consumed)
        visited[current_row][current_col] = true

        rdir := []int{current_row-1, current_row+1}
        cdir := []int{current_col-1, current_col+1}

        max := -1
        mr := -1
        mc := -1

        for _, c := range cdir {
            if c < 0 || c > ncols {
                continue
            }
            if visited[current_row][c] == true {
                continue
            }
            if g[current_row][c] > max {
                mr = current_row
                mc = c
                max = g[current_row][c]
            }
        }

        for _, r := range rdir {
            if r < 0 || r > nrows {
                continue
            }
            if visited[r][current_col] == true {
                continue
            }
            if g[r][current_col]  > max {
                mr = r
                mc = current_col
                max = g[r][current_col]
            }
        }
        if max == -1 || max == 0 {
            // no more carrots to consume
            break
        }
        current_row = mr
        current_col = mc
    }

    return carrots_consumed
}


