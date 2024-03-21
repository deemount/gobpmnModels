package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewTask ...
func NewTask() TaskRepository {
	return &Task{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (task *Task) SetID(typ string, suffix interface{}) {
	task.ID = SetID(typ, suffix)
}

// SetName ...
func (task *Task) SetName(name string) {
	task.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (task *Task) SetDocumentation() {
	task.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

// SetProperty ...
func (task *Task) SetProperty() {
	task.Property = make(attributes.PROPERTY_SLC, 1)
}

/*** Flow ***/

// SetDataInputAssociation ...
func (task *Task) SetDataInputAssociation(num int) {
	task.DataInputAssociation = make(flow.DATA_INPUT_ASSOCIATION_SLC, num)
}

/*** Marker ***/

// SetIncoming ...
func (task *Task) SetIncoming(num int) {
	task.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (task *Task) SetOutgoing(num int) {
	task.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (task Task) GetID() impl.STR_PTR {
	return &task.ID
}

// GetName ...
func (task Task) GetName() impl.STR_PTR {
	return &task.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (task Task) GetDocumentation() *attributes.Documentation {
	return &task.Documentation[0]
}

// GetProperty ...
func (task Task) GetProperty() *attributes.Property {
	return &task.Property[0]
}

/*** Flow ***/

// GetDataInputAssociation ...
func (task Task) GetDataInputAssociation(num int) flow.DATA_INPUT_ASSOCIATION_PTR {
	return &task.DataInputAssociation[num]
}

/*** Marker ***/

// GetIncoming ...
func (task Task) GetIncoming(num int) *marker.Incoming {
	return &task.Incoming[num]
}

// GetOutgoing ...
func (task Task) GetOutgoing(num int) *marker.Outgoing {
	return &task.Outgoing[num]
}

/*
 * @String
 */

// String ...
func (task Task) String() string {
	return fmt.Sprintf("id=%v, name=%v", task.ID, task.Name)
}
