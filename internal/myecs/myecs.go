package myecs

import (
	"github.com/bytearena/ecs"
	"timsims1717/magicmissile/pkg/object"
)

var (
	FullCount   = 0
	IDCount     = 0
	LoadedCount = 0
)

var (
	Manager = ecs.NewManager()

	// Components
	Object = Manager.NewComponent()
	Parent = Manager.NewComponent()
	Temp   = Manager.NewComponent()
	Update = Manager.NewComponent()

	Drawable = Manager.NewComponent()
	Animated = Manager.NewComponent()

	Payload = Manager.NewComponent()
	Health  = Manager.NewComponent()
	Movable = Manager.NewComponent()
	Attack  = Manager.NewComponent()
	Mob     = Manager.NewComponent()
	Hitbox  = Manager.NewComponent()

	Explosion = Manager.NewComponent()

	// Tags
	IsObject   = ecs.BuildTag(Object)
	IsTemp     = ecs.BuildTag(Temp, Object)
	HasParent  = ecs.BuildTag(Object, Parent)
	IsDrawable = ecs.BuildTag(Object, Drawable)
	HasAnim    = ecs.BuildTag(Animated)
	HasUpdate  = ecs.BuildTag(Update)

	HasHealth  = ecs.BuildTag(Object, Health, Hitbox)
	PlayerChar = ecs.BuildTag(Object, Movable, Health)
	HasPayload = ecs.BuildTag(Object, Payload)
	CanAttack  = ecs.BuildTag(Object, Attack)
	IsMob      = ecs.BuildTag(Object, Mob)

	IsExplosion = ecs.BuildTag(Object, Explosion)
)

type ClearFlag bool

func UpdateManager() {
	LoadedCount = 0
	IDCount = 0
	FullCount = 0
	for _, result := range Manager.Query(IsObject) {
		if t, ok := result.Components[Object].(*object.Object); ok {
			FullCount++
			if t.ID != "" {
				IDCount++
				if t.Load {
					LoadedCount++
				}
			}
		}
	}
}
