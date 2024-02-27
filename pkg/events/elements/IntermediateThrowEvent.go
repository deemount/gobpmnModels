package elements

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events/definitions"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewIntermediateThrowEvent ...
func NewIntermediateThrowEvent() IntermediateThrowEventRepository {
	return &IntermediateThrowEvent{}
}

/*
 * @Setters
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
	intermediateThrowEvent.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*** Marker ***/

// SetIncoming ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetIncoming(num int) {
	intermediateThrowEvent.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetOutgoing(num int) {
	intermediateThrowEvent.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*** Event Definitions ***/

// SetCompensateEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetCompensateEventDefinition() {
	intermediateThrowEvent.CompensateEventDefinition = make(definitions.COMPENSATE_EVENT_DEF_SLC, 1)
}

// SetLinkEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetLinkEventDefinition() {
	intermediateThrowEvent.LinkEventDefinition = make(definitions.LINK_EVENT_DEF_SLC, 1)
}

// SetEscalationEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetEscalationEventDefinition() {
	intermediateThrowEvent.EscalationEventDefinition = make(definitions.ESCALATION_EVENT_DEF_SLC, 1)
}

// SetMessageEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetMessageEventDefinition() {
	intermediateThrowEvent.MessageEventDefinition = make(definitions.MESSAGE_EVENT_DEF_SLC, 1)
}

// SetSignalEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetSignalEventDefinition() {
	intermediateThrowEvent.SignalEventDefinition = make(definitions.SIGNAL_EVENT_DEF_SLC, 1)
}

/*
 * @Getters
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
func (intermediateThrowEvent IntermediateThrowEvent) GetCompensateEventDefinition() definitions.COMPENSATE_EVENT_DEF_PTR {
	return &intermediateThrowEvent.CompensateEventDefinition[0]
}

// GetLinkEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetLinkEventDefinition() definitions.LINK_EVENT_DEF_PTR {
	return &intermediateThrowEvent.LinkEventDefinition[0]
}

// GetEscalationEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetEscalationEventDefinition() definitions.ESCALATION_EVENT_DEF_PTR {
	return &intermediateThrowEvent.EscalationEventDefinition[0]
}

// GetMessageEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetMessageEventDefinition() definitions.MESSAGE_EVENT_DEF_PTR {
	return &intermediateThrowEvent.MessageEventDefinition[0]
}

// GetSignalEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetSignalEventDefinition() definitions.SIGNAL_EVENT_DEF_PTR {
	return &intermediateThrowEvent.SignalEventDefinition[0]
}

/*
 * @String
 */

// String ...
func (intermediateThrowEvent IntermediateThrowEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", intermediateThrowEvent.ID, intermediateThrowEvent.Name)
}

// String ...
func (intermediateThrowEvent TIntermediateThrowEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", intermediateThrowEvent.ID, intermediateThrowEvent.Name)
}
