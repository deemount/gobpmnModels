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
	GetBoundaryEvent(num int) *elements.BoundaryEvent
	SetEndEvent(num int)
	GetEndEvent(num int) *elements.EndEvent
	SetIntermediateCatchEvent(num int)
	GetIntermediateCatchEvent(num int) *elements.IntermediateCatchEvent
	SetIntermediateThrowEvent(num int)
	GetIntermediateThrowEvent(num int) *elements.IntermediateThrowEvent
}

// CoreEventsElementsRepository ...
type CoreEventsElementsRepository interface {
	SetSignal(num int)
	GetSignal(num int) *elements.Signal
	SetMessage(num int)
	GetMessage(num int) *elements.Message
}
