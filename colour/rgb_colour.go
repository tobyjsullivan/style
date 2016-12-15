package colour

import (
	"strings"
	"strconv"
	"math"
)

type RGBColour struct {
	red uint8
	green uint8
	blue uint8
}

func (c *RGBColour) RGBColour() *RGBColour {
	return c
}

func (c *RGBColour) HSLColour() *HSLColour  {
	rPrime := float64(c.red) / 255
	gPrime := float64(c.green) / 255
	bPrime := float64(c.blue) / 255

	cMax := math.Max(rPrime, math.Max(gPrime, bPrime))
	cMin := math.Min(rPrime, math.Min(gPrime, bPrime))

	delta := cMax - cMin

	lightness := (cMax + cMin) / 2

	saturation := 0.000
	eps := 0.0001
	if delta > eps {
		saturation = delta / (1 - math.Abs(2*lightness - 1))
	}

	hue := uint32(0)
	println("Delta: ", delta)
	println("rPrime: ", rPrime)
	println("gPrime: ", gPrime)
	println("bPrime: ", bPrime)
	if delta > eps && math.Abs(cMax - rPrime) < eps {
		hue = 60 * (uint32((gPrime - bPrime) / delta) % 6)
	} else if delta > eps && math.Abs(cMax - gPrime) < eps {
		hue = uint32(60 * (((bPrime - rPrime) / delta) + 2))
	} else if delta > eps && math.Abs(cMax - bPrime) < eps {
		hue = uint32(60 * (((rPrime - gPrime) / delta) + 4))
	}

	return &HSLColour{
		hue: uint16(hue),
		saturation: saturation,
		lightness: lightness}
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
