package mod

import (
	"fmt"
	"sort"
)

func GenAllPoints(mtype int, vec [][]int) {
	lenx := len(vec)
	leny := len(vec[0])
	shape := NewShape(lenx - 1,leny - 1,mtype)
	tetrisUnit := NewTetrisUnit(vec)
	var res ShapePointss
	for count := shape.Count(); count > 0; count-- {
		shape.Roate()
		for i := leny - 1; i >= 0; i-- {
			shape.PosY = i
			shape.Posx = lenx - 1
			for tetrisUnit.CheckAvailable(shape) {
				shape.Posx--
			}
			if shape.Posx < lenx - 1 {
				shape.Posx++
				tetrisUnitTmp := tetrisUnit.Copy()
				if tetrisUnitTmp.MergeShape(shape) {
					evaluate := NewEvaluate(shape, tetrisUnitTmp)
					res = append(res, &ShapePoints{
						MShape: *shape,
						Points: evaluate.GetScore(),
						MTetrisUnit: *tetrisUnitTmp,
					})
				}
			}
		}
	}
	sort.Sort(res)
	fmt.Printf("%+v\n", res)
	fmt.Println(res[0].MShape)
	fmt.Println(res[0].MTetrisUnit)
	fmt.Println(res[0].MShape.GetShape())
	fmt.Println(res[0].Points)
}
