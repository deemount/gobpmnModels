package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

// NewReceiveTask ...
func NewReceiveTask() ReceiveTaskRepository {
	return &ReceiveTask{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (receiveTask *ReceiveTask) SetID(typ string, suffix interface{}) {
	receiveTask.ID = SetID(typ, suffix)
}

// SetName ...
func (receiveTask *ReceiveTask) SetName(name string) {
	receiveTask.Name = name
}

// SetMessageRef ...
func (receiveTask *ReceiveTask) SetMessageRef(suffix string) {
	receiveTask.MessageRef = fmt.Sprintf("Message_%s", suffix)
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (receiveTask *ReceiveTask) SetDocumentation() {
	receiveTask.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*** Marker ***/

// SetIncoming ...
func (receiveTask *ReceiveTask) SetIncoming(num int) {
	receiveTask.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (receiveTask *ReceiveTask) SetOutgoing(num int) {
	receiveTask.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (receiveTask ReceiveTask) GetID() impl.STR_PTR {
	return &receiveTask.ID
}

// GetName ...
func (receiveTask ReceiveTask) GetName() impl.STR_PTR {
	return &receiveTask.Name
}

// GetMessageRef ...
func (receiveTask ReceiveTask) GetMessageRef(suffix string) impl.STR_PTR {
	return &receiveTask.MessageRef
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (receiveTask ReceiveTask) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &receiveTask.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (receiveTask ReceiveTask) GetIncoming(num int) marker.INCOMING_PTR {
	return &receiveTask.Incoming[num]
}

// GetOutgoing ...
func (receiveTask ReceiveTask) GetOutgoing(num int) marker.OUTGOING_PTR {
	return &receiveTask.Outgoing[num]
}

/*
 * @String
 */

// String ...
func (receiveTask ReceiveTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", receiveTask.ID, receiveTask.Name)
}
