package collaboration

import "fmt"

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "collaboration":
		r = fmt.Sprintf("Collaboration_%v", suffix)
	case "participant":
		r = fmt.Sprintf("Participant_%v", suffix)
	case "id":
		r = fmt.Sprintf("%s", suffix)
	}
	return r
}
