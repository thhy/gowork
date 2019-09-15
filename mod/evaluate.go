package mod

type Evaluate struct {
	shape *Shape
	tetrisUnit *TetrisUnit
}

func NewEvaluate(shape *Shape, unit *TetrisUnit) (this *Evaluate) {
	this = new(Evaluate)
	this.shape = shape
	this.tetrisUnit = unit
	return
}

/*最大高度*/
func (this *Evaluate) GetMaxHeigh() int {
	var maxHeigh = 0
	for i := this.tetrisUnit.GetLenX() - 1; i >= 0; i-- {
		for j := range this.tetrisUnit.Vec[i] {
			if this.tetrisUnit.Vec[i][j] > 0 {
				return i
			}
		}
	}
	return maxHeigh
}

/*贡献块*/
func (this *Evaluate) GetRE() int {
	var delLine = 0
	var contributeCount = 0
	for i := 0; i < this.tetrisUnit.GetLenX(); i++ {
		var isDel = true
		for j := 0; j < this.tetrisUnit.GetLenY(); j++ {
			if this.tetrisUnit.Vec[i][j] == 0 {
				isDel = false
			}
		}
		if isDel {
			delLine++
			vec := this.shape.GetShape()
			for k := range vec[this.shape.Posx - i] {
				if k > 0 {
					contributeCount++
				}
			}
		}
	}
	return delLine * contributeCount
}

/*行变换*/
func (this *Evaluate) GetRT() int {
	changeLine := 0
	for i := this.tetrisUnit.GetLenX() - 1; i >= 0; i-- {
		preColor := this.tetrisUnit.Vec[i][0]
		for j := 1; j < this.tetrisUnit.GetLenY(); j++ {
			if this.tetrisUnit.Vec[i][j] != preColor {
				changeLine++
			}
		}
	}
	return changeLine
}

/*列变换*/
func (this *Evaluate) GetCT() int {
	changeColumn := 0
	for i := 0; i < this.tetrisUnit.GetLenY(); i++ {
		preColor := this.tetrisUnit.Vec[this.tetrisUnit.GetLenX() - 1][i]
		for j := this.tetrisUnit.GetLenX() - 2; j >= 0; j-- {
			if this.tetrisUnit.Vec[j][i] != preColor {
				changeColumn++
			}
		}
	}
	return changeColumn
}

/*空洞数*/
func (this *Evaluate) GetHoleCount() int  {
	holeCount := 0
	for i := 0; i < this.tetrisUnit.GetLenY() - 1; i++ {
		line := this.tetrisUnit.GetLenX() - 1
		for ; line >= 0; line-- {
			if this.tetrisUnit.Vec[line][i] > 0 {
				break
			}
		}
		for ; line >= 0; line-- {
			if this.tetrisUnit.Vec[line][i] == 0 {
				holeCount++
			}
		}
	}
	return holeCount
}

/*井深评分*/
func (this *Evaluate) GetWellCount() int {
	wellCount := 0
	//计算左边的
	thisHole := 0
	wellDepth := 0
	for i := this.tetrisUnit.GetLenX() - 1; i >= 0; i-- {
		if this.tetrisUnit.Vec[i][0] == 0 && this.tetrisUnit.Vec[i][1] > 0 {
			thisHole++
		} else {
			wellDepth += thisHole * (thisHole + 1) / 2
			thisHole = 0
		}
	}
	wellDepth += thisHole * (thisHole + 1) / 2
	wellCount += wellDepth

	for column := 1; column < this.tetrisUnit.GetLenY() - 1; column++ {
		thisHole = 0
		wellDepth = 0
		for row := this.tetrisUnit.GetLenX() - 1; row >= 0; row-- {
			if this.tetrisUnit.Vec[row][column] == 0 && this.tetrisUnit.Vec[row][column + 1] > 0 && this.tetrisUnit.Vec[row][column - 1] > 0 {
				thisHole++
			} else {
				wellDepth += thisHole * (thisHole + 1) / 2
				thisHole = 0
			}
		}
	}
	wellCount += wellDepth

	thisHole = 0
	wellDepth = 0
	for i := this.tetrisUnit.GetLenX() - 1; i >= 0; i-- {
		if this.tetrisUnit.Vec[i][this.tetrisUnit.GetLenY() - 2] > 0 && this.tetrisUnit.Vec[i][this.tetrisUnit.GetLenY() - 1] == 0 {
			thisHole++
		} else {
			wellDepth += thisHole * (thisHole + 1) / 2
			thisHole = 0
		}
	}
	wellDepth += thisHole * (thisHole + 1) / 2
	wellCount += wellDepth
	return wellCount
}

func (this *Evaluate) GetScore() float32 {

	return -4.500158825082766 * float32(this.GetMaxHeigh()) +
		3.4181268101392694 * float32(this.GetRE()) +
		-3.2178882868487753 * float32(this.GetRT()) +
		-9.348695305445199 * float32(this.GetCT()) +
		-7.899265427351652 * float32(this.GetHoleCount()) +
		-3.3855972247263626 * float32(this.GetWellCount())
}