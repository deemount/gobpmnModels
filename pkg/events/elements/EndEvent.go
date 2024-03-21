package elements

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events/definitions"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewEndEvent ...
func NewEndEvent() EndEventRepository {
	return &EndEvent{}
}

/*
 * @Setters
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
	endEvent.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

// SetIncoming ...
func (endEvent *EndEvent) SetIncoming(num int) {
	endEvent.Incoming = make(marker.INCOMING_SLC, num)
}

/*** Event Definitions ***/

// SetCompensateEventDefinition ...
func (endEvent *EndEvent) SetCompensateEventDefinition() {
	endEvent.CompensateEventDefinition = make(definitions.COMPENSATE_EVENT_DEF_SLC, 1)
}

// SetEscalationEventDefinition ...
func (endEvent *EndEvent) SetEscalationEventDefinition() {
	endEvent.EscalationEventDefinition = make(definitions.ESCALATION_EVENT_DEF_SLC, 1)
}

// SetMessageEventDefinition ...
func (endEvent *EndEvent) SetMessageEventDefinition() {
	endEvent.MessageEventDefinition = make(definitions.MESSAGE_EVENT_DEF_SLC, 1)
}

// SetErrorEventDefinition ...
func (endEvent *EndEvent) SetErrorEventDefinition() {
	endEvent.ErrorEventDefinition = make(definitions.ERROR_EVENT_DEF_SLC, 1)
}

// SetSignalEventDefinition ...
func (endEvent *EndEvent) SetSignalEventDefinition() {
	endEvent.SignalEventDefinition = make(definitions.SIGNAL_EVENT_DEF_SLC, 1)
}

// SetTerminateEventDefinition ...
func (endEvent *EndEvent) SetTerminateEventDefinition() {
	endEvent.TerminateEventDefinition = make(definitions.TERMINATE_EVENT_DEF_SLC, 1)
}

/*
 * @Getters
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
func (endEvent EndEvent) GetDocumentation() *attributes.Documentation {
	return &endEvent.Documentation[0]
}

// GetIncoming ...
func (endEvent EndEvent) GetIncoming(num int) *marker.Incoming {
	return &endEvent.Incoming[num]
}

/*** Event Definitions ***/

// GetCompensateEventDefinition ...
func (endEvent EndEvent) GetCompensateEventDefinition() definitions.COMPENSATE_EVENT_DEF_PTR {
	return &endEvent.CompensateEventDefinition[0]
}

// GetEscalationEventDefinition ...
func (endEvent EndEvent) GetEscalationEventDefinition() definitions.ESCALATION_EVENT_DEF_PTR {
	return &endEvent.EscalationEventDefinition[0]
}

// GetMessageEventDefinition ...
func (endEvent EndEvent) GetMessageEventDefinition() definitions.MESSAGE_EVENT_DEF_PTR {
	return &endEvent.MessageEventDefinition[0]
}

// GetErrorEventDefinition ...
func (endEvent EndEvent) GetErrorEventDefinition() definitions.ERROR_EVENT_DEF_PTR {
	return &endEvent.ErrorEventDefinition[0]
}

// GetSignalEventDefinition ...
func (endEvent EndEvent) GetSignalEventDefinition() definitions.SIGNAL_EVENT_DEF_PTR {
	return &endEvent.SignalEventDefinition[0]
}

// GetTerminateEventDefinition ...
func (endEvent EndEvent) GetTerminateEventDefinition() definitions.TERMINATE_EVENT_DEF_PTR {
	return &endEvent.TerminateEventDefinition[0]
}

/*
 * @String
 */

// String ...
func (endEvent EndEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", endEvent.ID, endEvent.Name)
}

// String ...
func (endEvent TEndEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", endEvent.ID, endEvent.Name)
}
