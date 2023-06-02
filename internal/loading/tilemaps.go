package loading

import (
	"github.com/faiface/pixel"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/pkg/img"
)

func LoadTileMaps() {
	CreateScrollMap()
}

func CreateScrollMap() {
	// Scroll Top
	for x := 0; x < int(data.TowerScrollWidth); x++ {
		str := "scroll_t"
		if x == 0 {
			str = "scroll_tl"
		} else if x == int(data.TowerScrollWidth)-1 {
			str = "scroll_tr"
		}
		offset := pixel.V(float64(x)*data.TileSize-data.TowerScrollWidth/2, 0.)
		spr := img.NewOffsetSprite(str, data.UIKey, offset)
		data.ScrollTop = append(data.ScrollTop, spr)
	}
	// Scroll Mid
	for x := 0; x < int(data.TowerScrollWidth); x++ {
		str := "scroll_m"
		if x == 0 {
			str = "scroll_ml"
		} else if x == int(data.TowerScrollWidth)-1 {
			str = "scroll_mr"
		}
		offset := pixel.V(float64(x)*data.TileSize-data.TowerScrollWidth/2, 0.)
		spr := img.NewOffsetSprite(str, data.UIKey, offset)
		data.ScrollMid = append(data.ScrollMid, spr)
	}
	// Scroll Bottom
	for x := 0; x < int(data.TowerScrollWidth); x++ {
		str := "scroll_b"
		if x == 0 {
			str = "scroll_bl"
		} else if x == int(data.TowerScrollWidth)-1 {
			str = "scroll_br"
		}
		offset := pixel.V(float64(x)*data.TileSize-data.TowerScrollWidth/2, 0.)
		spr := img.NewOffsetSprite(str, data.UIKey, offset)
		data.ScrollBot = append(data.ScrollBot, spr)
	}
}
