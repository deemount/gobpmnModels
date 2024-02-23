package marker

import "fmt"

func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	// TODO: Checkout CategoryValue case for empty string
	case "":
		r = fmt.Sprintf("_%s", suffix)
		break
	case "category":
		r = fmt.Sprintf("Category_%s", suffix)
		break
	case "group":
		r = fmt.Sprintf("Group_%v", suffix)
		break
	case "id":
		r = fmt.Sprintf("_%s", suffix)
		break
	}
	return r
}
