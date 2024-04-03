package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewManualTask ...
func NewManualTask() ManualTaskRepository {
	return &ManualTask{}
}

/*
 * @Setters
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
	manualTask.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

// SetExtensionElements ...
func (manualTask *ManualTask) SetExtensionElements() {
	manualTask.ExtensionElements = make(attributes.EXTENSION_ELEMENTS_SLC, 1)
}

/*** Marker ***/

// SetIncoming ...
func (manualTask *ManualTask) SetIncoming(num int) {
	manualTask.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (manualTask *ManualTask) SetOutgoing(num int) {
	manualTask.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*
 * @Getters
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
func (manualTask ManualTask) GetDocumentation() *attributes.Documentation {
	return &manualTask.Documentation[0]
}

// GetExtensionElements ...
func (manualTask ManualTask) GetExtensionElements() *attributes.ExtensionElements {
	return &manualTask.ExtensionElements[0]
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
 * @String
 */

// String ...
func (manualTask ManualTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", manualTask.ID, manualTask.Name)
}
