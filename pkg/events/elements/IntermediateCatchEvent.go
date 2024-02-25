package elements

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events/definitions"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

// NewIntermediateCatchEvent ...
func NewIntermediateCatchEvent() IntermediateCatchEventRepository {
	return &IntermediateCatchEvent{}
}

/*
 * Default Setters
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
	ice.Documentation = make([]attributes.Documentation, 1)
}

/*** Marker ***/

// SetIncoming ...
func (ice *IntermediateCatchEvent) SetIncoming(num int) {
	ice.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (ice *IntermediateCatchEvent) SetOutgoing(num int) {
	ice.Outgoing = make([]marker.Outgoing, num)
}

/*** Event Definitions ***/

// SetConditionalEventDefinition ...
func (ice *IntermediateCatchEvent) SetConditionalEventDefinition() {
	ice.ConditionalEventDefinition = make([]definitions.ConditionalEventDefinition, 1)
}

// SetLinkEventDefinition ...
func (ice *IntermediateCatchEvent) SetLinkEventDefinition() {
	ice.LinkEventDefinition = make([]definitions.LinkEventDefinition, 1)
}

// SetTimerEventDefinition ...
func (ice *IntermediateCatchEvent) SetTimerEventDefinition() {
	ice.TimerEventDefinition = make([]definitions.TimerEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (ice *IntermediateCatchEvent) SetMessageEventDefinition() {
	ice.MessageEventDefinition = make([]definitions.MessageEventDefinition, 1)
}

/*
 * Default Getters
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
func (ice IntermediateCatchEvent) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &ice.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (ice IntermediateCatchEvent) GetIncoming(num int) marker.INCOMING_PTR {
	return &ice.Incoming[num]
}

// GetOutgoing ...
func (ice IntermediateCatchEvent) GetOutgoing(num int) marker.OUTGOING_PTR {
	return &ice.Outgoing[num]
}

/*** Event Definitions ***/

// GetConditionalEventDefinition ...
func (ice IntermediateCatchEvent) GetConditionalEventDefinition() *definitions.ConditionalEventDefinition {
	return &ice.ConditionalEventDefinition[0]
}

// GetLinkEventDefinition ...
func (ice IntermediateCatchEvent) GetLinkEventDefinition() *definitions.LinkEventDefinition {
	return &ice.LinkEventDefinition[0]
}

// GetTimerEventDefinition ...
func (ice IntermediateCatchEvent) GetTimerEventDefinition() *definitions.TimerEventDefinition {
	return &ice.TimerEventDefinition[0]
}

// GetMessageEventDefinition ...
func (ice IntermediateCatchEvent) GetMessageEventDefinition() *definitions.MessageEventDefinition {
	return &ice.MessageEventDefinition[0]
}

/*
 * Default String
 */

// String ...
func (ice IntermediateCatchEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", ice.ID, ice.Name)
}

// String ...
func (ice TIntermediateCatchEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", ice.ID, ice.Name)
}
