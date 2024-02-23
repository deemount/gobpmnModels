package tasks

import "fmt"

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "activity":
		r = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		r = fmt.Sprintf("%s", suffix)
		break
	}
	return r
}
