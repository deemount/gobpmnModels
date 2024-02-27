package flow

import "fmt"

/*
 * @Functions
 */

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "association":
		r = fmt.Sprintf("Association_%v", suffix)
	case "datainputassociation":
		r = fmt.Sprintf("DataInputAssociation_%v", suffix)
	case "flow":
		r = fmt.Sprintf("Flow_%v", suffix)
	case "id":
		r = fmt.Sprintf("%s", suffix)
	}
	return r
}

// SetSourceTargetRef ...
func SetSourceTargetRef(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "activity":
		r = fmt.Sprintf("Activity_%s", suffix)
	case "event":
		r = fmt.Sprintf("Event_%s", suffix)
	case "id":
		r = fmt.Sprintf("%s", suffix)
	case "participant":
		r = fmt.Sprintf("Participant_%s", suffix)
	case "startevent":
		r = fmt.Sprintf("StartEvent_%s", suffix)
	}
	return r
}
