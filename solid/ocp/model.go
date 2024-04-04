package solid

import "fmt"

type Color int
type Size int

const (
	Red Color = iota
	Green
	Blue
	Yellow
	White
	Black
)

const (
	Small Size = iota
	Medium
	Large
	XLarge
	XXLarge
)

var colorMap = map[Color]string{
	Red:    "Red",
	Green:  "Green",
	Blue:   "Blue",
	Yellow: "Yellow",
	White:  "White",
	Black:  "Black",
}

var sizeMap = map[Size]string{
	Small:   "Small",
	Large:   "Large",
	XLarge:  "XLarge",
	XXLarge: "XXLarge",
}

type Shirt struct {
	Color Color
	Mrp   float64
	Size  Size
}

func (s *Shirt) String() string {
	return fmt.Sprintf("\t \t \t color:%s size:%s MRP:%f ", colorMap[s.Color], sizeMap[s.Size], s.Mrp)
}
