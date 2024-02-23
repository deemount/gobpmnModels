package elements

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events/definitions"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

// NewStartEvent ...
func NewStartEvent() StartEventRepository {
	return &StartEvent{}
}

/*
 * Default Setters
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

/*** Make Elements ***/

// SetDocumentation ...
func (startEvent *StartEvent) SetDocumentation() {
	startEvent.Documentation = make([]attributes.Documentation, 1)
}

// SetConditionalEventDefinition ...
func (startEvent *StartEvent) SetConditionalEventDefinition() {
	startEvent.ConditionalEventDef = make([]definitions.ConditionalEventDefinition, 1)
}

// SetMessagEventDefinition ...
func (startEvent *StartEvent) SetMessageEventDefinition() {
	startEvent.MessageEventDefinition = make([]definitions.MessageEventDefinition, 1)
}

// SetTimerEventDefinition ...
func (startEvent *StartEvent) SetTimerEventDefinition() {
	startEvent.TimerEventDef = make([]definitions.TimerEventDefinition, 1)
}

// SetOutgoing ...
func (startEvent *StartEvent) SetOutgoing(num int) {
	startEvent.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
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
func (startEvent StartEvent) GetIsInterrupting() *bool {
	return &startEvent.IsInterrupting
}

/* Elements */

// GetDocumentation ...
func (startEvent StartEvent) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &startEvent.Documentation[0]
}

// GetConditionalEventDefinition ...
func (startEvent StartEvent) GetConditionalEventDefinition() *definitions.ConditionalEventDefinition {
	return &startEvent.ConditionalEventDef[0]
}

// GetMessageEventDefinition ...
func (startEvent StartEvent) GetMessageEventDefinition() *definitions.MessageEventDefinition {
	return &startEvent.MessageEventDefinition[0]
}

// GetTimerEventDefinition ...
func (startEvent StartEvent) GetTimerEventDefinition() *definitions.TimerEventDefinition {
	return &startEvent.TimerEventDef[0]
}

// GetOutgoing ...
func (startEvent StartEvent) GetOutgoing(num int) marker.OUTGOING_PTR {
	return &startEvent.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (startEvent StartEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", startEvent.ID, startEvent.Name)
}

// String ...
func (startEvent TStartEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", startEvent.ID, startEvent.Name)
}
