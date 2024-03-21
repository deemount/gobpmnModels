package elements

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events/definitions"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewStartEvent ...
func NewStartEvent() StartEventRepository {
	return &StartEvent{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (startEvent *StartEvent) SetID(typ string, suffix interface{}) {
	startEvent.ID = SetID(typ, suffix)
}

// SetName ...
func (startEvent *StartEvent) SetName(name string) {
	startEvent.Name = name
}

// SetIsInterrupting ...
func (startEvent *StartEvent) SetIsInterrupting(isInterrupt bool) {
	startEvent.IsInterrupting = isInterrupt
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (startEvent *StartEvent) SetDocumentation() {
	startEvent.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*** Marker ***/

// SetOutgoing ...
func (startEvent *StartEvent) SetOutgoing(num int) {
	startEvent.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*** Event Definitions ***/

// SetConditionalEventDefinition ...
func (startEvent *StartEvent) SetConditionalEventDefinition() {
	startEvent.ConditionalEventDef = make(definitions.CONDITIONAL_EVENT_DEF_SLC, 1)
}

// SetMessagEventDefinition ...
func (startEvent *StartEvent) SetMessageEventDefinition() {
	startEvent.MessageEventDefinition = make(definitions.MESSAGE_EVENT_DEF_SLC, 1)
}

// SetTimerEventDefinition ...
func (startEvent *StartEvent) SetTimerEventDefinition() {
	startEvent.TimerEventDef = make(definitions.TIMER_EVENT_DEF_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (startEvent StartEvent) GetID() impl.STR_PTR {
	return &startEvent.ID
}

// GetName ...
func (startEvent StartEvent) GetName() impl.STR_PTR {
	return &startEvent.Name
}

// SetIsInterrupting ...
func (startEvent StartEvent) GetIsInterrupting() impl.BOOL_PTR {
	return &startEvent.IsInterrupting
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (startEvent StartEvent) GetDocumentation() *attributes.Documentation {
	return &startEvent.Documentation[0]
}

/*** Marker ***/

// GetOutgoing ...
func (startEvent StartEvent) GetOutgoing(num int) *marker.Outgoing {
	return &startEvent.Outgoing[num]
}

/*** Event Definitions ***/

// GetConditionalEventDefinition ...
func (startEvent StartEvent) GetConditionalEventDefinition() definitions.CONDITIONAL_EVENT_DEF_PTR {
	return &startEvent.ConditionalEventDef[0]
}

// GetMessageEventDefinition ...
func (startEvent StartEvent) GetMessageEventDefinition() definitions.MESSAGE_EVENT_DEF_PTR {
	return &startEvent.MessageEventDefinition[0]
}

// GetTimerEventDefinition ...
func (startEvent StartEvent) GetTimerEventDefinition() definitions.TIMER_EVENT_DEF_PTR {
	return &startEvent.TimerEventDef[0]
}

/*
 * @String
 */

// String ...
func (startEvent StartEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", startEvent.ID, startEvent.Name)
}

// String ...
func (startEvent TStartEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", startEvent.ID, startEvent.Name)
}
