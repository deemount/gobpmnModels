package attributes

import "fmt"

func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "property":
		r = fmt.Sprintf("Property_%v", suffix)
		break
	case "attribute":
		r = fmt.Sprintf("Attribute_%v", suffix)
		break
	case "id":
		r = fmt.Sprintf("%s", suffix)
		break
	}
	return r
}
