package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/rombintu/plantspawn/internal"
)

func TestColors(t *testing.T) {
	c := internal.NewColor(255, 0, 255)
	fmt.Println(c.HexString)
	fmt.Println(internal.ParseHexColor(c.HexString))
}

func TestRareConstColors(t *testing.T) {
	fmt.Println(
		internal.CommonColor.HexString, internal.RareColor.HexString,
		internal.EpicColor.HexString, internal.LegendaryColor.HexString,
	)
}

func TestNewRarity(t *testing.T) {
	var common int
	var rare int
	var epic int
	var leg int
	summ := 10000
	for i := 0; i < summ; i++ {
		r := internal.NewRarity()
		if r.Title == internal.Common {
			common++
		} else if r.Title == internal.Rare {
			rare++
		} else if r.Title == internal.Epic {
			epic++
		} else if r.Title == internal.Legendary {
			leg++
		}
	}
	fmt.Printf(
		"Common: %d\nRare: %d\nEpic: %d\nLegendary: %d\nSUMMA: %d",
		common, rare, epic, leg, summ,
	)
}

func TestColorsCli(t *testing.T) {
	for _, c := range []internal.Color{
		internal.CommonColor, internal.RareColor,
		internal.EpicColor, internal.LegendaryColor} {
		text := internal.GetStyleText(c.HexString, c, true)
		fmt.Println(text)
	}
}

func TestMixColors(t *testing.T) {
	red := internal.NewColor(76, 50, 50)
	white := internal.NewColor(240, 240, 240)
	fmt.Printf(
		"БАЗА-0 [%d]: %s + БАЗА-0 [%d]: %s = ", red.A,
		internal.GetStyleText(red.HexString, red, true), white.A,
		internal.GetStyleText(white.HexString, white, true),
	)
	newC := internal.MixColors(red, white)
	fmt.Printf("Чайлд-0 [%d]: %s\n", newC.A, internal.GetStyleText(newC.HexString, newC, true))
	for i := 1; i <= 10; i++ {
		fmt.Printf(
			"Чайлд-%d [%d]: %s + БАЗА-0 [%d]: %s = ", i-1, newC.A,
			internal.GetStyleText(newC.HexString, newC, true), white.A,
			internal.GetStyleText(white.HexString, white, true),
		)
		newC = internal.MixColors(newC, white)
		fmt.Printf("Чайлд-%d [%d]: %s\n", i, newC.A, internal.GetStyleText(newC.HexString, newC, true))
		// time.Sleep(1 * time.Second)
	}
}

func TestPercent(t *testing.T) {
	c1 := internal.NewColor(255, 55, 255)
	if (255-c1.R) <= 55 && (255-c1.G) <= 55 && (255-c1.B) <= 55 {
		fmt.Println(true)
	}
}

func TestMixBase(t *testing.T) {
	rand.Seed(time.Now().Unix())
	colors := []internal.Color{
		internal.BaseBlack,
		internal.BaseWhite,
		internal.BaseRed,
		internal.BaseGreen,
		internal.BaseBlue,
	}
	for i := 0; i < 30; i++ {
		c1 := colors[rand.Intn(len(colors))]
		c2 := colors[rand.Intn(len(colors))]
		newC := internal.MixColors(c1, c2)
		fmt.Printf(
			"[%d]: %s + [%d]: %s = ", c1.A,
			internal.GetStyleText(c1.HexString, c1, true), c2.A,
			internal.GetStyleText(c2.HexString, c2, true),
		)
		fmt.Printf("[%d]: %s\n", newC.A, internal.GetStyleText(newC.HexString, newC, true))
	}
}
