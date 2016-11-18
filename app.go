package main

import (
	"github.com/tobyjsullivan/style/apparel"
	"github.com/tobyjsullivan/style/colour"
	"github.com/tobyjsullivan/style/stylist"
	"fmt"
)

func main() {
	println("Bleep bloop.")

	blue, err := colour.ColourFromHex("000099")
	if err != nil {
		panic(err)
	}

	blueShirt := &apparel.Shirt{
		Colour: blue,
	}

	khaki, err := colour.ColourFromHex("cc9900")
	if err != nil {
		panic(err)
	}

	khakiPants := &apparel.Trousers{
		Colour: khaki,
	}

	outfitA := &stylist.Outfit{
		Shirt: blueShirt,
		Trousers: khakiPants,
	}

	res := stylist.EvaluateOutfit(outfitA)
	println(fmt.Sprintf("Outfit evaluation (A): %.02f", res))

	black, err := colour.ColourFromHex("000000")
	if err != nil {
		panic(err)
	}

	blackShirt := &apparel.Shirt{
		Colour: black,
	}

	white, err := colour.ColourFromHex("ffffff")
	if err != nil {
		panic(err)
	}

	whitePants := &apparel.Trousers{
		Colour: white,
	}

	outfitB := &stylist.Outfit{
		Shirt: blackShirt,
		Trousers: whitePants,
	}

	res = stylist.EvaluateOutfit(outfitB)
	println(fmt.Sprintf("Outfit evaluation (B): %.02f", res))

}
