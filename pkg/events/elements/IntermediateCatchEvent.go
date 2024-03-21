package elements

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events/definitions"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewIntermediateCatchEvent ...
func NewIntermediateCatchEvent() IntermediateCatchEventRepository {
	return &IntermediateCatchEvent{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ice *IntermediateCatchEvent) SetID(typ string, suffix interface{}) {
	ice.ID = SetID(typ, suffix)
}

// SetName ...
func (ice *IntermediateCatchEvent) SetName(name string) {
	ice.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (ice *IntermediateCatchEvent) SetDocumentation() {
	ice.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*** Marker ***/

// SetIncoming ...
func (ice *IntermediateCatchEvent) SetIncoming(num int) {
	ice.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (ice *IntermediateCatchEvent) SetOutgoing(num int) {
	ice.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*** Event Definitions ***/

// SetConditionalEventDefinition ...
func (ice *IntermediateCatchEvent) SetConditionalEventDefinition() {
	ice.ConditionalEventDefinition = make(definitions.CONDITIONAL_EVENT_DEF_SLC, 1)
}

// SetLinkEventDefinition ...
func (ice *IntermediateCatchEvent) SetLinkEventDefinition() {
	ice.LinkEventDefinition = make(definitions.LINK_EVENT_DEF_SLC, 1)
}

// SetTimerEventDefinition ...
func (ice *IntermediateCatchEvent) SetTimerEventDefinition() {
	ice.TimerEventDefinition = make(definitions.TIMER_EVENT_DEF_SLC, 1)
}

// SetMessageEventDefinition ...
func (ice *IntermediateCatchEvent) SetMessageEventDefinition() {
	ice.MessageEventDefinition = make(definitions.MESSAGE_EVENT_DEF_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ice IntermediateCatchEvent) GetID() impl.STR_PTR {
	return &ice.ID
}

// GetName ...
func (ice IntermediateCatchEvent) GetName() impl.STR_PTR {
	return &ice.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (ice IntermediateCatchEvent) GetDocumentation() *attributes.Documentation {
	return &ice.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (ice IntermediateCatchEvent) GetIncoming(num int) *marker.Incoming {
	return &ice.Incoming[num]
}

// GetOutgoing ...
func (ice IntermediateCatchEvent) GetOutgoing(num int) *marker.Outgoing {
	return &ice.Outgoing[num]
}

/*** Event Definitions ***/

// GetConditionalEventDefinition ...
func (ice IntermediateCatchEvent) GetConditionalEventDefinition() definitions.CONDITIONAL_EVENT_DEF_PTR {
	return &ice.ConditionalEventDefinition[0]
}

// GetLinkEventDefinition ...
func (ice IntermediateCatchEvent) GetLinkEventDefinition() definitions.LINK_EVENT_DEF_PTR {
	return &ice.LinkEventDefinition[0]
}

// GetTimerEventDefinition ...
func (ice IntermediateCatchEvent) GetTimerEventDefinition() definitions.TIMER_EVENT_DEF_PTR {
	return &ice.TimerEventDefinition[0]
}

// GetMessageEventDefinition ...
func (ice IntermediateCatchEvent) GetMessageEventDefinition() definitions.MESSAGE_EVENT_DEF_PTR {
	return &ice.MessageEventDefinition[0]
}

/*
 * @String
 */

// String ...
func (ice IntermediateCatchEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", ice.ID, ice.Name)
}

// String ...
func (ice TIntermediateCatchEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", ice.ID, ice.Name)
}
