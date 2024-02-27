package elements

import (
	"fmt"
)

/*
 * @Functions
 */

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "event":
		r = fmt.Sprintf("Event_%v", suffix)
	case "message":
		r = fmt.Sprintf("Message_%v", suffix)
	case "signal":
		r = fmt.Sprintf("Signal_%v", suffix)
	case "startevent":
		r = fmt.Sprintf("StartEvent_%v", suffix)
	case "id":
		r = fmt.Sprintf("%s", suffix)
	}
	return r
}
