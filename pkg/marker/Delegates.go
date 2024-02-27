package marker

import "fmt"

/*
 * @Functions
 */

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	// TODO: Checkout CategoryValue case for empty string
	case "":
		r = fmt.Sprintf("_%s", suffix)
	case "category":
		r = fmt.Sprintf("Category_%s", suffix)
	case "group":
		r = fmt.Sprintf("Group_%v", suffix)
	case "id":
		r = fmt.Sprintf("_%s", suffix)
	}
	return r
}
