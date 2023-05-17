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
	Name    string         `json:"name"`
	Key     string         `json:"key"`
	Desc    string         `json:"desc"`
	SprKey  string         `json:"sprite"`
	Count   int            `json:"count"`
	Delay   float64        `json:"delay"`
	Spread  float64        `json:"spread"`
	Tier    int            `json:"tier"`
	Target  pixel.Vec      `json:"target"`
	Speed   float64        `json:"speed"`
	Colors  []string       `json:"colors"`
	Object  *object.Object `json:"-"`
	Sprite  *img.Sprite    `json:"-"`
	Payload []Payload      `json:"payloads"`
	Tiers   []*Missile     `json:"tiers"`
}

type Explosion struct {
	CurrRadius float64       `json:"-"`
	FullRadius float64       `json:"radius"`
	ExpandRate float64       `json:"expandRate"`
	Dissipate  float64       `json:"dissipateAfter"`
	DisRate    float64       `json:"dissipateRate"`
	DisRadius  float64       `json:"-"`
	StartColor color.RGBA    `json:"-"`
	FullColor  color.RGBA    `json:"-"`
	EndColor   color.RGBA    `json:"-"`
	Timer      *timing.Timer `json:"-"`
}

type Payload struct {
	Missile   *Missile   `json:"missile"`
	Explosion *Explosion `json:"explosion"`
	Spell     *string    `json:"spell"`
	Function  func()     `json:"-"`
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
