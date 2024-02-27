package subprocesses

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/gateways"
	"github.com/deemount/gobpmnModels/pkg/loop"
	"github.com/deemount/gobpmnModels/pkg/marker"
	"github.com/deemount/gobpmnModels/pkg/tasks"
	impl "github.com/deemount/gobpmnTypes"
)

// NewSubProcess ...
func NewSubProcess() SubProcessRepository {
	return &SubProcess{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (subprocess *SubProcess) SetID(typ string, suffix interface{}) {
	subprocess.ID = SetID(typ, suffix)
}

// SetName ...
func (subprocess *SubProcess) SetName(name string) {
	subprocess.Name = name
}

// SetTriggeredByEvent ...
func (subprocess *SubProcess) SetTriggeredByEvent(triggered bool) {
	subprocess.TriggeredByEvent = triggered
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (subprocess *SubProcess) SetDocumentation() {
	subprocess.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*** Loop ***/

// SetMultiInstaceLoopCharacteristics ...
func (subprocess *SubProcess) SetMultiInstanceLoopCharacteristics() {
	subprocess.MultiInstanceLoopCharacteristics = make([]loop.MultiInstanceLoopCharacteristics, 1)
}

/*** Events ***/

// SetStartEvent ...
func (subprocess *SubProcess) SetStartEvent(num int) {
	subprocess.StartEvent = make(events.START_EVENT_SLC, num)
}

// SetEndEvent ...
func (subprocess *SubProcess) SetEndEvent(num int) {
	subprocess.EndEvent = make(events.END_EVENT_SLC, num)
}

// SetBoundaryEvent ...
func (subprocess *SubProcess) SetBoundaryEvent(num int) {
	subprocess.BoundaryEvent = make(events.BOUNDARY_EVENT_SLC, num)
}

// SetIntermediateCatchEvent ...
func (subprocess *SubProcess) SetIntermediateCatchEvent(num int) {
	subprocess.IntermediateCatchEvent = make(events.INTERMEDIATE_CATCH_EVENT_SLC, num)
}

// SetIntermediateThrowEvent ...
func (subprocess *SubProcess) SetIntermediateThrowEvent(num int) {
	subprocess.IntermediateThrowEvent = make(events.INTERMEDIATE_THROW_EVENT_SLC, num)
}

/*** Tasks ***/

// SetBusinessRuleTask ...
func (subprocess *SubProcess) SetBusinessRuleTask(num int) {
	subprocess.BusinessRuleTask = make(tasks.BUSINESS_RULE_TASK_SLC, num)
}

// SetTask ...
func (subprocess *SubProcess) SetTask(num int) {
	subprocess.Task = make(tasks.TASK_SLC, num)
}

// SetUserTask ...
func (subprocess *SubProcess) SetUserTask(num int) {
	subprocess.UserTask = make(tasks.USER_TASK_SLC, num)
}

// SetManualTask ...
func (subprocess *SubProcess) SetManualTask(num int) {
	subprocess.ManualTask = make(tasks.MANUAL_TASK_SLC, num)
}

// SetReceiveTask ...
func (subprocess *SubProcess) SetReceiveTask(num int) {
	subprocess.ReceiveTask = make(tasks.RECEIVE_TASK_SLC, num)
}

// SetScriptTask ...
func (subprocess *SubProcess) SetScriptTask(num int) {
	subprocess.ScriptTask = make(tasks.SCRIPT_TASK_SLC, num)
}

// SetSendTask ...
func (subprocess *SubProcess) SetSendTask(num int) {
	subprocess.SendTask = make(tasks.SEND_TASK_SLC, num)
}

// SetServiceTask ...
func (subprocess *SubProcess) SetServiceTask(num int) {
	subprocess.ServiceTask = make(tasks.SERVICE_TASK_SLC, num)
}

/*** Subprocess ***/

// SetCallActivity ...
func (subprocess *SubProcess) SetCallActivity(num int) {
	subprocess.CallActivity = make(CALL_ACTIVITY_SLC, num)
}

// SetSubProcess ...
func (subprocess *SubProcess) SetSubProcess(num int) {
	subprocess.SubProcess = make(SUBPROCESS_SLC, num)
}

// SetAdHocSubProcess ...
func (subprocess *SubProcess) SetAdHocSubProcess(num int) {
	subprocess.AdHocSubProcess = make(ADHOC_SUBPROCESS_SLC, num)
}

// SetTransaction ...
func (subprocess *SubProcess) SetTransaction(num int) {
	subprocess.Transaction = make(TRANSACTION_SLC, num)
}

/*** Gateways ***/

// SetExclusiveGateway
func (subprocess *SubProcess) SetExclusiveGateway(num int) {
	subprocess.ExclusiveGateway = make(gateways.EXCLUSIVE_GATEWAY_SLC, num)
}

// SetInclsuiveGateway
func (subprocess *SubProcess) SetInclusiveGateway(num int) {
	subprocess.InclusiveGateway = make(gateways.INCLUSIVE_GATEWAY_SLC, num)
}

// SetParallelGateway
func (subprocess *SubProcess) SetParallelGateway(num int) {
	subprocess.ParallelGateway = make(gateways.PARALLEL_GATEWAY_SLC, num)
}

// SetComplexGateway
func (subprocess *SubProcess) SetComplexGateway(num int) {
	subprocess.ComplexGateway = make(gateways.COMPLEX_GATEWAY_SLC, num)
}

// SetEventBasedGateway
func (subprocess *SubProcess) SetEventBasedGateway(num int) {
	subprocess.EventBasedGateway = make(gateways.EVENT_BASED_GATEWAYS_SLC, num)
}

/*** Marker ***/

// SetIncoming ...
func (subprocess *SubProcess) SetIncoming(num int) {
	subprocess.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (subprocess *SubProcess) SetOutgoing(num int) {
	subprocess.Outgoing = make(marker.OUTGOING_SLC, num)
}

// SetSequenceFlow ...
func (subprocess *SubProcess) SetSequenceFlow(num int) {
	subprocess.SequenceFlow = make(flow.SEQUENCE_FLOW_SLC, num)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (subprocess SubProcess) GetID() impl.STR_PTR {
	return &subprocess.ID
}

// GetName ...
func (subprocess SubProcess) GetName() impl.STR_PTR {
	return &subprocess.Name
}

// GetTriggeredByEvent ...
func (subprocess SubProcess) GetTriggeredByEvent() impl.BOOL_PTR {
	return &subprocess.TriggeredByEvent
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (subprocess SubProcess) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &subprocess.Documentation[0]
}

/*** Loop ***/

// GetMultiInstaceLoopCharacteristics ...
func (subprocess SubProcess) GetMultiInstanceLoopCharacteristics() loop.MULTI_INSTANCE_LOOP_CHARACTERISTICS_PTR {
	return &subprocess.MultiInstanceLoopCharacteristics[0]
}

/*** Event ***/

// GetStartEvent ...
func (subprocess SubProcess) GetStartEvent(num int) events.START_EVENT_PTR {
	return &subprocess.StartEvent[num]
}

// GetEndEvent ...
func (subprocess SubProcess) GetEndEvent(num int) events.END_EVENT_PTR {
	return &subprocess.EndEvent[num]
}

// GetBoundaryEvent ...
func (subprocess SubProcess) GetBoundaryEvent(num int) events.BOUNDARY_EVENT_PTR {
	return &subprocess.BoundaryEvent[num]
}

// GetIntermediateCatchEvent ...
func (subprocess SubProcess) GetIntermediateCatchEvent(num int) events.INTERMEDIATE_CATCH_EVENT_PTR {
	return &subprocess.IntermediateCatchEvent[num]
}

// GetIntermediateThrowEvent ...
func (subprocess SubProcess) GetIntermediateThrowEvent(num int) events.INTERMEDIATE_THROW_EVENT_PTR {
	return &subprocess.IntermediateThrowEvent[num]
}

/*** Tasks ***/

// SetBusinessRuleTask ...
func (subprocess SubProcess) GetBusinessRuleTask(num int) tasks.BUSINESS_RULE_TASK_PTR {
	return &subprocess.BusinessRuleTask[num]
}

// GetTask ...
func (subprocess SubProcess) GetTask(num int) tasks.TASK_PTR {
	return &subprocess.Task[num]
}

// GetUserTask ...
func (subprocess SubProcess) GetUserTask(num int) tasks.USER_TASK_PTR {
	return &subprocess.UserTask[num]
}

// GetManualTask ...
func (subprocess SubProcess) GetManualTask(num int) tasks.MANUAL_TASK_PTR {
	return &subprocess.ManualTask[num]
}

// GetReceiveTask ...
func (subprocess SubProcess) GetReceiveTask(num int) tasks.RECEIVE_TASK_PTR {
	return &subprocess.ReceiveTask[num]
}

// GetScriptTask ...
func (subprocess SubProcess) GetScriptTask(num int) tasks.SCRIPT_TASK_PTR {
	return &subprocess.ScriptTask[num]
}

// GetSendTask ...
func (subprocess SubProcess) GetSendTask(num int) tasks.SEND_TASK_PTR {
	return &subprocess.SendTask[num]
}

// GetServiceTask ...
func (subprocess SubProcess) GetServiceTask(num int) tasks.SERVICE_TASK_PTR {
	return &subprocess.ServiceTask[num]
}

/*** Subprocesses ***/

// GetCallActivity ...
func (subprocess SubProcess) GetCallActivity(num int) CALL_ACTIVITY_PTR {
	return &subprocess.CallActivity[num]
}

// GetSubProcess ...
func (subprocess SubProcess) GetSubProcess(num int) SUBPROCESS_PTR {
	return &subprocess.SubProcess[num]
}

// GetAdHocSubProcess ...
func (subprocess SubProcess) GetAdHocSubProcess(num int) ADHOC_SUBPROCESS_PTR {
	return &subprocess.AdHocSubProcess[num]
}

// GetTransaction ...
func (subprocess SubProcess) GetTransaction(num int) TRANSACTION_PTR {
	return &subprocess.Transaction[num]
}

/*** Gateways ***/

// GetExclusiveGateway
func (subprocess SubProcess) GetExclusiveGateway(num int) gateways.EXCLUSIVE_GATEWAY_PTR {
	return &subprocess.ExclusiveGateway[num]
}

// GetInclsuiveGateway
func (subprocess SubProcess) GetInclusiveGateway(num int) gateways.INCLUSIVE_GATEWAY_PTR {
	return &subprocess.InclusiveGateway[num]
}

// GetParallelGateway
func (subprocess SubProcess) GetParallelGateway(num int) gateways.PARALLEL_GATEWAY_PTR {
	return &subprocess.ParallelGateway[num]
}

// GetComplexGateway
func (subprocess SubProcess) GetComplexGateway(num int) gateways.COMPLEX_GATEWAY_PTR {
	return &subprocess.ComplexGateway[num]
}

// GetEventBasedGateway
func (subprocess SubProcess) GetEventBasedGateway(num int) gateways.EVENT_BASED_GATEWAYS_PTR {
	return &subprocess.EventBasedGateway[num]
}

/*** Marker ***/

// GetIncoming ...
func (subprocess SubProcess) GetIncoming(num int) marker.INCOMING_PTR {
	return &subprocess.Incoming[num]
}

// GetOutgoing ...
func (subprocess SubProcess) GetOutgoing(num int) marker.OUTGOING_PTR {
	return &subprocess.Outgoing[num]
}

// GetSequenceFlow ...
func (subprocess SubProcess) GetSequenceFlow(num int) flow.SEQUENCE_FLOW_PTR {
	return &subprocess.SequenceFlow[num]
}
