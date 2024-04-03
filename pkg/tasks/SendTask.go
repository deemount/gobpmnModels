package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewSendTask ...
func NewSendTask() SendTaskRepository {
	return &SendTask{}
}

/*
 * @Setters
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

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (sendTask *SendTask) SetDocumentation() {
	sendTask.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

// SetExtensionElements ...
func (sendTask *SendTask) SetExtensionElements() {
	sendTask.ExtensionElements = make(attributes.EXTENSION_ELEMENTS_SLC, 1)
}

/*** Marker ***/

// SetIncoming ...
func (sendTask *SendTask) SetIncoming(num int) {
	sendTask.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (sendTask *SendTask) SetOutgoing(num int) {
	sendTask.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*
 * @Getters
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

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (sendTask SendTask) GetDocumentation() *attributes.Documentation {
	return &sendTask.Documentation[0]
}

// GetExtensionElements ...
func (sendTask SendTask) GetExtensionElements() *attributes.ExtensionElements {
	return &sendTask.ExtensionElements[0]
}

/*** Marker ***/

// GetIncoming ...
func (sendTask SendTask) GetIncoming(num int) *marker.Incoming {
	return &sendTask.Incoming[num]
}

// GetOutgoing ...
func (sendTask SendTask) GetOutgoing(num int) *marker.Outgoing {
	return &sendTask.Outgoing[num]
}

/*
 * @String
 */

// String ...
func (sendTask SendTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", sendTask.ID, sendTask.Name)
}
