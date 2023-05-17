package loading

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
	"timsims1717/magicmissile/internal/data"
)

func LoadRealms(path string) error {
	errMsg := "load realms"
	content, err := os.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	var realms []*data.Realm
	err = json.Unmarshal(content, &realms)
	if err != nil {
		return errors.Wrap(err, errMsg)
	}
	data.AllRealms = make(map[string]*data.Realm)
	for _, realm := range realms {
		for _, background := range realm.Backgrounds {
			for _, layer := range background.Layers {
				switch layer.VFnCode {
				case data.Peak:
					layer.VOffset = data.RandPeak()
				case data.Valley:
					layer.VOffset = data.RandValley()
				case data.None:
					layer.VOffset = data.NoShape
				}
			}
		}
		data.AllRealms[realm.Code] = realm
	}
	return nil
}
