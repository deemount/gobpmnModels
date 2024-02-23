package subprocesses

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/loop"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

func NewCallActivity() CallActivityRepository {
	return &CallActivity{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ca *CallActivity) SetID(typ string, suffix interface{}) {
	ca.ID = SetID(typ, suffix)
}

// SetName ...
func (ca *CallActivity) SetName(name string) {
	ca.Name = name
}

// SetCalledElement ...
func (ca *CallActivity) SetCalledElement(element string) {
	ca.CalledElement = element
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (ca *CallActivity) SetDocumentation() {
	ca.Documentation = make([]attributes.Documentation, 1)
}

/*** Marker ***/

// SetIncoming ...
func (ca *CallActivity) SetIncoming(num int) {
	ca.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (ca *CallActivity) SetOutgoing(num int) {
	ca.Outgoing = make([]marker.Outgoing, num)
}

/*** Loop ***/

// SetStandardLoopCharacteristics ...
func (ca *CallActivity) SetStandardLoopCharacteristics() {
	ca.StandardLoopCharacteristics = make([]loop.StandardLoopCharacteristics, 1)
}

// SetMultiInstanceLoopCharacteristics ...
func (ca *CallActivity) SetMultiInstanceLoopCharacteristics() {
	ca.MultiInstanceLoopCharacteristics = make([]loop.MultiInstanceLoopCharacteristics, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ca CallActivity) GetID() impl.STR_PTR {
	return &ca.ID
}

// GetName ...
func (ca CallActivity) GetName() impl.STR_PTR {
	return &ca.Name
}

// GetCalledElement ...
func (ca CallActivity) GetCalledElement() *string {
	return &ca.CalledElement
}

/*** Make Elements ***/

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (ca CallActivity) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &ca.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (ca CallActivity) GetIncoming(num int) marker.INCOMING_PTR {
	return &ca.Incoming[num]
}

// GetOutgoing ...
func (ca CallActivity) GetOutgoing(num int) marker.OUTGOING_PTR {
	return &ca.Outgoing[num]
}

/*** Loop ***/

// GetStandardLoopCharacteristics ...
func (ca CallActivity) GetStandardLoopCharacteristics() *loop.StandardLoopCharacteristics {
	return &ca.StandardLoopCharacteristics[0]
}

// GetMultiInstanceLoopCharacteristics ...
func (ca CallActivity) GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics {
	return &ca.MultiInstanceLoopCharacteristics[0]
}