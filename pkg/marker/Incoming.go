package marker

import (
	"fmt"

	impl "github.com/deemount/gobpmnTypes"
)

// NewIncoming ...
func NewIncoming() IncomingRepository {
	return &Incoming{}
}

/*
 * @Setters
 */

/* Content */

/** BPMN **/

// SetFlow ...
func (incoming *Incoming) SetFlow(suffix interface{}) {
	incoming.Flow = fmt.Sprintf("Flow_%s", suffix)
}

/*
 * @Getters
 */

/* Content */

/** BPMN **/

// GetFlow ...
func (incoming Incoming) GetFlow() impl.STR_PTR {
	return &incoming.Flow
}
