package colour

import (
	"strings"
	"strconv"
)

type HexColour struct {
	red uint8
	green uint8
	blue uint8
}

func (c *HexColour) GetHexColour() *HexColour {
	return c
}

func (c *HexColour) R() uint8 {
	return c.red
}

func (c *HexColour) G() uint8 {
	return c.green
}

func (c *HexColour) B() uint8 {
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

	return &HexColour{
		red: uint8(r),
		green: uint8(g),
		blue: uint8(b) }, nil
}
