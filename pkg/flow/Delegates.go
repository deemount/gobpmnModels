package flow

import "fmt"

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "flow":
		r = fmt.Sprintf("Flow_%v", suffix)
		break
	case "association":
		r = fmt.Sprintf("Association_%v", suffix)
		break
	case "datainputassociation":
		r = fmt.Sprintf("DataInputAssociation_%v", suffix)
		break
	case "id":
		r = fmt.Sprintf("%s", suffix)
		break
	}
	return r
}

// SetSourceTargetRef ...
func SetSourceTargetRef(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "activity":
		r = fmt.Sprintf("Activity_%s", suffix)
		break
	case "event":
		r = fmt.Sprintf("Event_%s", suffix)
		break
	case "id":
		r = fmt.Sprintf("%s", suffix)
		break
	case "participant":
		r = fmt.Sprintf("Participant_%s", suffix)
		break
	}
	return r
}
