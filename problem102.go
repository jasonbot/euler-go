/*
Problem 102.
   
Three distinct points are plotted at random on a Cartesian plane, for which -1000  x, y  1000, such that a triangle is formed.

Consider the following two triangles:

A(-340,495), B(-153,-910), C(835,-947)

X(-175,41), Y(-421,-714), Z(574,-645)

It can be verified that triangle ABC contains the origin, whereas triangle XYZ does not.

Using triangles.txt (right click and 'Save Link/Target As...'), a 27K text file containing the co-ordinates of one thousand "random" triangles, find the number of triangles for which the interior contains the origin.

NOTE: The first two examples in the file represent the triangles in the example given above.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y float64
}

type LineSegment struct {
	P1, P2 Point
}

func (self *LineSegment) OnRight() bool {
	x1, x2, y1, y2 := self.P1.X, self.P2.X, self.P1.Y, self.P2.Y

	slope := (y2 - y1) / (x2 - x1)
	intercept := y1 - (x1 * slope)

	return intercept < 0
}

type Triangle struct {
	Lines [3]LineSegment
}

func (self *Triangle) CrossesOrigin() bool {
	crossings := 0
	if self.Lines[0].OnRight() {
		crossings += 1
	}
	if self.Lines[1].OnRight() {
		crossings += 1
	}
	if self.Lines[2].OnRight() {
		crossings += 1
	}

	return (crossings % 2) == 0
}

func MakeTriangle(point_array [6]float64) Triangle {

	p1, p2, p3 := Point{point_array[0], point_array[1]}, Point{point_array[2], point_array[3]}, Point{point_array[4], point_array[5]}
	l1, l2, l3 := LineSegment{p1, p2}, LineSegment{p2, p3}, LineSegment{p3, p1}
	triangle := Triangle{[3]LineSegment{l1, l2, l3}}

	return triangle
}

func main() {
	iohandle, err := os.Open("./triangles.txt")
	if err != nil {
		fmt.Printf("Failed to open triangles.txt\n")
		return
	}

	reader := bufio.NewReader(iohandle)
	crossed_triangles := 0

	for {
		if line, _, err := reader.ReadLine(); err != nil {
			break
		} else {
			lineString := string(line)
			coordinates := strings.Split(lineString, ",")
			if len(coordinates) != 6 {
				// fmt.Printf("Found a badly-formed line")
				return
			}

			var items [6]float64

			for index := 0; index < 6; index++ {
				items[index], err = strconv.ParseFloat(coordinates[index], 64)
				if err != nil {
					fmt.Printf("Found a badly-formed number")
					return
				}
			}

			triangle := MakeTriangle(items)
			if triangle.CrossesOrigin() {
				//fmt.Printf("Item crosses %s\n", line)
				crossed_triangles += 1
			}
		}
	}

	fmt.Printf("%v\n", crossed_triangles)
}
