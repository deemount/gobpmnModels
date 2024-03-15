package subprocesses

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events"
	"github.com/deemount/gobpmnModels/pkg/events/elements"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/gateways"
	"github.com/deemount/gobpmnModels/pkg/loop"
	"github.com/deemount/gobpmnModels/pkg/marker"
	"github.com/deemount/gobpmnModels/pkg/tasks"
	impl "github.com/deemount/gobpmnTypes"
)

/*
 * @Base
 */

// SubproceesesBase ...
type SubprocessesBase interface {
	impl.IFBaseID
	impl.IFBaseName
	attributes.AttributesBaseElements
	marker.MarkerIncomingOutgoing
}

/*
 * @Repositories
 */

// SubprocessesElementsRepository ...
type SubprocessesElementsRepository interface {
	SetCallActivity(num int)
	GetCallActivity(num int) CALL_ACTIVITY_PTR
	SetSubProcess(num int)
	GetSubProcess(num int) SUBPROCESS_PTR
	SetTransaction(num int)
	GetTransaction(num int) TRANSACTION_PTR
	SetAdHocSubProcess(num int)
	GetAdHocSubProcess(num int) ADHOC_SUBPROCESS_PTR
}

// AdHocSubProcessRepository ...
type AdHocSubProcessRepository interface {
	SubprocessesBase
	flow.FlowSequenceFlow
	gateways.GatewaysElementsRepository
	tasks.TasksElementsRepository

	SetTriggeredByEvent(triggered bool)
	GetTriggeredByEvent() impl.BOOL_PTR

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() loop.MULTI_INSTANCE_LOOP_CHARACTERISTICS_PTR

	SetStartEvent(num int)
	GetStartEvent(num int) *elements.StartEvent
	SetEndEvent(num int)
	GetEndEvent(num int) *elements.EndEvent

	SetSubProcess(num int)
	GetSubProcess(num int) SUBPROCESS_PTR
	SetAdHocSubProcess(num int)
	GetAdHocSubProcess(num int) ADHOC_SUBPROCESS_PTR
}

// CallActivityRepository ...
type CallActivityRepository interface {
	SubprocessesBase

	SetCalledElement(element string)
	GetCalledElement() impl.STR_PTR

	SetStandardLoopCharacteristics()
	GetStandardLoopCharacteristics() loop.STANDARD_LOOP_CHARACTERISTICS_PTR

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() loop.MULTI_INSTANCE_LOOP_CHARACTERISTICS_PTR
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
	GetTriggeredByEvent() impl.BOOL_PTR

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() loop.MULTI_INSTANCE_LOOP_CHARACTERISTICS_PTR
}

// TransactionRepository ...
type TransactionRepository interface {
	SubprocessesBase
	flow.FlowSequenceFlow
}
