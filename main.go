package main

import (
	"fmt"
)

type ShapeType struct {
	Vec [][]int
}

func CopyTwoDmsArray(vec [][]int) (newVec [][]int) {
	newVec = make([][]int, len(vec))

	for i := range vec {
		newVec[i] = make([]int, len(vec[i]))
		for j := range vec[i] {
			newVec[i][j] = vec[i][j]
		}
	}
	return
}

func main() {
	vec := [][]int{
		{1,2,3,4},
		{2,0,0,0},
		{3,4,5,6},
	}
	newVec := CopyTwoDmsArray(vec)
	fmt.Printf("%+v\n",vec)
	newVec[0][1] = 10
	fmt.Printf("%+v\n",newVec)
}
