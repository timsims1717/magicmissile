package data

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/typeface"
	"timsims1717/magicmissile/pkg/viewport"
)

var (
	InventoryView    *viewport.ViewPort
	LeftTowerScroll  *TowerScroll
	MidTowerScroll   *TowerScroll
	RightTowerScroll *TowerScroll

	LeftScrollPos  = pixel.V(-TileSize*(TowerScrollWidth)*1.5, BaseHeight*0.5-TileSize*8.)
	MidScrollPos   = pixel.V(-TileSize*(TowerScrollWidth-4)*0.5, BaseHeight*0.5-TileSize*8.)
	RightScrollPos = pixel.V(TileSize*(TowerScrollWidth+8)*0.5, BaseHeight*0.5-TileSize*8.)
	ScrollScale    = (TowerScrollHeight + 6) * 0.25

	TowerViewPos  = pixel.V(20., -55.)
	TowerTitlePos = pixel.V(112., -40.)
	SubtitlePosX  = 124.

	SlotsHeadX = 5.
	LevelHeadX = SlotsHeadX + TileSize*3.
	NameHeadX  = LevelHeadX + TileSize*3.

	SlotListViewX = 155.
	SlotListViewY = -136.
)

type TowerScroll struct {
	Tower    *Tower
	Object   *object.Object
	Entities []*ecs.Entity

	TowerViewObj *object.Object
	TitleText    *typeface.Text
	SubclassText *typeface.Text
	LevelText    *typeface.Text
	TableHeadY   float64

	ListViewObj *object.Object
	ListView    *viewport.ViewPort
	TopOfListY  float64
	BotOfListY  float64
}

func (t *TowerScroll) AddEntity(e *ecs.Entity) {
	t.Entities = append(t.Entities, e)
}
