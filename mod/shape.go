package mod

const (
	Lshape = iota
	Jshape
)

type ShapeType struct {
	Vec [][][]int
}

func (this *ShapeType) Count() int {
	return len(this.Vec)
}

func (this *ShapeType) GetIdx(idx int) [][]int {
	if len(this.Vec) > idx && idx >= 0 {
		return this.Vec[idx]
	}
	return nilShape
}

var shapeTypes map[int]*ShapeType
var nilShape [][]int

func init() {
	shapeTypes = make(map[int]*ShapeType)
	shapeTypes[Lshape] = &ShapeType{Vec: [][][]int{
		{
			{1, 0},
			{1, 0},
			{1, 1},
		},
		{
			{1,1,1},
			{0,0,1},
		},

	}}
	shapeTypes[Jshape] = &ShapeType{Vec: [][][]int{
		{
			{},
		},
	}}
}

type Shape struct {
	Posx int
	PosY int
	shapeType *ShapeType
	Idx int
	MType int
}

func NewShape(posx, posy, shape int) (this *Shape) {
	this = new(Shape)
	this.Posx = posx
	this.PosY = posy
	this.Idx = 0
	this.MType = shape
	this.shapeType = shapeTypes[this.MType]
	return
}

func (this *Shape) Copy() *Shape  {
	res := new(Shape)
	res.Posx = this.Posx
	res.PosY = this.PosY
	res.shapeType = this.shapeType
	res.Idx = this.Idx
	res.MType = this.MType
	return res
}

func (this *Shape) GetShape() [][]int {
	count := this.shapeType.Count()
	if count > this.Idx && this.Idx >= 0 {
		return this.shapeType.GetIdx(this.Idx)
	}
	return nil
}

func (this *Shape) Count() int{
	return len(shapeTypes[this.MType].Vec)
}

func (this *Shape) Roate() {
	count := this.Count()
	this.Idx = (this.Idx + 1)%count
}

type ShapePoints struct {
	MShape Shape
	Points float32
	MTetrisUnit	TetrisUnit
}

type ShapePointss []*ShapePoints

func (this ShapePointss) Len() int{
	return len(this)
}

func (this ShapePointss) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this ShapePointss) Less(i, j int) bool{
	return this[i].Points > this[j].Points
}

