package gateways

import "fmt"

/*
 * @Functions
 */

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "gateway":
		r = fmt.Sprintf("Gateway_%v", suffix)
	case "id":
		r = fmt.Sprintf("%s", suffix)
	}
	return r
}
