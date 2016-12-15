package colour

type Colour interface {
	RGBColour() *RGBColour
	HSLColour() *HSLColour
}
