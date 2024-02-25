package definitions

import (
	"github.com/deemount/gobpmnModels/pkg/conditional"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/time"
)

/*
 * @Base
 */

// DefinitionsGetTerminateBase ...
// @EndEvent only
type DefinitionsGetTerminateBase interface {
	GetTerminateEventDefinition() *TerminateEventDefinition
}

// DefinitionsGetElementsBase ...
type DefinitionsGetElementsBase interface {
	GetMessageEventDefinition() *MessageEventDefinition
	GetEscalationEventDefinition() *EscalationEventDefinition
	GetErrorEventDefinition() *ErrorEventDefinition
	GetSignalEventDefinition() *SignalEventDefinition
	GetCompensateEventDefinition() *CompensateEventDefinition
}

type DefinitionsGetElements interface {
	DefinitionsGetElementsBase
	// @BoundaryEvent
	GetCancelEventDefinition() *CancelEventDefinition
	GetTimerEventDefinition() *TimerEventDefinition
	GetConditionalEventDefinition() *ConditionalEventDefinition
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
	GetSignalRef() *string
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
