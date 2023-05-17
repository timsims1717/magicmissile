package loading

import (
	"encoding/json"
	"github.com/faiface/pixel"
	"github.com/pkg/errors"
	"os"
	"timsims1717/magicmissile/internal/data"
)

func LoadSpells(path string) error {
	errMsg := "load spells"
	data.Missiles = make(map[string][data.MaxSpellTier]*data.Missile)
	content, err := os.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	var missiles []*data.Missile
	err = json.Unmarshal(content, &missiles)
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	data.Missiles = make(map[string][data.MaxSpellTier]*data.Missile)
	for _, missile := range missiles {
		data.SpellKeys = append(data.SpellKeys, missile.Key)
		for i := 1; i <= data.MaxSpellTier; i++ {
			if missile.Tier <= i {
				// create the lowest tier missile
				nMissile := &data.Missile{
					Name:    missile.Name,
					Key:     missile.Key,
					Desc:    missile.Desc,
					Tier:    i,
					Target:  missile.Target,
					SprKey:  missile.SprKey,
					Count:   missile.Count,
					Delay:   missile.Delay,
					Spread:  missile.Spread,
					Speed:   missile.Speed,
					Colors:  missile.Colors,
					Payload: missile.Payload,
				}
				// upgrade if this is not the lowest tier
				if missile.Tier < i {
					j := i - missile.Tier - 1
					if j < len(missile.Tiers) {
						tier := missile.Tiers[j]
						if tier.Desc != "" {
							nMissile.Desc = tier.Desc
						}
						if tier.SprKey != "" {
							nMissile.SprKey = tier.SprKey
						}
						if tier.Target != pixel.ZV {
							nMissile.Target = tier.Target
						}
						if tier.Count > 0 {
							nMissile.Count = tier.Count
						}
						if tier.Delay > 0 {
							nMissile.Delay = tier.Delay
						}
						if tier.Spread > 0 {
							nMissile.Spread = tier.Spread
						}
						if len(tier.Colors) > 0 {
							nMissile.Colors = tier.Colors
						}
						if len(tier.Payload) > 0 {
							nMissile.Payload = tier.Payload
						}
					}
				}
				// set the payload settings
				//for j, f := range nMissile.Payload {
				//	if f.Explosion != nil {
				//
				//	}
				//}
				// add the missile to its set
				if _, ok := data.Missiles[missile.Key]; ok {
					set := data.Missiles[missile.Key]
					set[i-1] = nMissile
					data.Missiles[missile.Key] = set
				} else {
					var set [data.MaxSpellTier]*data.Missile
					set[i-1] = nMissile
					data.Missiles[missile.Key] = set
				}
			}
		}
	}
	return nil
}
