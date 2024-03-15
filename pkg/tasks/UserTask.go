package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewUserTask ...
func NewUserTask() UserTaskRepository {
	return &UserTask{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (utask *UserTask) SetID(typ string, suffix interface{}) {
	utask.ID = SetID(typ, suffix)
}

// SetName ...
func (utask *UserTask) SetName(name string) {
	utask.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (utask *UserTask) SetDocumentation() {
	utask.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*** Marker ***/

// SetIncoming ...
func (utask *UserTask) SetIncoming(num int) {
	utask.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (utask *UserTask) SetOutgoing(num int) {
	utask.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (utask UserTask) GetID() impl.STR_PTR {
	return &utask.ID
}

// GetName ...
func (utask UserTask) GetName() impl.STR_PTR {
	return &utask.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (utask UserTask) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &utask.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (utask UserTask) GetIncoming(num int) *marker.Incoming {
	return &utask.Incoming[num]
}

// GetOutgoing ...
func (utask UserTask) GetOutgoing(num int) *marker.Outgoing {
	return &utask.Outgoing[num]
}

/*
 * @String
 */

// String ...
func (utask UserTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", utask.ID, utask.Name)
}
