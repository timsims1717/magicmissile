package data

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/typeface"
	"timsims1717/magicmissile/pkg/viewport"
)

var (
	InventoryView    *viewport.ViewPort
	LeftTowerScroll  *TowerScroll
	MidTowerScroll   *TowerScroll
	RightTowerScroll *TowerScroll

	// InventoryState is the current inventory item that is open
	// -1 is for none
	// 3 is for all three tower scrolls open
	// 0-2 is for a certain tower scroll being edited
	InventoryState = -1
	InventoryTrans = false

	TowerScrollWidth  = 27.
	ScrollHeight      = 35.
	TileSize          = 16.
	TowerScrollYStart = (BaseHeight + ScrollHeight*TileSize) * 0.5
	LeftScrollStart   = pixel.V(-BaseWidth*0.33+25., TowerScrollYStart)
	MidScrollStart    = pixel.V(0., TowerScrollYStart)
	RightScrollStart  = pixel.V(BaseWidth*0.33-25., TowerScrollYStart)
	TowerScrollY      = 60.
	LeftScrollX       = -BaseWidth*0.33 + 25.
	//MidScrollX        = 0.
	//RightScrollX      = BaseWidth*0.33 - 25.

	SpellInventoryPos   = pixel.V(BaseWidth*0.167, TowerScrollY)
	SpellInventoryWidth = TowerScrollWidth * 2.
	InventoryTitlePos   = pixel.V(5., -5.)

	TowerViewPos  = pixel.V(10., -15.)
	TowerTitlePos = pixel.V(21., -15.)
	SubtitlePos   = pixel.V(31., -18.)
	EditButtonPos = pixel.V(21., -21.)

	SlotsHeadX = 5.
	LevelHeadX = SlotsHeadX + TileSize*3.
	NameHeadX  = LevelHeadX + TileSize*3.

	SlotListViewX = 155.
	SlotListViewY = -136.

	SpellInventory  = &spellInventory{}
	MovingSpellSlot *InvSpellSlot
	MSSEntities     []*ecs.Entity
)

type spellInventory struct {
	Spells   map[string]int
	Scroll   *Scroll
	Entities []*ecs.Entity

	TitleText *typeface.Text

	ListViewObj *object.Object
	ListView    *viewport.ViewPort
}

func (t *spellInventory) AddEntity(e *ecs.Entity) {
	t.Entities = append(t.Entities, e)
}

type TowerScroll struct {
	Tower    *Tower
	Scroll   *Scroll
	Entities []*ecs.Entity
	PosX     float64

	TowerViewObj *object.Object
	TitleText    *typeface.Text
	SubclassText *typeface.Text
	LevelText    *typeface.Text
	EditButton   *object.Object
	SaveButton   *object.Object
	EditText     *typeface.Text
	SlotHead     *typeface.Text
	LvlHead      *typeface.Text
	NameHead     *typeface.Text
	TableHeadY   float64
	InvSlots     []*InvSpellSlot

	ListViewObj *object.Object
	ListView    *viewport.ViewPort
}

func (t *TowerScroll) AddEntity(e *ecs.Entity) {
	t.Entities = append(t.Entities, e)
}

type Scroll struct {
	Object   *object.Object
	Inters   []*object.Interpolation
	Entity   *ecs.Entity
	FullDim  pixel.Vec
	CurrDim  pixel.Vec
	TLPos    pixel.Vec
	Opened   bool
	Closed   bool
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

type InvSpellSlot struct {
	SlotObj *object.Object
	SlotSpr *img.Sprite
	SlotTxt *typeface.Text

	TierObj *object.Object
	TierSpr *img.Sprite
	TierTxt *typeface.Text

	NameMObj *object.Object
	NameMSpr *img.Sprite
	NameLObj *object.Object
	NameLSpr *img.Sprite
	NameRObj *object.Object
	NameRSpr *img.Sprite
	NameTxt  *typeface.Text

	Slot     *SpellSlot
	PrevSlot *SpellSlot
}
