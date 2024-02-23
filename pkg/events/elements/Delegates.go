package elements

import (
	"fmt"
)

var (
	structMessage = "message"
)

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "event":
		r = fmt.Sprintf("Event_%v", suffix)
		break
	case "message":
		r = fmt.Sprintf("Message_%v", suffix)
		break
	case "signal":
		r = fmt.Sprintf("Signal_%v", suffix)
		break
	case "startevent":
		r = fmt.Sprintf("StartEvent_%v", suffix)
		break
	case "id":
		r = fmt.Sprintf("%s", suffix)
		break
	}
	return r
}
