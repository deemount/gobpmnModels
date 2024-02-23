package data

import "fmt"

func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "dataobject":
		r = fmt.Sprintf("DataObject_%v", suffix)
		break
	case "dataobjectreference":
		r = fmt.Sprintf("DataObjectReference_%v", suffix)
		break
	case "datastorereference":
		r = fmt.Sprintf("DataStoreReference_%v", suffix)
		break
	case "id":
		r = fmt.Sprintf("%s", suffix)
		break
	}
	return r
}
