package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

// NewTask ...
func NewTask() TaskRepository {
	return &Task{}
}

/*
 * Default Setters
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

/*** Make Elements ***/

/** BPMN **/

// SetProperty ...
func (task *Task) SetProperty() {
	task.Property = make([]attributes.Property, 1)
}

// SetDocumentation ...
func (task *Task) SetDocumentation() {
	task.Documentation = make([]attributes.Documentation, 1)
}

// SetDataInputAssociation ...
func (task *Task) SetDataInputAssociation(num int) {
	task.DataInputAssociation = make([]flow.DataInputAssociation, num)
}

// SetIncoming ...
func (task *Task) SetIncoming(num int) {
	task.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (task *Task) SetOutgoing(num int) {
	task.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
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

// GetProperty ...
func (task Task) GetProperty() *attributes.Property {
	return &task.Property[0]
}

// GetDocumentation ...
func (task Task) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &task.Documentation[0]
}

// GetDataInputAssociation ...
func (task Task) GetDataInputAssociation(num int) *flow.DataInputAssociation {
	return &task.DataInputAssociation[num]
}

// GetIncoming ...
func (task Task) GetIncoming(num int) *marker.Incoming {
	return &task.Incoming[num]
}

// GetOutgoing ...
func (task Task) GetOutgoing(num int) *marker.Outgoing {
	return &task.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (task Task) String() string {
	return fmt.Sprintf("id=%v, name=%v", task.ID, task.Name)
}
