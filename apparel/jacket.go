package apparel

import "github.com/tobyjsullivan/style/colour"

var (
	LapelNotch = &JacketLapel{Name: "Notch"}
	LapelShawl = &JacketLapel{Name: "Shawl"}
	LapelPeak = &JacketLapel{Name: "Peak"}
)

type Jacket struct {
	Colour colour.Colour
	LapelType *JacketLapel
}

type JacketLapel struct {
	Name string
}
