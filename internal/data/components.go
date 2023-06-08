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
	Arc     float64        `json:"arc"`
	Angle   float64        `json:"angle"`
	Tier    int            `json:"tier"`
	Target  pixel.Vec      `json:"target"`
	Limit   float64        `json:"limit"`
	Travel  float64        `json:"-"`
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
	Shrink     bool          `json:"shrink"`
	Shrinking  bool          `json:"-"`
	Movement   pixel.Vec     `json:"movement"`
	MoveSpeed  float64       `json:"moveSpeed"`
	CurrMove   pixel.Vec     `json:"-"`
	DisRadius  float64       `json:"-"`
	Color      color.RGBA    `json:"-"`
	Timer      *timing.Timer `json:"-"`
}

func (e *Explosion) Copy() *Explosion {
	return &Explosion{
		FullRadius: e.FullRadius,
		ExpandRate: e.ExpandRate,
		Dissipate:  e.Dissipate,
		DisRate:    e.DisRate,
		Movement:   e.Movement,
		MoveSpeed:  e.MoveSpeed,
		Color:      e.Color,
	}
}

type Payload struct {
	Missile   *Missile   `json:"missile"`
	Explosion *Explosion `json:"explosion"`
	Script    *string    `json:"script"`

	Function func(*Missile, *object.Object) `json:"-"`
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
	Health   *Health
	Object   *object.Object
	Sprite   *img.Sprite
	Entity   *ecs.Entity
	Origin   pixel.Vec
	Slots    []*SpellSlot
	CurrSlot int
	Exp      int
	Level    int
}

type SpellSlot struct {
	Tier  int
	Spell string
	Name  string
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
