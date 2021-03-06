package galaxy

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Jest0r/starex_go/coords"
)

const (
	LumMin  = 10.0
	LumExp  = 0.20
	LumMult = 15
)

type Color struct {
	R int32
	G int32
	B int32
	A int32
}

type System struct {
	CenterObject *CenterObject
	Lum          float64 `json:"lum"`
	Colorstr     string  `json:"color"`
	Coords       coords.CoordsI16
	Color        Color
}

// to satisfy kdtree interface
func (s System) Dimensions() int {
	return 3
}

// to satisfy kdtree interface
func (s System) Dimension(i int) float64 {
	switch i {
	case 0:
		return float64(s.Coords.X)
	case 1:
		return float64(s.Coords.Y)
	default:
		return float64(s.Coords.Z)
	}
}

func (s *System) print() {
	fmt.Printf("System - Coords %v", s.Coords)
}

func (s *System) PlaceCenterObject(co *CenterObject) {
	s.CenterObject = co
}

func (s *System) SetColor(colorstr string, lum float64) {
	// Error handling - Catching missing colors
	if len(colorstr) < 5 {
		s.Color.R = 200
		s.Color.G = 0
		s.Color.B = 200
		s.Lum = 100000
		s.Color.A = 50
	} else {
		r, _ := strconv.ParseInt(colorstr[1:3], 16, 16)
		g, _ := strconv.ParseInt(colorstr[3:5], 16, 16)
		b, _ := strconv.ParseInt(colorstr[5:7], 16, 16)
		s.Color.R = int32(r)
		s.Color.G = int32(g)
		s.Color.B = int32(b)
		s.Lum = lum

		alpha := math.Pow(lum, LumExp) * LumMult
		alpha = math.Max(alpha, LumMin)
		s.Color.A = int32(math.Min(alpha, 255))
	}
}
