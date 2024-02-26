package attributes

import "fmt"

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "property":
		r = fmt.Sprintf("Property_%v", suffix)
	case "attribute":
		r = fmt.Sprintf("Attribute_%v", suffix)
	case "id":
		r = fmt.Sprintf("%s", suffix)
	}
	return r
}
