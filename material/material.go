package material

var (
	MaterialWool = &Material{Name: "Wool"}
	MaterialCotton = &Material{Name: "Cotton"}
	MaterialLinen = &Material{Name: "Linen"}
	MaterialFlannel = &Material{Name: "Flannel"}
	MaterialTweed = &Material{Name: "Tweed"}
)

type Material struct {
	Name string
}
