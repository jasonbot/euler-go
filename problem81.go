/*
Problem 81

In the 5 by 5 matrix below, the minimal path sum from the top left to the
bottom right, by only moving to the right and down, is indicated in bold red
and is equal to 2427.


131	673	234	103	18
201	96	342	965	150
630	803	746	422	111
537	699	497	121	956
805	732	524	37	331

Find the minimal path sum, in matrix.txt (right click and 'Save Link/Target
As...'), a 31K text file containing a 80 by 80 matrix, from the top left to
the bottom right by only moving right and down.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	UNSET = 0
	LEFT  = 1
	UP    = 2
)

type RouteNode struct {
	minscore, cost uint64
	direction      uint8
}

func CreateNode(node_score uint64) RouteNode {
	return RouteNode{0, node_score, UNSET}
}

func main() {
	iohandle, err := os.Open("./matrix.txt")
	if err != nil {
		fmt.Printf("Failed to open matrix.txt\n")
		return
	}

	var matrix [][]RouteNode

	reader := bufio.NewReader(iohandle)

	for {
		if line, _, err := reader.ReadLine(); err != nil {
			break
		} else {
			row := make([]RouteNode, 0)
			lineString := string(line)
			nodeValues := strings.Split(lineString, ",")
			for index := 0; index < len(nodeValues); index++ {
				val, err := strconv.ParseUint(nodeValues[index], 10, 64)
				if err != nil {
					fmt.Printf("Error reading data\n")
					return
				}
				row = append(row, CreateNode(val))
			}
			matrix = append(matrix, row)
		}
	}

	maxrow := len(matrix)
	for rowindex := 0; rowindex < maxrow; rowindex++ {
		maxcol := len(matrix[rowindex])
		for colindex := 0; colindex < maxcol; colindex += 1 {
			newmax := matrix[rowindex][colindex].minscore + matrix[rowindex][colindex].cost
			if (colindex < maxcol - 1) {
				if (matrix[rowindex][colindex + 1].direction == UNSET || matrix[rowindex][colindex + 1].minscore > newmax) {
					matrix[rowindex][colindex + 1].minscore = newmax
					matrix[rowindex][colindex + 1].direction = LEFT
				}
			}
			if (rowindex < maxrow - 1) {
				if (matrix[rowindex + 1][colindex].direction == UNSET || matrix[rowindex + 1][colindex].minscore > newmax) {
					matrix[rowindex + 1][colindex].minscore = newmax
					matrix[rowindex + 1][colindex].direction = UP
				}
			}
		}
	}
}
