package definitions

import (
	"github.com/deemount/gobpmnModels/pkg/conditional"
	"github.com/deemount/gobpmnModels/pkg/time"
	impl "github.com/deemount/gobpmnTypes"
)

/*
 * @Base
 */

// DefinitionsGetTerminateBase ...
// @EndEvent only
type DefinitionsGetTerminateBase interface {
	GetTerminateEventDefinition() TERMINATE_EVENT_DEF_PTR
}

// DefinitionsGetElementsBase ...
type DefinitionsGetElementsBase interface {
	GetMessageEventDefinition() MESSAGE_EVENT_DEF_PTR
	GetEscalationEventDefinition() ESCALATION_EVENT_DEF_PTR
	GetErrorEventDefinition() ERROR_EVENT_DEF_PTR
	GetSignalEventDefinition() SIGNAL_EVENT_DEF_PTR
	GetCompensateEventDefinition() COMPENSATE_EVENT_DEF_PTR
}

type DefinitionsGetElements interface {
	DefinitionsGetElementsBase
	// Used @BoundaryEvent
	GetCancelEventDefinition() CANCEL_EVENT_DEF_PTR
	GetTimerEventDefinition() TIMER_EVENT_DEF_PTR
	GetConditionalEventDefinition() CONDITIONAL_EVENT_DEF_PTR
}

// DefinitionsBase ...
type DefinitionsBase interface {
	impl.IFBaseID
	impl.IFBaseName
}

/*
 * @Repositories
 */

// CancelEventDefinitionRepository ...
type CancelEventDefinitionRepository interface {
	impl.IFBaseID
}

// CompensateEventDefinitionRepository ...
type CompensateEventDefinitionRepository interface{}

// ConditionalEventDefinitionRepository ...
type ConditionalEventDefinitionRepository interface {
	impl.IFBaseID

	SetCondition()
	GetCondition() *conditional.Condition
}

// ErrorEventDefinitionRepository ...
type ErrorEventDefinitionRepository interface {
	impl.IFBaseID
}

// EscalationEventDefinitionRepository ...
type EscalationEventDefinitionRepository interface {
	impl.IFBaseID
}

// LinkEventDefinitionRepository ...
type LinkEventDefinitionRepository interface {
	impl.IFBaseID
}

// MessageEventDefinitionRepository ...
type MessageEventDefinitionRepository interface {
	impl.IFBaseID

	SetMsgRef(suffix string)
	GetMsgRef() impl.STR_PTR
}

// SignalEventDefinitionRepository ...
type SignalEventDefinitionRepository interface {
	impl.IFBaseID

	SetSignalRef(suffix string)
	GetSignalRef() impl.STR_PTR
}

// TerminateEventDefinitionRepository ...
type TerminateEventDefinitionRepository interface {
	impl.IFBaseID
}

// TimerEventDefinitionRepository ...
type TimerEventDefinitionRepository interface {
	impl.IFBaseID

	SetTimeDate()
	GetTimeDate() *time.TimeDate
	SetTimeDuration()
	GetTimeDuration() *time.TimeDuration
}
