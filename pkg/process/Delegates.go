package process

import "fmt"

/*
 * @Functions
 */

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "process":
		r = fmt.Sprintf("Process_%v", suffix)
	case "id":
		r = fmt.Sprintf("%s", suffix)
	}
	return r
}
