package process

import "fmt"

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "process":
		r = fmt.Sprintf("Process_%v", suffix)
		break
	case "id":
		r = fmt.Sprintf("%s", suffix)
		break
	}
	return r
}
