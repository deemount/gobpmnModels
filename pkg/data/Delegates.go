package data

import "fmt"

/*
 * @Functions
 */

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "dataobject":
		r = fmt.Sprintf("DataObject_%v", suffix)
	case "dataobjectreference":
		r = fmt.Sprintf("DataObjectReference_%v", suffix)
	case "datastorereference":
		r = fmt.Sprintf("DataStoreReference_%v", suffix)
	case "id":
		r = fmt.Sprintf("%s", suffix)
	}
	return r
}
