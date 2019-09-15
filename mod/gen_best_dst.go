package mod

import "fmt"

func GenAllPoints(mtype int, vec [][]int) {
	lenx := len(vec)
	leny := len(vec[0])
	shape := NewShape(lenx - 1,leny - 1,mtype)
	tetrisUnit := NewTetrisUnit(vec)
	fmt.Printf("%+v\n", shape.GetShape())
	fmt.Println("lenx:", lenx)
	fmt.Println("x:", tetrisUnit.GetLenX(), "Y:", tetrisUnit.GetLenY())
	var res ShapePointss
	for count := shape.Count(); count > 0; count-- {
		shape.Roate()
		for i := leny - 1; i >= 0; i-- {
			shape.PosY = i
			shape.Posx = lenx - 1
			fmt.Println("shape posx:", shape.Posx)
			if tetrisUnit.CheckAvailable(shape) {
				shape.Posx--
				fmt.Println(*shape)

			}
			if shape.Posx > 0 {
				shape.Posx++
				tetrisUnitTmp := tetrisUnit.Copy()
				fmt.Println(*shape)
				if tetrisUnitTmp.MergeShape(shape) {
					evaluate := NewEvaluate(shape, tetrisUnitTmp)
					res = append(res, &ShapePoints{
						MShape: *shape,
						Points: evaluate.GetScore(),
					})
					fmt.Printf("%+v\n", ShapePoints{
						MShape: *shape,
						Points: evaluate.GetScore(),
					})
				}
			}
		}
	}
	fmt.Printf("%+v\n", res)
}
