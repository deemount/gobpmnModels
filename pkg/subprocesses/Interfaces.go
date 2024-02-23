package subprocesses

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events"
	"github.com/deemount/gobpmnModels/pkg/events/elements"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/gateways"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/loop"
	"github.com/deemount/gobpmnModels/pkg/marker"
	"github.com/deemount/gobpmnModels/pkg/tasks"
)

type SubprocessesBase interface {
	impl.IFBaseID
	impl.IFBaseName
	attributes.AttributesBaseElements
	marker.MarkerIncomingOutgoing
}

// SubprocessesElementsRepository ...
type SubprocessesElementsRepository interface {
	SetCallActivity(num int)
	GetCallActivity(num int) *CallActivity
	SetSubProcess(num int)
	GetSubProcess(num int) *SubProcess
	SetTransaction(num int)
	GetTransaction(num int) *Transaction
	SetAdHocSubProcess(num int)
	GetAdHocSubProcess(num int) *AdHocSubProcess
}

// AdHocSubProcessRepository ...
type AdHocSubProcessRepository interface {
	SubprocessesBase
	flow.FlowSequenceFlow
	gateways.GatewaysElementsRepository
	tasks.TasksElementsRepository

	SetTriggeredByEvent(triggered bool)
	GetTriggeredByEvent() *bool

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics

	SetStartEvent(num int)
	GetStartEvent(num int) *elements.StartEvent
	SetEndEvent(num int)
	GetEndEvent(num int) *elements.EndEvent

	SetSubProcess(num int)
	GetSubProcess(num int) *SubProcess
	SetAdHocSubProcess(num int)
	GetAdHocSubProcess(num int) *AdHocSubProcess
}

// CallActivityRepository ...
type CallActivityRepository interface {
	SubprocessesBase

	SetCalledElement(element string)
	GetCalledElement() *string

	SetStandardLoopCharacteristics()
	GetStandardLoopCharacteristics() *loop.StandardLoopCharacteristics

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics
}

// SubProcessRepository ...
type SubProcessRepository interface {
	SubprocessesBase
	flow.FlowSequenceFlow
	tasks.TasksElementsRepository
	gateways.GatewaysElementsRepository
	events.ProcessEventsElementsRepository
	SubprocessesElementsRepository

	SetTriggeredByEvent(triggered bool)
	GetTriggeredByEvent() *bool

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics
}

// TransactionRepository ...
type TransactionRepository interface {
	SubprocessesBase
	flow.FlowSequenceFlow
}
