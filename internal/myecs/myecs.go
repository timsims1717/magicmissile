package myecs

import "github.com/bytearena/ecs"

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

	// Tags
	IsObject   = ecs.BuildTag(Object)
	IsTemp     = ecs.BuildTag(Temp, Object)
	HasParent  = ecs.BuildTag(Object, Parent)
	IsDrawable = ecs.BuildTag(Object, Drawable)
	HasAnim    = ecs.BuildTag(Animated)
	HasUpdate  = ecs.BuildTag(Update, Object)

	HasHealth  = ecs.BuildTag(Object, Health, Hitbox)
	PlayerChar = ecs.BuildTag(Object, Movable, Health)
	HasPayload = ecs.BuildTag(Object, Payload)
	CanAttack  = ecs.BuildTag(Object, Attack)
	IsMob      = ecs.BuildTag(Object, Mob)
)

type ClearFlag bool