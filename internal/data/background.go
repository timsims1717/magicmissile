package data

import (
	"bytes"
	"encoding/json"
	"math/rand"
)

type VFnCode int

const (
	Peak = iota
	Valley
	None
)

var toString = map[VFnCode]string{
	Peak:   "Peak",
	Valley: "Valley",
	None:   "None",
}

var toID = map[string]VFnCode{
	"Peak":   Peak,
	"Valley": Valley,
	"None":   None,
}

func (c *VFnCode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[*c])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (c *VFnCode) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*c = toID[j]
	return nil
}

func NoShape(_ float64) float64 {
	return 0.
}

func RandPeak() func(float64) float64 {
	a := rand.Float64()*0.000075 + 0.00025
	b := rand.Float64()*0.05 + 0.025
	c := rand.Float64()*50. - 15.
	return func(x float64) float64 {
		y := -(a * x * x) + b*x + c
		return y
	}
}

func RandValley() func(float64) float64 {
	a := rand.Float64()*0.000075 + 0.00025
	b := rand.Float64()*0.05 + 0.025
	c := rand.Float64()*50. - 15.
	return func(x float64) float64 {
		y := a*x*x - b*x - c
		return y
	}
}

func RandPeaks(p int) func(float64) float64 {
	a := rand.Float64() * 0.005
	b := rand.Float64() * 0.02
	c := rand.Float64() * 50.
	return func(x float64) float64 {
		y := -(a * x * x) + b*x + c
		return y
	}
}
