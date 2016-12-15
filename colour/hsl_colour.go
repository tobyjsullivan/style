package colour

type HSLColour struct {
	hue uint16
	saturation float64
	lightness float64
}

func (c *HSLColour) Hue() uint16 {
	return c.hue
}

func (c *HSLColour) Saturation() float64 {
	return c.saturation
}

func (c *HSLColour) Lightness() float64 {
	return c.lightness
}

//func ColourFromHSL(h float64, s float64, l float64) (Colour, error) {
//	return &HSLColour{
//		hue: h,
//		saturation: s,
//		lightness: l}, nil
//}
