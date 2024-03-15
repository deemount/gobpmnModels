package elements

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/events/definitions"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewBoundaryEvent ...
func NewBoundaryEvent() BoundaryEventRepository {
	return &BoundaryEvent{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (be *BoundaryEvent) SetID(typ string, suffix interface{}) {
	be.ID = SetID(typ, suffix)
}

// SetName ...
func (be *BoundaryEvent) SetName(name string) {
	be.ID = name
}

// SetAttachedToRef ...
func (be *BoundaryEvent) SetAttachedToRef(ref string) {
	be.AttachedToRef = ref
}

// SetCancelActivity ...
func (be *BoundaryEvent) SetCancelActivity(cancel bool) {
	be.CancelActivity = cancel
}

/* Elements */

/** BPMN **/

/*** Event Definition ***/

// SetMessageEventDefinition ...
func (be *BoundaryEvent) SetCancelEventDefinition() {
	be.CancelEventDefinition = make(definitions.CANCEL_EVENT_DEF_SLC, 1)
}

// SetMessageEventDefinition ...
func (be *BoundaryEvent) SetMessageEventDefinition() {
	be.MessageEventDefinition = make(definitions.MESSAGE_EVENT_DEF_SLC, 1)
}

// SetTimerEventDefinition ...
func (be *BoundaryEvent) SetTimerEventDefinition() {
	be.TimerEventDefinition = make(definitions.TIMER_EVENT_DEF_SLC, 1)
}

// SetEscalationEventDefinition ...
func (be *BoundaryEvent) SetEscalationEventDefinition() {
	be.EscalationEventDefinition = make(definitions.ESCALATION_EVENT_DEF_SLC, 1)
}

// SetErrorEventDefinition ...
func (be *BoundaryEvent) SetErrorEventDefinition() {
	be.ErrorEventDefinition = make(definitions.ERROR_EVENT_DEF_SLC, 1)
}

// SetSignalEventDefinition ...
func (be *BoundaryEvent) SetSignalEventDefinition() {
	be.SignalEventDefinition = make(definitions.SIGNAL_EVENT_DEF_SLC, 1)
}

// SetCompensateEventDefinition ...
func (be *BoundaryEvent) SetCompensateEventDefinition() {
	be.CompensateEventDefinition = make(definitions.COMPENSATE_EVENT_DEF_SLC, 1)
}

// SetConditionalEventDefinition ...
func (be *BoundaryEvent) SetConditionalEventDefinition() {
	be.ConditionalEventDefinition = make(definitions.CONDITIONAL_EVENT_DEF_SLC, 1)
}

/*** Marker ***/

// SetOutgoing ...
func (be *BoundaryEvent) SetOutgoing(num int) {
	be.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (be BoundaryEvent) GetID() impl.STR_PTR {
	return &be.ID
}

// GetName ...
func (be BoundaryEvent) GetName() impl.STR_PTR {
	return &be.Name
}

// GetAttachedToRef ...
func (be BoundaryEvent) GetAttachedToRef() impl.STR_PTR {
	return &be.AttachedToRef
}

// GetCancelActivity ...
func (be BoundaryEvent) GetCancelActivity() impl.BOOL_PTR {
	return &be.CancelActivity
}

/* Elements */

/** BPMN **/

/*** Event Definition ***/

// GetMessageEventDefinition ...
func (be BoundaryEvent) GetCancelEventDefinition() definitions.CANCEL_EVENT_DEF_PTR {
	return &be.CancelEventDefinition[0]
}

// GetMessageEventDefinition ...
func (be BoundaryEvent) GetMessageEventDefinition() definitions.MESSAGE_EVENT_DEF_PTR {
	return &be.MessageEventDefinition[0]
}

// GetTimerEventDefinition ...
func (be BoundaryEvent) GetTimerEventDefinition() definitions.TIMER_EVENT_DEF_PTR {
	return &be.TimerEventDefinition[0]
}

// GetEscalationEventDefinition ...
func (be BoundaryEvent) GetEscalationEventDefinition() definitions.ESCALATION_EVENT_DEF_PTR {
	return &be.EscalationEventDefinition[0]
}

// GetErrorEventDefinition ...
func (be BoundaryEvent) GetErrorEventDefinition() definitions.ERROR_EVENT_DEF_PTR {
	return &be.ErrorEventDefinition[0]
}

// GetSignalEventDefinition ...
func (be BoundaryEvent) GetSignalEventDefinition() definitions.SIGNAL_EVENT_DEF_PTR {
	return &be.SignalEventDefinition[0]
}

// GetCompensateEventDefinition ...
func (be BoundaryEvent) GetCompensateEventDefinition() definitions.COMPENSATE_EVENT_DEF_PTR {
	return &be.CompensateEventDefinition[0]
}

// GetConditionalEventDefinition ...
func (be BoundaryEvent) GetConditionalEventDefinition() definitions.CONDITIONAL_EVENT_DEF_PTR {
	return &be.ConditionalEventDefinition[0]
}

/*** Marker ***/

// SetOutgoing ...
func (be BoundaryEvent) GetOutgoing(num int) *marker.Outgoing {
	return &be.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (be BoundaryEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", be.ID, be.Name)
}

// String ...
func (be TBoundaryEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", be.ID, be.Name)
}
