package pool

import "fmt"

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "lane":
		r = fmt.Sprintf("Lane_%v", suffix)
	case "laneset":
		r = fmt.Sprintf("LaneSet_%v", suffix)
	case "id":
		r = fmt.Sprintf("%s", suffix)
	}
	return r
}
