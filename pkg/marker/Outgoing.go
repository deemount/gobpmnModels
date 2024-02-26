package marker

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/impl"
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
