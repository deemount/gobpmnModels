package gateways

import "fmt"

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "gateway":
		r = fmt.Sprintf("Gateway_%v", suffix)
		break
	case "id":
		r = fmt.Sprintf("%s", suffix)
		break
	}
	return r
}
