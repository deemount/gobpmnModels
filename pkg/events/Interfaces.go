package events

/*
 * @Repositories
 */

// ProcessEventsElementsRepository ...
type ProcessEventsElementsRepository interface {
	SetStartEvent(num int)
	GetStartEvent(num int) START_EVENT_PTR
	SetBoundaryEvent(num int)
	GetBoundaryEvent(num int) BOUNDARY_EVENT_PTR
	SetEndEvent(num int)
	GetEndEvent(num int) END_EVENT_PTR
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
