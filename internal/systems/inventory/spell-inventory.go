package inventory

import (
	"github.com/faiface/pixel"
	"math/rand"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/typeface"
)

func CreateSpellInventory() {
	DisposeSpellInventory()
	data.SpellInventory.Spells = make(map[string]int)
	for key, spell := range data.Missiles {
		if key == "arcaneshot" {
			data.SpellInventory.Spells[key] = -1
		} else {
			cnt := rand.Intn(5) - spell[0].Tier
			if cnt < 0 {
				cnt = 0
			}
			data.SpellInventory.Spells[key] = 0
		}
	}
	data.SpellInventory.Scroll = systems.CreateScroll(data.SpellInventoryPos, 101, pixel.V(data.SpellInventoryWidth*data.TileSize, data.ScrollHeight*data.TileSize), pixel.V(data.SpellInventoryWidth*data.TileSize, data.ScrollHeight*data.TileSize))
	data.SpellInventory.TitleText = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.15, 300., 0.)
	data.SpellInventory.TitleText.Obj.Layer = 103
	data.SpellInventory.TitleText.SetOffset(data.SpellInventory.Scroll.TLPos.Add(data.InventoryTitlePos))
	data.SpellInventory.TitleText.SetColor(data.ScrollText)
	data.SpellInventory.TitleText.SetText("Spell Inventory")
	data.SpellInventory.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.SpellInventory.TitleText.Obj).
		AddComponent(myecs.Parent, data.SpellInventory.Scroll.Object).
		AddComponent(myecs.Drawable, data.SpellInventory.TitleText).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))
}

func DisposeSpellInventory() {
	for _, e := range data.SpellInventory.Entities {
		myecs.Manager.DisposeEntity(e)
	}
	if data.SpellInventory.Scroll != nil {
		systems.DisposeScroll(data.SpellInventory.Scroll)
	}
	data.SpellInventory.TitleText = nil
	data.SpellInventory.ListViewObj = nil
	data.SpellInventory.ListView = nil
}
