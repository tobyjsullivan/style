package stylist

var (
	SeasonWinter = &Season{Name: "Winter"}
	SeasonSpring = &Season{Name: "Spring"}
	SeasonSummer = &Season{Name: "Summer"}
	SeasonAutumn = &Season{Name: "Autumn"}
)

type Season struct {
	Name string
}
