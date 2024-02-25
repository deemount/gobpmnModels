package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

// NewManualTask ...
func NewManualTask() ManualTaskRepository {
	return &ManualTask{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (manualTask *ManualTask) SetID(typ string, suffix interface{}) {
	manualTask.ID = SetID(typ, suffix)
}

// SetName ...
func (manualTask *ManualTask) SetName(name string) {
	manualTask.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (manualTask *ManualTask) SetDocumentation() {
	manualTask.Documentation = make([]attributes.Documentation, 1)
}

/*** Marker ***/

// SetIncoming ...
func (manualTask *ManualTask) SetIncoming(num int) {
	manualTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (manualTask *ManualTask) SetOutgoing(num int) {
	manualTask.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (manualTask ManualTask) GetID() impl.STR_PTR {
	return &manualTask.ID
}

// GetName ...
func (manualTask ManualTask) GetName() impl.STR_PTR {
	return &manualTask.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (manualTask ManualTask) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &manualTask.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (manualTask ManualTask) GetIncoming(num int) *marker.Incoming {
	return &manualTask.Incoming[num]
}

// GetOutgoing ...
func (manualTask ManualTask) GetOutgoing(num int) *marker.Outgoing {
	return &manualTask.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (manualTask ManualTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", manualTask.ID, manualTask.Name)
}
