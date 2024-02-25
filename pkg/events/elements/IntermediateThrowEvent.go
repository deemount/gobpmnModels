package elements

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events/definitions"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

// NewIntermediateThrowEvent ...
func NewIntermediateThrowEvent() IntermediateThrowEventRepository {
	return &IntermediateThrowEvent{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetID(typ string, suffix interface{}) {
	intermediateThrowEvent.ID = SetID(typ, suffix)
}

// SetName ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetName(name string) {
	intermediateThrowEvent.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetDocumentation() {
	intermediateThrowEvent.Documentation = make([]attributes.Documentation, 1)
}

/*** Marker ***/

// SetIncoming ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetIncoming(num int) {
	intermediateThrowEvent.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetOutgoing(num int) {
	intermediateThrowEvent.Outgoing = make([]marker.Outgoing, num)
}

/*** Event Definitions ***/

// SetCompensateEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetCompensateEventDefinition() {
	intermediateThrowEvent.CompensateEventDefinition = make([]definitions.CompensateEventDefinition, 1)
}

// SetLinkEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetLinkEventDefinition() {
	intermediateThrowEvent.LinkEventDefinition = make([]definitions.LinkEventDefinition, 1)
}

// SetEscalationEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetEscalationEventDefinition() {
	intermediateThrowEvent.EscalationEventDefinition = make([]definitions.EscalationEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetMessageEventDefinition() {
	intermediateThrowEvent.MessageEventDefinition = make([]definitions.MessageEventDefinition, 1)
}

// SetSignalEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetSignalEventDefinition() {
	intermediateThrowEvent.SignalEventDefinition = make([]definitions.SignalEventDefinition, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (intermediateThrowEvent IntermediateThrowEvent) GetID() impl.STR_PTR {
	return &intermediateThrowEvent.ID
}

// GetName ...
func (intermediateThrowEvent IntermediateThrowEvent) GetName() impl.STR_PTR {
	return &intermediateThrowEvent.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (intermediateThrowEvent IntermediateThrowEvent) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &intermediateThrowEvent.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (intermediateThrowEvent IntermediateThrowEvent) GetIncoming(num int) marker.INCOMING_PTR {
	return &intermediateThrowEvent.Incoming[num]
}

// GetOutgoing ...
func (intermediateThrowEvent IntermediateThrowEvent) GetOutgoing(num int) marker.OUTGOING_PTR {
	return &intermediateThrowEvent.Outgoing[num]
}

/*** Event Definitions ***/

// GetCompensateEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetCompensateEventDefinition() *definitions.CompensateEventDefinition {
	return &intermediateThrowEvent.CompensateEventDefinition[0]
}

// GetLinkEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetLinkEventDefinition() *definitions.LinkEventDefinition {
	return &intermediateThrowEvent.LinkEventDefinition[0]
}

// GetEscalationEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetEscalationEventDefinition() *definitions.EscalationEventDefinition {
	return &intermediateThrowEvent.EscalationEventDefinition[0]
}

// GetMessageEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetMessageEventDefinition() *definitions.MessageEventDefinition {
	return &intermediateThrowEvent.MessageEventDefinition[0]
}

// GetSignalEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetSignalEventDefinition() *definitions.SignalEventDefinition {
	return &intermediateThrowEvent.SignalEventDefinition[0]
}

/*
 * Default String
 */

// String ...
func (intermediateThrowEvent IntermediateThrowEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", intermediateThrowEvent.ID, intermediateThrowEvent.Name)
}

// String ...
func (intermediateThrowEvent TIntermediateThrowEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", intermediateThrowEvent.ID, intermediateThrowEvent.Name)
}
