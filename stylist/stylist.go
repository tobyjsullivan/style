package stylist

import (
	"github.com/tobyjsullivan/style/colour"
	"fmt"
	"math"
	"github.com/tobyjsullivan/style/material"
)

func EvaluateOutfit(o *Outfit) float64 {
	factor1 := computeContrast(o.Shirt.Colour, o.Trousers.Colour)
	return math.Abs(factor1)
}

func computeContrast(c1 colour.Colour, c2 colour.Colour) float64 {
	q1 := calcQIYContrast(c1.GetHexColour())
	q2 := calcQIYContrast(c2.GetHexColour())

	println(fmt.Sprintf("Computed QIY's: %f, %f", q1, q2))

	return q1 - q2
}

func calcQIYContrast(hc *colour.HexColour) float64 {
	r := uint32(hc.R())
	g := uint32(hc.G())
	b := uint32(hc.B())
	return float64((r * 299) + (g * 587) + (b * 114))/1000
}

func materialMatchesSeason(m *material.Material, s *Season) bool {
	switch m {
	case material.MaterialWool:
		return true
	case material.MaterialFlannel:
		return s == SeasonWinter
	case material.MaterialCotton:
		return s == SeasonSpring || s == SeasonSummer || s == SeasonAutumn
	case material.MaterialLinen:
		return s  == SeasonSummer
	case material.MaterialTweed:
		return s == SeasonWinter || s == SeasonAutumn
	}

	return true
}