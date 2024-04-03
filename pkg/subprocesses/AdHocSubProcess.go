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

// NewSubprocess ...
func NewAdHocSubProcess() AdHocSubProcessRepository {
	return &AdHocSubProcess{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (adhoc *AdHocSubProcess) SetID(typ string, suffix interface{}) {
	adhoc.ID = SetID(typ, suffix)
}

// SetName ...
func (adhoc *AdHocSubProcess) SetName(name string) {
	adhoc.Name = name
}

// SetTriggeredByEvent ...
func (adhoc *AdHocSubProcess) SetTriggeredByEvent(triggered bool) {
	adhoc.TriggeredByEvent = triggered
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (adhoc *AdHocSubProcess) SetDocumentation() {
	adhoc.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

// SetExtensionElements ...
func (adhoc *AdHocSubProcess) SetExtensionElements() {
	adhoc.ExtensionElements = make(attributes.EXTENSION_ELEMENTS_SLC, 1)
}

/** Marker ***/

// SetIncoming ...
func (adhoc *AdHocSubProcess) SetIncoming(num int) {
	adhoc.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (adhoc *AdHocSubProcess) SetOutgoing(num int) {
	adhoc.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*** Loop ***/

// SetMultiInstaceLoopCharacteristics ...
func (adhoc *AdHocSubProcess) SetMultiInstanceLoopCharacteristics() {
	adhoc.MultiInstanceLoopCharacteristics = make([]loop.MultiInstanceLoopCharacteristics, 1)
}

/*** Events ***/

// SetStartEvent ...
func (adhoc *AdHocSubProcess) SetStartEvent(num int) {
	adhoc.StartEvent = make(events.START_EVENT_SLC, num)
}

// SetEndEvent ...
func (adhoc *AdHocSubProcess) SetEndEvent(num int) {
	adhoc.EndEvent = make(events.END_EVENT_SLC, num)
}

/*** Tasks ***/

// SetBusinessRuleTask ...
func (adhoc *AdHocSubProcess) SetBusinessRuleTask(num int) {
	adhoc.BusinessRuleTask = make(tasks.BUSINESS_RULE_TASK_SLC, num)
}

// SetTask ...
func (adhoc *AdHocSubProcess) SetTask(num int) {
	adhoc.Task = make(tasks.TASK_SLC, num)
}

// SetUserTask ...
func (adhoc *AdHocSubProcess) SetUserTask(num int) {
	adhoc.UserTask = make(tasks.USER_TASK_SLC, num)
}

// SetManualTask ...
func (adhoc *AdHocSubProcess) SetManualTask(num int) {
	adhoc.ManualTask = make(tasks.MANUAL_TASK_SLC, num)
}

// SetReceiveTask ...
func (adhoc *AdHocSubProcess) SetReceiveTask(num int) {
	adhoc.ReceiveTask = make(tasks.RECEIVE_TASK_SLC, num)
}

// SetScriptTask ...
func (adhoc *AdHocSubProcess) SetScriptTask(num int) {
	adhoc.ScriptTask = make(tasks.SCRIPT_TASK_SLC, num)
}

// SetSendTask ...
func (adhoc *AdHocSubProcess) SetSendTask(num int) {
	adhoc.SendTask = make(tasks.SEND_TASK_SLC, num)
}

// SetServiceTask ...
func (adhoc *AdHocSubProcess) SetServiceTask(num int) {
	adhoc.ServiceTask = make(tasks.SERVICE_TASK_SLC, num)
}

/*** Subprocess ***/

// SetSubProcess ...
func (adhoc *AdHocSubProcess) SetSubProcess(num int) {
	adhoc.SubProcess = make(SUBPROCESS_SLC, num)
}

// SetAdHocSubProcess ...
func (adhoc *AdHocSubProcess) SetAdHocSubProcess(num int) {
	adhoc.AdHocSubProcess = make(ADHOC_SUBPROCESS_SLC, num)
}

/*** Gateway ***/

// SetExclusiveGateway
func (adhoc *AdHocSubProcess) SetExclusiveGateway(num int) {
	adhoc.ExclusiveGateway = make(gateways.EXCLUSIVE_GATEWAY_SLC, num)
}

// SetInclsuiveGateway
func (adhoc *AdHocSubProcess) SetInclusiveGateway(num int) {
	adhoc.InclusiveGateway = make(gateways.INCLUSIVE_GATEWAY_SLC, num)
}

// SetParallelGateway
func (adhoc *AdHocSubProcess) SetParallelGateway(num int) {
	adhoc.ParallelGateway = make(gateways.PARALLEL_GATEWAY_SLC, num)
}

// SetComplexGateway
func (adhoc *AdHocSubProcess) SetComplexGateway(num int) {
	adhoc.ComplexGateway = make(gateways.COMPLEX_GATEWAY_SLC, num)
}

// SetEventBasedGateway
func (adhoc *AdHocSubProcess) SetEventBasedGateway(num int) {
	adhoc.EventBasedGateway = make(gateways.EVENT_BASED_GATEWAYS_SLC, num)
}

/*** Marker ***/

// SetSequenceFlow ...
func (adhoc *AdHocSubProcess) SetSequenceFlow(num int) {
	adhoc.SequenceFlow = make(flow.SEQUENCE_FLOW_SLC, num)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (adhoc *AdHocSubProcess) GetID() impl.STR_PTR {
	return &adhoc.ID
}

// GetName ...
func (adhoc AdHocSubProcess) GetName() impl.STR_PTR {
	return &adhoc.Name
}

// GetTriggeredByEvent ...
func (adhoc AdHocSubProcess) GetTriggeredByEvent() impl.BOOL_PTR {
	return &adhoc.TriggeredByEvent
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (adhoc AdHocSubProcess) GetDocumentation() *attributes.Documentation {
	return &adhoc.Documentation[0]
}

// GetExtensionElements ...
func (adhoc AdHocSubProcess) GetExtensionElements() *attributes.ExtensionElements {
	return &adhoc.ExtensionElements[0]
}

/*** Marker ***/

// GetIncoming ...
func (adhoc AdHocSubProcess) GetIncoming(num int) *marker.Incoming {
	return &adhoc.Incoming[num]
}

// GetOutgoing ...
func (adhoc AdHocSubProcess) GetOutgoing(num int) *marker.Outgoing {
	return &adhoc.Outgoing[num]
}

/*** Loop ***/

// GetMultiInstaceLoopCharacteristics ...
func (adhoc AdHocSubProcess) GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics {
	return &adhoc.MultiInstanceLoopCharacteristics[0]
}

/*** Events ***/

// GetStartEvent ...
func (adhoc AdHocSubProcess) GetStartEvent(num int) *elements.StartEvent {
	return &adhoc.StartEvent[num]
}

// GetEndEvent ...
func (adhoc AdHocSubProcess) GetEndEvent(num int) *elements.EndEvent {
	return &adhoc.EndEvent[num]
}

/*** Tasks ***/

// GetBusinessRuleTask ...
func (adhoc AdHocSubProcess) GetBusinessRuleTask(num int) *tasks.BusinessRuleTask {
	return &adhoc.BusinessRuleTask[num]
}

// GetTask ...
func (adhoc AdHocSubProcess) GetTask(num int) *tasks.Task {
	return &adhoc.Task[num]
}

// GetUserTask ...
func (adhoc AdHocSubProcess) GetUserTask(num int) *tasks.UserTask {
	return &adhoc.UserTask[num]
}

// GetManualTask ...
func (adhoc AdHocSubProcess) GetManualTask(num int) *tasks.ManualTask {
	return &adhoc.ManualTask[num]
}

// GetReceiveTask ...
func (adhoc AdHocSubProcess) GetReceiveTask(num int) *tasks.ReceiveTask {
	return &adhoc.ReceiveTask[num]
}

// GetScriptTask ...
func (adhoc AdHocSubProcess) GetScriptTask(num int) *tasks.ScriptTask {
	return &adhoc.ScriptTask[num]
}

// GetSendTask ...
func (adhoc AdHocSubProcess) GetSendTask(num int) *tasks.SendTask {
	return &adhoc.SendTask[num]
}

// GetServiceTask ...
func (adhoc AdHocSubProcess) GetServiceTask(num int) *tasks.ServiceTask {
	return &adhoc.ServiceTask[num]
}

/*** Subprocess ***/

// GetSubProcess ...
func (adhoc AdHocSubProcess) GetSubProcess(num int) *SubProcess {
	return &adhoc.SubProcess[num]
}

// GetAdHocSubProcess ...
func (adhoc AdHocSubProcess) GetAdHocSubProcess(num int) *AdHocSubProcess {
	return &adhoc.AdHocSubProcess[num]
}

/*** Gateway ***/

// GetExclusiveGateway
func (adhoc AdHocSubProcess) GetExclusiveGateway(num int) *gateways.ExclusiveGateway {
	return &adhoc.ExclusiveGateway[num]
}

// GetInclsuiveGateway
func (adhoc AdHocSubProcess) GetInclusiveGateway(num int) *gateways.InclusiveGateway {
	return &adhoc.InclusiveGateway[num]
}

// GetParallelGateway
func (adhoc AdHocSubProcess) GetParallelGateway(num int) *gateways.ParallelGateway {
	return &adhoc.ParallelGateway[num]
}

// GetComplexGateway
func (adhoc AdHocSubProcess) GetComplexGateway(num int) *gateways.ComplexGateway {
	return &adhoc.ComplexGateway[num]
}

// GetEventBasedGateway
func (adhoc AdHocSubProcess) GetEventBasedGateway(num int) *gateways.EventBasedGateway {
	return &adhoc.EventBasedGateway[num]
}

/*** Marker ***/

// GetSequenceFlow ...
func (adhoc AdHocSubProcess) GetSequenceFlow(num int) *flow.SequenceFlow {
	return &adhoc.SequenceFlow[num]
}
