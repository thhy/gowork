package mod

import (
	"log"
	"fmt"
)

type TetrisUnit struct {
	Vec [][]int
	Lenx int
	Leny int
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

func NewTetrisUnit(vec [][]int) (this *TetrisUnit) {
	this = new(TetrisUnit)
	this.Vec = CopyTwoDmsArray(vec)
	this.Lenx = len(this.Vec)
	this.Leny = len(this.Vec[0])
	return
}

func (this *TetrisUnit) Copy() (newTe *TetrisUnit) {
	newTe = new(TetrisUnit)
	newTe.Lenx = this.Lenx
	newTe.Leny = this.Leny
	newTe.Vec = CopyTwoDmsArray(this.Vec)
	return
}

func (this *TetrisUnit) CheckAvailable(shape *Shape) bool {
	vec := shape.GetShape()
	for i := 0; i < len(vec); i++ {
		for j := 0; j < len(vec[i]); j++ {
			if vec[i][j] > 0 {
				if shape.Posx - i >= this.Lenx || shape.PosY - j >= this.Leny || shape.Posx - i < 0 || shape.PosY - j < 0 {
					fmt.Println("out of range", "shape posx:", shape.Posx + i, " shape posy:", shape.PosY + j, " this Lenx:", this.GetLenX(), " this lenY:", this.GetLenY())
					return false
				}
				if this.Vec[shape.Posx - i][shape.PosY - j] > 0 {
					return false
				}
			}
		}
	}
	return true
}

func (this *TetrisUnit) MergeShape(shape *Shape) bool {
	vec := shape.GetShape()
	if !this.CheckAvailable(shape) {
		return false
	}
	for i := 0; i < len(vec); i++ {
		for j := 0; j < len(vec[i]); j++ {
			if vec[i][j] > 0 {
				fmt.Println("x:", shape.Posx - i, " Y:", shape.PosY - j)
				if this.Vec[shape.Posx - i][shape.PosY - j] > 0 {
					log.Fatal("invalid pos, please check available")
				}
				this.Vec[shape.Posx - i][shape.PosY - j] = vec[i][j]
			}
		}
	}
	return true
}

func (this *TetrisUnit) GetLenX() int {
	return this.Lenx
}
func (this *TetrisUnit) GetLenY() int {
	return this.Leny
}