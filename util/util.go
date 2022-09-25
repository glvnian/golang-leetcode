package util

import "encoding/json"

func InterfaceToString(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}
