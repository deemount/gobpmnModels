package events

import "github.com/deemount/gobpmnModels/pkg/events/elements"

/*
 * @Repositories
 */

// ProcessEventsElementsRepository ...
type ProcessEventsElementsRepository interface {
	SetStartEvent(num int)
	GetStartEvent(num int) *elements.StartEvent
	SetBoundaryEvent(num int)
	GetBoundaryEvent(num int) BOUNDARY_EVENT_PTR
	SetEndEvent(num int)
	GetEndEvent(num int) *elements.EndEvent
	SetIntermediateCatchEvent(num int)
	GetIntermediateCatchEvent(num int) INTERMEDIATE_CATCH_EVENT_PTR
	SetIntermediateThrowEvent(num int)
	GetIntermediateThrowEvent(num int) INTERMEDIATE_THROW_EVENT_PTR
}

// CoreEventsElementsRepository ...
type CoreEventsElementsRepository interface {
	SetSignal(num int)
	GetSignal(num int) SIGNAL_PTR
	SetMessage(num int)
	GetMessage(num int) MESSAGE_PTR
}
