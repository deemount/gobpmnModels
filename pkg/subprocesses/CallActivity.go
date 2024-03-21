package subprocesses

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/loop"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewCallActivity ...
func NewCallActivity() CallActivityRepository {
	return &CallActivity{}
}

/**
 * @Setters
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

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (ca *CallActivity) SetDocumentation() {
	ca.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*** Marker ***/

// SetIncoming ...
func (ca *CallActivity) SetIncoming(num int) {
	ca.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (ca *CallActivity) SetOutgoing(num int) {
	ca.Outgoing = make(marker.OUTGOING_SLC, num)
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
 * @Getters
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
func (ca CallActivity) GetCalledElement() impl.STR_PTR {
	return &ca.CalledElement
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (ca CallActivity) GetDocumentation() *attributes.Documentation {
	return &ca.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (ca CallActivity) GetIncoming(num int) *marker.Incoming {
	return &ca.Incoming[num]
}

// GetOutgoing ...
func (ca CallActivity) GetOutgoing(num int) *marker.Outgoing {
	return &ca.Outgoing[num]
}

/*** Loop ***/

// GetStandardLoopCharacteristics ...
func (ca CallActivity) GetStandardLoopCharacteristics() loop.STANDARD_LOOP_CHARACTERISTICS_PTR {
	return &ca.StandardLoopCharacteristics[0]
}

// GetMultiInstanceLoopCharacteristics ...
func (ca CallActivity) GetMultiInstanceLoopCharacteristics() loop.MULTI_INSTANCE_LOOP_CHARACTERISTICS_PTR {
	return &ca.MultiInstanceLoopCharacteristics[0]
}
