package collaboration

import "fmt"

func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "collaboration":
		r = fmt.Sprintf("Collaboration_%v", suffix)
		break
	case "participant":
		r = fmt.Sprintf("Participant_%v", suffix)
		break
	case "id":
		r = fmt.Sprintf("%s", suffix)
		break
	}
	return r
}
