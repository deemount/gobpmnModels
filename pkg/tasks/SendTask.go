package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

func NewSendTask() SendTaskRepository {
	return &SendTask{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (sendTask *SendTask) SetID(typ string, suffix interface{}) {

	sendTask.ID = SetID(typ, suffix)
}

// SetName ...
func (sendTask *SendTask) SetName(name string) {
	sendTask.Name = name
}

/* Elements */

// SetDocumentation ...
func (sendTask *SendTask) SetDocumentation() {
	sendTask.Documentation = make([]attributes.Documentation, 1)
}

// SetIncoming ...
func (sendTask *SendTask) SetIncoming(num int) {
	sendTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (sendTask *SendTask) SetOutgoing(num int) {
	sendTask.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (sendTask SendTask) GetID() impl.STR_PTR {
	return &sendTask.ID
}

// GetName ...
func (sendTask SendTask) GetName() impl.STR_PTR {
	return &sendTask.Name
}

/* Elements */

// GetDocumentation ...
func (sendTask SendTask) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &sendTask.Documentation[0]
}

// GetIncoming ...
func (sendTask SendTask) GetIncoming(num int) *marker.Incoming {
	return &sendTask.Incoming[num]
}

// GetOutgoing ...
func (sendTask SendTask) GetOutgoing(num int) *marker.Outgoing {
	return &sendTask.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (sendTask SendTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", sendTask.ID, sendTask.Name)
}
