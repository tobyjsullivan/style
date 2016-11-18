package material

var (
	PatternTwill = &Pattern{Name: "Twill"}
	PatternHoundstooth = &Pattern{Name: "Houndstooth"}
	PatternHerringbone = &Pattern{Name: "Herringbone"}
	PatternGlenurquhartCheck = &Pattern{Name: "Glenurquhart Check"}
	PatternWindowpaneCheck = &Pattern{Name: "Windowpane Check"}
	PatternPinstripe = &Pattern{Name: "Pinstripe"}
)

type Pattern struct {
	Name string
}
