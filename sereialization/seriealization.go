package sereialization

import (
	"encoding/json"
)

func Seriealize(val interface{}) string {
	s, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	return string(s)
}
