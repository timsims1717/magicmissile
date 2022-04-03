package data

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
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
	Target  pixel.Vec
	Speed   float64
	Finish  func(pixel.Vec)
}

type Explosion struct {
	Radius     float64
	CurrRadius float64
	Expansion  float64
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
}

type Town struct {
	Health *Health
	Obj    *object.Object
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

type Animations int

const (
	Idle = iota
	Move
	WindUp
	Atk
)