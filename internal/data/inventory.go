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

	TowerScrollWidth  = 27.
	TowerScrollHeight = 35.
	TileSize          = 16.
	TowerScrollYStart = (BaseHeight + TowerScrollHeight*TileSize) * 0.5
	LeftScrollStart   = pixel.V(-BaseWidth*0.33+25., TowerScrollYStart)
	MidScrollStart    = pixel.V(0., TowerScrollYStart)
	RightScrollStart  = pixel.V(BaseWidth*0.33-25., TowerScrollYStart)
	TowerScrollY      = 60.
	LeftScrollPos     = pixel.V(-BaseWidth*0.33+25., TowerScrollY)
	MidScrollPos      = pixel.V(0., TowerScrollY)
	RightScrollPos    = pixel.V(BaseWidth*0.33-25., TowerScrollY)
	ScrollScale       = (TowerScrollHeight + 6) * 0.25

	TowerViewPos  = pixel.V(10., -15.)
	TowerTitlePos = pixel.V(108., -15.)
	SubtitlePos   = pixel.V(121., -18.)

	SlotsHeadX = 5.
	LevelHeadX = SlotsHeadX + TileSize*3.
	NameHeadX  = LevelHeadX + TileSize*3.

	SlotListViewX = 155.
	SlotListViewY = -136.
)

type TowerScroll struct {
	Tower    *Tower
	Scroll   *Scroll
	Entities []*ecs.Entity

	TowerViewObj *object.Object
	TitleText    *typeface.Text
	SubclassText *typeface.Text
	LevelText    *typeface.Text
	SlotHead     *typeface.Text
	LvlHead      *typeface.Text
	NameHead     *typeface.Text
	TableHeadY   float64

	ListViewObj *object.Object
	ListView    *viewport.ViewPort
}

func (t *TowerScroll) AddEntity(e *ecs.Entity) {
	t.Entities = append(t.Entities, e)
}

type Scroll struct {
	Object   *object.Object
	Entity   *ecs.Entity
	FullDim  pixel.Vec
	CurrDim  pixel.Vec
	TLPos    pixel.Vec
	Opened   bool
	Freeze   bool
	Entities []*ecs.Entity

	TopLeft  *object.Object
	TopMid   *object.Object
	TopRight *object.Object
	MidLeft  *object.Object
	MidMid   *object.Object
	MidRight *object.Object
	BotLeft  *object.Object
	BotMid   *object.Object
	BotRight *object.Object
}

func (t *Scroll) AddEntity(e *ecs.Entity) {
	t.Entities = append(t.Entities, e)
}
