package elements

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events/definitions"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

// Notice: belongs to definitions package
type EventsSetTerminateBase interface {
	SetTerminateEventDefinition()
}

// Notice: belongs to definitions package
type EventsSetDefinitionsBase interface {
	SetMessageEventDefinition()
	SetEscalationEventDefinition()
	SetErrorEventDefinition()
	SetSignalEventDefinition()
	SetCompensateEventDefinition()
}

// Notice: belongs to definitions package
type EventsSetDefinitions interface {
	EventsSetDefinitionsBase

	SetTimerEventDefinition()
	SetCancelEventDefinition()
	SetConditionalEventDefinition()
}

type EventElementsDefinitions interface {
	definitions.DefinitionsGetElements
	EventsSetDefinitions
}

type EventElementsCoreThrowCatchElements interface {
	SetLinkEventDefinition()
	GetLinkEventDefinition() *definitions.LinkEventDefinition
	SetMessageEventDefinition()
	GetMessageEventDefinition() *definitions.MessageEventDefinition
}

// BoundaryEventRepository ...
type BoundaryEventRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	EventElementsDefinitions
	marker.MarkerOutgoing

	SetAttachedToRef(ref string)
	GetAttachedToRef() *string

	// maybe @deprecated @7.17 execution platform
	// Notice: maybe token out of a older modeler version.
	// Needs a checkout
	SetCancelActivity(cancel bool)
	GetCancelActivity() *bool

	String() string
}

// TBoundaryEventRepository ...
type TBoundaryEventRepository interface {
	String() string
}

// EndEventRepository ...
type EndEventRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	marker.MarkerIncoming
	attributes.AttributesBaseElements
	EventsSetDefinitionsBase
	EventsSetTerminateBase
	definitions.DefinitionsGetElementsBase
	definitions.DefinitionsGetTerminateBase

	String() string
}

// TEndEventRepository ...
type TEndEventRepository interface {
	String() string
}

// IntermediateCatchEventRepository ...
type IntermediateCatchEventRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	attributes.AttributesBaseElements
	marker.MarkerIncomingOutgoing
	EventElementsCoreThrowCatchElements

	SetConditionalEventDefinition()
	GetConditionalEventDefinition() *definitions.ConditionalEventDefinition
	SetTimerEventDefinition()
	GetTimerEventDefinition() *definitions.TimerEventDefinition

	String() string
}

// TIntermediateCatchEventRepository ...
type TIntermediateCatchEventRepository interface {
	String() string
}

// IntermediateThrowEventRepository ...
type IntermediateThrowEventRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	marker.MarkerIncomingOutgoing
	attributes.AttributesBaseElements
	EventElementsCoreThrowCatchElements

	SetCompensateEventDefinition()
	GetCompensateEventDefinition() *definitions.CompensateEventDefinition
	SetEscalationEventDefinition()
	GetEscalationEventDefinition() *definitions.EscalationEventDefinition

	String() string
}

// TIntermediateThrowEventRepository ...
type TIntermediateThrowEventRepository interface {
	String() string
}

// StartEventRepository ...
type StartEventRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	marker.MarkerOutgoing
	attributes.AttributesBaseElements

	SetIsInterrupting(isInterrupt bool)
	GetIsInterrupting() *bool

	SetConditionalEventDefinition()
	GetConditionalEventDefinition() *definitions.ConditionalEventDefinition
	SetTimerEventDefinition()
	GetTimerEventDefinition() *definitions.TimerEventDefinition

	SetMessageEventDefinition()
	GetMessageEventDefinition() *definitions.MessageEventDefinition

	String() string
}

// TStartEventRepository ...
type TStartEventRepository interface {
	String() string
}

// MessageRepository ...
type MessageRepository interface {
	impl.IFBaseID
	impl.IFBaseName
}

// SignalRepository ...
type SignalRepository interface {
	impl.IFBaseID
	impl.IFBaseName
}
