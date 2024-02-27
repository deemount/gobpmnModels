package definitions

import (
	"fmt"
)

/*
 * @Functions
 */

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "cancel":
		r = fmt.Sprintf("CancelEventDefinition_%v", suffix)
	case "conditional":
		r = fmt.Sprintf("ConditionalEventDefinition_%v", suffix)
	case "error":
		r = fmt.Sprintf("ErrorEventDefinition_%v", suffix)
	case "escalation":
		r = fmt.Sprintf("EscalationEventDefinition_%v", suffix)
	case "link":
		r = fmt.Sprintf("LinkEventDefinition_%v", suffix)
	case "message":
		r = fmt.Sprintf("MessageEventDefinition_%v", suffix)
	case "signal":
		r = fmt.Sprintf("SignalEventDefinition_%v", suffix)
	case "terminate":
		r = fmt.Sprintf("TerminateEventDefinition_%v", suffix)
	case "timer":
		r = fmt.Sprintf("TimerEventDefinition_%v", suffix)
	case "id":
		r = fmt.Sprintf("%s", suffix)
	}
	return r
}
