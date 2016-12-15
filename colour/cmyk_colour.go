package colour

import (
	"errors"
	"fmt"
)

type CMYKColour struct {
	cyan float32
	magenta float32
	yellow float32
	key float32
}

func (c *CMYKColour) RGBColour() *RGBColour {
	r := 255 * (1 - c.cyan) * (1 - c.key)
	g := 255 * (1 - c.magenta) * (1 - c.key)
	b := 255 * (1 - c.yellow) * (1 - c.key)

	return &RGBColour{
		red: uint8(r),
		green: uint8(g),
		blue: uint8(b)}
}

func (c *CMYKColour) HSLColour() *HSLColour {
	return c.RGBColour().HSLColour()
}

func ColourFromCMYK(c float32, m float32, y float32, k float32) (Colour, error) {
	if c > 1.0 {
		return nil, errors.New(fmt.Sprintf("Cyan value is out of range: %f", c))
	}

	if m > 1.0 {
		return nil, errors.New(fmt.Sprintf("Cyan value is out of range: %f", c))
	}

	if y > 1.0 {
		return nil, errors.New(fmt.Sprintf("Cyan value is out of range: %f", c))
	}

	if k > 1.0 {
		return nil, errors.New(fmt.Sprintf("Cyan value is out of range: %f", c))
	}

	return &CMYKColour{
		cyan: c,
		magenta: m,
		yellow: y,
		key: k}, nil
}

