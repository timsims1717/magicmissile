package systems

import (
	"github.com/faiface/pixel"
	"math/rand"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
)

func CreateTowers() {
	if data.CurrBackground != nil {
		bg := data.CurrBackground.Backgrounds[data.TownLayer]
		p := bg.Perlin
		div := float64(data.BaseWidth / 10)
		data.Towns = []*data.Town{}
		for i := 0; i < 9; i++ {
			if i == 0 || i == 4 || i == 8 {
				spr := img.NewSprite("wizard_tower", data.ObjectKey)
				obj := object.New()
				x := div * float64(-4+i)
				obj.Pos.X = x
				obj.Pos.Y = p.Noise1D(x/bg.Layer.WaveLength)*bg.Layer.Scale + data.ForeOffset + bg.Layer.VOffset(x)
				obj.Layer = data.TownLayer
				obj.Flip = rand.Intn(2) == 0
				obj.Offset.Y += img.Batchers[data.ObjectKey].GetSprite(spr.Key).Frame().H()*0.5 - 10.
				obj.Pos.Y += 6.
				var slots []*data.SpellSlot
				for j := 0; j < data.SpellSlotNum; j++ {
					spellKey := data.SpellKeys[rand.Intn(len(data.SpellKeys))]
					baseTier := data.Missiles[spellKey][0]
					slotTier := baseTier.Tier
					for slotTier < data.MaxSpellTier {
						if rand.Intn(10) == 0 {
							slotTier++
						} else {
							break
						}
					}
					slots = append(slots, &data.SpellSlot{
						Tier:  slotTier,
						Spell: spellKey,
						Name:  baseTier.Name,
					})
				}
				e := myecs.Manager.NewEntity()
				e.AddComponent(myecs.Object, obj).
					AddComponent(myecs.Drawable, spr)
				data.Towers = append(data.Towers, &data.Tower{
					Health: nil,
					Object: obj,
					Sprite: spr,
					Entity: e,
					Origin: pixel.V(0, 96.),
					Slots:  slots,
				})
			}
		}
	} else {
		panic("towers can't be created without a background")
	}
}

func CreateTowersNoBG() {
	div := float64(data.BaseWidth / 10)
	data.Towns = []*data.Town{}
	for i := 0; i < 9; i++ {
		if i == 0 || i == 4 || i == 8 {
			spr := img.NewSprite("wizard_tower", data.ObjectKey)
			obj := object.New()
			x := div * float64(-4+i)
			obj.Pos.X = x
			obj.Pos.Y = data.ForeOffset
			obj.Layer = data.TownLayer
			obj.Flip = rand.Intn(2) == 0
			obj.Offset.Y += img.Batchers[data.ObjectKey].GetSprite(spr.Key).Frame().H()*0.5 - 10.
			obj.Pos.Y += 6.
			var slots []*data.SpellSlot
			for j := 0; j < data.SpellSlotNum; j++ {
				spellKey := data.SpellKeys[rand.Intn(len(data.SpellKeys))]
				baseTier := data.Missiles[spellKey][0]
				slotTier := baseTier.Tier
				for slotTier < data.MaxSpellTier {
					if rand.Intn(10) == 0 {
						slotTier++
					} else {
						break
					}
				}
				slots = append(slots, &data.SpellSlot{
					Tier:  slotTier,
					Spell: spellKey,
					Name:  baseTier.Name,
				})
			}
			e := myecs.Manager.NewEntity()
			e.AddComponent(myecs.Object, obj).
				AddComponent(myecs.Drawable, spr)
			data.Towers = append(data.Towers, &data.Tower{
				Health: nil,
				Object: obj,
				Sprite: spr,
				Entity: e,
				Origin: pixel.V(0, 96.),
				Slots:  slots,
			})
		}
	}
}

func DisposeTowers() {
	for _, tower := range data.Towers {
		myecs.Manager.DisposeEntity(tower.Entity)
	}
	data.Towers = []*data.Tower{}
}
