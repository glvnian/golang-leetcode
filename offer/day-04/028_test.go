package day_04

import "encoding/json"

func StruckToString(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}
