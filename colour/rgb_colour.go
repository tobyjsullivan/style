package colour

import (
	"strings"
	"strconv"
)

type RGBColour struct {
	red uint8
	green uint8
	blue uint8
}

func (c *RGBColour) RGBColour() *RGBColour {
	return c
}

func (c *RGBColour) R() uint8 {
	return c.red
}

func (c *RGBColour) G() uint8 {
	return c.green
}

func (c *RGBColour) B() uint8 {
	return c.blue
}

func ColourFromHex(hexString string) (Colour, error) {
	hexString = strings.ToUpper(hexString)

	r, err := strconv.ParseUint(hexString[0:2], 16, 8)
	if err != nil {
		return nil, err
	}

	g, err := strconv.ParseUint(hexString[2:4], 16, 8)
	if err != nil {
		return nil, err
	}

	b, err := strconv.ParseUint(hexString[4:6], 16, 8)
	if err != nil {
		return nil, err
	}

	return &RGBColour{
		red: uint8(r),
		green: uint8(g),
		blue: uint8(b) }, nil
}
