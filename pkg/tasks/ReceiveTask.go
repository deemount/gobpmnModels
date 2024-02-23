package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

func NewReceiveTask() ReceiveTaskRepository {
	return &ReceiveTask{}
}

/**
 * Default Setters
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

/*** Make Elements ***/

// SetDocumentation ...
func (receiveTask *ReceiveTask) SetDocumentation() {
	receiveTask.Documentation = make([]attributes.Documentation, 1)
}

// SetIncoming ...
func (receiveTask *ReceiveTask) SetIncoming(num int) {
	receiveTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (receiveTask *ReceiveTask) SetOutgoing(num int) {
	receiveTask.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
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

// GetDocumentation ...
func (receiveTask ReceiveTask) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &receiveTask.Documentation[0]
}

// GetIncoming ...
func (receiveTask ReceiveTask) GetIncoming(num int) *marker.Incoming {
	return &receiveTask.Incoming[num]
}

// GetOutgoing ...
func (receiveTask ReceiveTask) GetOutgoing(num int) *marker.Outgoing {
	return &receiveTask.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (receiveTask ReceiveTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", receiveTask.ID, receiveTask.Name)
}
