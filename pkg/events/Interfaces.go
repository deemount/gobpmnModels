package events

import (
	"github.com/deemount/gobpmnModels/pkg/events/elements"
)

/*
 * @Repositories
 */

// ProcessEventsElementsRepository ...
type ProcessEventsElementsRepository interface {
	SetStartEvent(num int)
	GetStartEvent(num int) *elements.StartEvent
	SetBoundaryEvent(num int)
	GetBoundaryEvent(num int) *elements.BoundaryEvent
	SetEndEvent(num int)
	GetEndEvent(num int) END_EVENT_PTR
	SetIntermediateCatchEvent(num int)
	GetIntermediateCatchEvent(num int) *elements.IntermediateCatchEvent
	SetIntermediateThrowEvent(num int)
	GetIntermediateThrowEvent(num int) *elements.IntermediateThrowEvent
}

// CoreEventsElementsRepository ...
type CoreEventsElementsRepository interface {
	SetSignal(num int)
	GetSignal(num int) SIGNAL_PTR
	SetMessage(num int)
	GetMessage(num int) MESSAGE_PTR
}
