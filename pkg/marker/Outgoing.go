package marker

import (
	"fmt"

	impl "github.com/deemount/gobpmnTypes"
)

// NewOutgoing ...
func NewOutgoing() OutgoingRepository {
	return &Outgoing{}
}

/*
 * @Setters
 */

/* Content */

/** BPMN **/

// SetFlow ...
func (outgoing *Outgoing) SetFlow(suffix interface{}) {
	outgoing.Flow = fmt.Sprintf("Flow_%s", suffix)
}

/*
 * @Getters
 */

/* Content */

/** BPMN **/

// GetFlow ...
func (outgoing Outgoing) GetFlow() impl.STR_PTR {
	return &outgoing.Flow
}
