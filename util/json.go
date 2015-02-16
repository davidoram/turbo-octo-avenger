package util

import (
	"encoding/json"
	//  "log"
)

func MustMarshalJSON(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
