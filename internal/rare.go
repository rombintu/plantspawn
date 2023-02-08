package internal

import (
	"math/rand"
	"time"
)

// Common, Rare, Epic, Legendary

const (
	Common    string = "Common"
	Rare      string = "Rare"
	Epic      string = "Epic"
	Legendary string = "Legendary"
	RangeRare int    = 1000
)

var (
	CommonColor    Color = NewColor(102, 153, 153) // 74.5%
	RareColor      Color = NewColor(0, 51, 204)    // 20%
	EpicColor      Color = NewColor(153, 102, 204) // 5%
	LegendaryColor Color = NewColor(255, 153, 0)   // 0.5%
)

type Rarity struct {
	Title string
	Color Color
}

var (
	CommonRariry    Rarity = Rarity{Title: Common, Color: CommonColor}
	RareRariry      Rarity = Rarity{Title: Rare, Color: RareColor}
	EpicRariry      Rarity = Rarity{Title: Epic, Color: EpicColor}
	LegendaryRariry Rarity = Rarity{Title: Legendary, Color: LegendaryColor}
)

func NewRarity() Rarity {
	rand.Seed(time.Now().UnixNano())
	rare := rand.Intn(RangeRare)
	if rare > 995 {
		return LegendaryRariry
	} else if rare > 945 {
		return EpicRariry
	} else if rare > 800 {
		return RareRariry
	} else {
		return CommonRariry
	}
}
