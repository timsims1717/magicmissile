package data

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"image/color"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
)

type Wizard struct {
	Base *ecs.Entity
	Arm  *ecs.Entity
	Obj  *object.Object
	HP   *Health
	Move *Moving
}

type Fighter struct {
	Obj    *object.Object
	HP     *Health
	Move   *Moving
	Attack *Attack
}

type Arm struct {
	Obj     *object.Object
	Spr     *img.Sprite
	Resting float64
	WindUp  float64
	Strike  float64
}

type PC struct {
	Char *Character
	Move *Moving
}

type Character struct {
	Obj    *object.Object
	Spr    *img.Sprite
	Health *Health
}

type Missile struct {
	Object *object.Object
	Sprite *img.Sprite
	Target pixel.Vec
	Speed  float64
	Finish []interface{}
}

type Explosion struct {
	CurrRadius float64
	FullRadius float64
	ExpandRate float64
	Dissipate  float64
	DisRate    float64
	DisRadius  float64
	DisYOffset float64
	StartColor color.RGBA
	FullColor  color.RGBA
	EndColor   color.RGBA
	Timer      *timing.Timer
}

type Health struct {
	Dead bool
	HP   int
	Team Team
}

type Moving struct {
	Selected bool
	Moving   bool
	Speed    float64
	Key      string
	Up       bool
	Wait     bool
}

type Town struct {
	Health *Health
	Object *object.Object
	Sprite *img.Sprite
	Entity *ecs.Entity
}

type Tower struct {
	Health *Health
	Object *object.Object
	Sprite *img.Sprite
	Entity *ecs.Entity
	Origin pixel.Vec
}

type Mob struct {
	Char   *Character
	Attack *Attack
	Speed  float64
	Target *Town
}

type Attack struct {
	Attacking bool
	WindUp    float64
	WindDown  float64
	Recover   float64
	Damage    int
	Range     float64
	Team      Team
	Timer     *timing.Timer
	Target    *ecs.Entity
}

type Team int

const (
	NoTeam = iota
	Player
	Enemy
)
