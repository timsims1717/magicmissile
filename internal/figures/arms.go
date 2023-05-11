package figures

import (
	"image/color"
	"math"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
)

func ArmObj() *object.Object {
	armObj := object.New()
	armObj.Pos.X += 12.
	armObj.Pos.Y += 12.
	return armObj
}

func ArmSpr(key string, col color.RGBA) *img.Sprite {
	return &img.Sprite{
		Key:   key,
		Color: col,
		Batch: "figures",
	}
}

func WandArm(col color.RGBA) *data.Arm {
	return &data.Arm{
		Obj: ArmObj(),
		Spr: ArmSpr("wand", col),
	}
}

func AxeArm(col color.RGBA) *data.Arm {
	return &data.Arm{
		Obj:     ArmObj(),
		Spr:     ArmSpr("axe", col),
		Resting: -0.2,
		WindUp:  math.Pi * 0.6,
		Strike:  -1.0,
	}
}

func ZombieArm(col color.RGBA) *data.Arm {
	return &data.Arm{
		Obj:     ArmObj(),
		Spr:     ArmSpr("zomb_arm", col),
		Resting: math.Pi * -0.05,
		WindUp:  math.Pi * 0.2,
		Strike:  math.Pi * -0.15,
	}
}
