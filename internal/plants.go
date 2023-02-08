package internal

type Plant struct {
	Title string
	Color string
	Owner int
}

// Plant Constructor
func NewPlant() *Plant {
	return &Plant{}
}

func (p *Plant) GetColor() {

}
