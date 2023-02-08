package internal

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dariubs/percent"
	paint "github.com/gookit/color"
)

const (
	domi int = 70
	resi int = 30
)

var (
	BaseWhite Color = NewColor(245, 245, 245)
	BaseBlack Color = NewColor(10, 10, 10)
	BaseRed   Color = NewColor(245, 10, 10)
	BaseGreen Color = NewColor(10, 245, 10)
	BaseBlue  Color = NewColor(10, 10, 245)
)

type Color struct {
	R, G, B, A uint8
	HexString  string
	Honorable  bool
}

func (cBase *Color) Mutation(cMut Color) {

}

func NewColor(r, g, b uint8) Color {
	a := GetRandA()
	honor := false
	if a > 200 {
		honor = true
	}
	return Color{
		R: r, G: g, B: b, A: a,
		HexString: HexFromRgbColor(r, g, b),
		Honorable: honor,
	}
}

func MixColors(c1, c2 Color) Color {
	var r, g, b uint8
	if (255-c1.R) <= 55 && (255-c1.G) <= 55 && (255-c1.B) <= 55 {
		c1.A = uint8(0)
	} else if (255-c2.R) <= 55 && (255-c2.G) <= 55 && (255-c2.B) <= 55 {
		c2.A = uint8(0)
	}
	if c1.A >= c2.A {
		r = toUInt8((percent.Percent(domi, int(c1.R))) + (percent.Percent(resi, int(c2.R))))
		g = toUInt8((percent.Percent(domi, int(c1.G))) + (percent.Percent(resi, int(c2.G))))
		b = toUInt8((percent.Percent(domi, int(c1.B))) + (percent.Percent(resi, int(c2.B))))
	} else {
		r = toUInt8((percent.Percent(resi, int(c1.R))) + (percent.Percent(domi, int(c2.R))))
		g = toUInt8((percent.Percent(resi, int(c1.G))) + (percent.Percent(domi, int(c2.G))))
		b = toUInt8((percent.Percent(resi, int(c1.B))) + (percent.Percent(domi, int(c2.B))))
	}
	return NewColor(r, g, b)
}

func toUInt8(f float64) uint8 {
	if int(f) > 245 {
		return uint8(245)
	}
	return uint8(f)
}

func HexFromRgbColor(r, g, b uint8) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

func ParseHexColor(s string) (c Color, err error) {
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

func GetRandA() uint8 {
	rand.Seed(time.Now().UnixNano())
	return uint8(rand.Intn(256))
}

func GetStyleText(text string, color Color, bg ...bool) string {
	style := paint.RGBFromHEX(color.HexString, bg...)
	return style.Sprint(text)
}
