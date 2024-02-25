package elements

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events/definitions"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

// NewEndEvent ...
func NewEndEvent() EndEventRepository {
	return &EndEvent{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (endEvent *EndEvent) SetID(typ string, suffix interface{}) {
	endEvent.ID = SetID(typ, suffix)
}

// SetName ...
func (endEvent *EndEvent) SetName(name string) {
	endEvent.Name = name
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (endEvent *EndEvent) SetDocumentation() {
	endEvent.Documentation = make([]attributes.Documentation, 1)
}

// SetIncoming ...
func (endEvent *EndEvent) SetIncoming(num int) {
	endEvent.Incoming = make([]marker.Incoming, num)
}

/*** Event Definitions ***/

// SetCompensateEventDefinition ...
func (endEvent *EndEvent) SetCompensateEventDefinition() {
	endEvent.CompensateEventDefinition = make([]definitions.CompensateEventDefinition, 1)
}

// SetEscalationEventDefinition ...
func (endEvent *EndEvent) SetEscalationEventDefinition() {
	endEvent.EscalationEventDefinition = make([]definitions.EscalationEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (endEvent *EndEvent) SetMessageEventDefinition() {
	endEvent.MessageEventDefinition = make([]definitions.MessageEventDefinition, 1)
}

// SetErrorEventDefinition ...
func (endEvent *EndEvent) SetErrorEventDefinition() {
	endEvent.ErrorEventDefinition = make([]definitions.ErrorEventDefinition, 1)
}

// SetSignalEventDefinition ...
func (endEvent *EndEvent) SetSignalEventDefinition() {
	endEvent.SignalEventDefinition = make([]definitions.SignalEventDefinition, 1)
}

// SetTerminateEventDefinition ...
func (endEvent *EndEvent) SetTerminateEventDefinition() {
	endEvent.TerminateEventDefinition = make([]definitions.TerminateEventDefinition, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (endEvent EndEvent) GetID() impl.STR_PTR {
	return &endEvent.ID
}

// GetName ...
func (endEvent EndEvent) GetName() impl.STR_PTR {
	return &endEvent.Name
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (endEvent EndEvent) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &endEvent.Documentation[0]
}

// GetIncoming ...
func (endEvent EndEvent) GetIncoming(num int) marker.INCOMING_PTR {
	return &endEvent.Incoming[num]
}

/*** Event Definitions ***/

// GetCompensateEventDefinition ...
func (endEvent EndEvent) GetCompensateEventDefinition() *definitions.CompensateEventDefinition {
	return &endEvent.CompensateEventDefinition[0]
}

// GetEscalationEventDefinition ...
func (endEvent EndEvent) GetEscalationEventDefinition() *definitions.EscalationEventDefinition {
	return &endEvent.EscalationEventDefinition[0]
}

// GetMessageEventDefinition ...
func (endEvent EndEvent) GetMessageEventDefinition() *definitions.MessageEventDefinition {
	return &endEvent.MessageEventDefinition[0]
}

// GetErrorEventDefinition ...
func (endEvent EndEvent) GetErrorEventDefinition() *definitions.ErrorEventDefinition {
	return &endEvent.ErrorEventDefinition[0]
}

// GetSignalEventDefinition ...
func (endEvent EndEvent) GetSignalEventDefinition() *definitions.SignalEventDefinition {
	return &endEvent.SignalEventDefinition[0]
}

// GetTerminateEventDefinition ...
func (endEvent EndEvent) GetTerminateEventDefinition() *definitions.TerminateEventDefinition {
	return &endEvent.TerminateEventDefinition[0]
}

/*
 * Default String
 */

// String ...
func (endEvent EndEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", endEvent.ID, endEvent.Name)
}

// String ...
func (endEvent TEndEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", endEvent.ID, endEvent.Name)
}
