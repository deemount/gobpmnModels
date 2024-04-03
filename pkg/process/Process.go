package process

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/data"
	"github.com/deemount/gobpmnModels/pkg/events"
	"github.com/deemount/gobpmnModels/pkg/events/elements"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/gateways"
	"github.com/deemount/gobpmnModels/pkg/pool"
	"github.com/deemount/gobpmnModels/pkg/subprocesses"
	"github.com/deemount/gobpmnModels/pkg/tasks"
	impl "github.com/deemount/gobpmnTypes"
)

// NewProcess ...
func NewProcess() ProcessRepository {
	return &Process{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (process *Process) SetID(typ string, suffix interface{}) {
	process.ID = SetID(typ, suffix)
}

// SetName ...
func (process *Process) SetName(name string) {
	process.Name = name
}

// SetIsExecutable ...
func (process *Process) SetIsExecutable(isExec bool) {
	process.IsExecutable = isExec
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (process *Process) SetDocumentation() {
	process.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

// SetExtensionElements ...
func (process *Process) SetExtensionElements() {
	process.ExtensionElements = make(attributes.EXTENSION_ELEMENTS_SLC, 1)
}

/** Pool **/

// SetLaneSet ...
func (process *Process) SetLaneSet() {
	process.LaneSet = make(pool.LANE_SET_SLC, 1)
}

/*** Events ***/

// SetStartEvent ...
func (process *Process) SetStartEvent(num int) {
	process.StartEvent = make(events.START_EVENT_SLC, num)
}

// SetBoundaryEvent ...
func (process *Process) SetBoundaryEvent(num int) {
	process.BoundaryEvent = make(events.BOUNDARY_EVENT_SLC, num)
}

// SetEndEvent ...
func (process *Process) SetEndEvent(num int) {
	process.EndEvent = make(events.END_EVENT_SLC, num)
}

// SetIntermedCatchEvent ...
func (process *Process) SetIntermediateCatchEvent(num int) {
	process.IntermediateCatchEvent = make(events.INTERMEDIATE_CATCH_EVENT_SLC, num)
}

// SetIntermedThrowEvent ...
func (process *Process) SetIntermediateThrowEvent(num int) {
	process.IntermediateThrowEvent = make(events.INTERMEDIATE_THROW_EVENT_SLC, num)
}

/*** Tasks ***/

// SetBusinessRuleTask ...
func (process *Process) SetBusinessRuleTask(num int) {
	process.BusinessRuleTask = make(tasks.BUSINESS_RULE_TASK_SLC, num)
}

// SetTask ...
func (process *Process) SetTask(num int) {
	process.Task = make(tasks.TASK_SLC, num)
}

// SetUserTask ...
func (process *Process) SetUserTask(num int) {
	process.UserTask = make(tasks.USER_TASK_SLC, num)
}

// SetManualTask ...
func (process *Process) SetManualTask(num int) {
	process.ManualTask = make(tasks.MANUAL_TASK_SLC, num)
}

// SetReceiveTask ...
func (process *Process) SetReceiveTask(num int) {
	process.ReceiveTask = make(tasks.RECEIVE_TASK_SLC, num)
}

// SetScriptTask ...
func (process *Process) SetScriptTask(num int) {
	process.ScriptTask = make(tasks.SCRIPT_TASK_SLC, num)
}

// SetSendTask ...
func (process *Process) SetSendTask(num int) {
	process.SendTask = make(tasks.SEND_TASK_SLC, num)
}

// SetServiceTask ...
func (process *Process) SetServiceTask(num int) {
	process.ServiceTask = make(tasks.SERVICE_TASK_SLC, num)
}

/*** Subprocesses ***/

// SetCallActivity ...
func (process *Process) SetCallActivity(num int) {
	process.CallActivity = make(subprocesses.CALL_ACTIVITY_SLC, num)
}

// SetSubProcess ...
func (process *Process) SetSubProcess(num int) {
	process.SubProcess = make(subprocesses.SUBPROCESS_SLC, num)
}

// SetTransaction ...
func (process *Process) SetTransaction(num int) {
	process.Transaction = make(subprocesses.TRANSACTION_SLC, num)
}

// SetAdHocSubProcess ...
func (process *Process) SetAdHocSubProcess(num int) {
	process.AdHocSubProcess = make(subprocesses.ADHOC_SUBPROCESS_SLC, num)
}

/*** Gateways ***/

// SetExclusiveGateway
func (process *Process) SetExclusiveGateway(num int) {
	process.ExclusiveGateway = make(gateways.EXCLUSIVE_GATEWAY_SLC, num)
}

// SetInclsuiveGateway
func (process *Process) SetInclusiveGateway(num int) {
	process.InclusiveGateway = make(gateways.INCLUSIVE_GATEWAY_SLC, num)
}

// SetParallelGateway
func (process *Process) SetParallelGateway(num int) {
	process.ParallelGateway = make(gateways.PARALLEL_GATEWAY_SLC, num)
}

// SetComplexGateway
func (process *Process) SetComplexGateway(num int) {
	process.ComplexGateway = make(gateways.COMPLEX_GATEWAY_SLC, num)
}

// SetEventBasedGateway
func (process *Process) SetEventBasedGateway(num int) {
	process.EventBasedGateway = make(gateways.EVENT_BASED_GATEWAYS_SLC, num)
}

/*** Marker ***/

// SetAssociation ...
func (process *Process) SetAssociation(num int) {
	process.Association = make(flow.ASSOCIATION_SLC, num)
}

// SetSequenceFlow ...
func (process *Process) SetSequenceFlow(num int) {
	process.SequenceFlow = make(flow.SEQUENCE_FLOW_SLC, num)
}

/*** Data ***/

// SetDataObject ...
func (process *Process) SetDataObject(num int) {
	process.DataObject = make([]data.DataObject, num)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (process Process) GetID() impl.STR_PTR {
	return &process.ID
}

// GetName ...
func (process Process) GetName() impl.STR_PTR {
	return &process.Name
}

// GetIsExecutable ...
func (process Process) GetIsExecutable() impl.BOOL_PTR {
	return &process.IsExecutable
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (process Process) GetDocumentation() *attributes.Documentation {
	return &process.Documentation[0]
}

// GetExtensionElements ...
func (process Process) GetExtensionElements() *attributes.ExtensionElements {
	return &process.ExtensionElements[0]
}

/*** Pool ***/

// GetLaneSet ...
func (process Process) GetLaneSet() *pool.LaneSet {
	return &process.LaneSet[0]
}

/*** Events ***/

// GetStartEvent ...
func (process Process) GetStartEvent(num int) *elements.StartEvent {
	return &process.StartEvent[num]
}

// GetBoundaryEvent ...
func (process Process) GetBoundaryEvent(num int) *elements.BoundaryEvent {
	return &process.BoundaryEvent[num]
}

// GetEndEvent ...
func (process Process) GetEndEvent(num int) *elements.EndEvent {
	return &process.EndEvent[num]
}

// GetIntermedCatchEvent ...
func (process Process) GetIntermediateCatchEvent(num int) *elements.IntermediateCatchEvent {
	return &process.IntermediateCatchEvent[num]
}

// GetIntermedThrowEvent ...
func (process Process) GetIntermediateThrowEvent(num int) *elements.IntermediateThrowEvent {
	return &process.IntermediateThrowEvent[num]
}

/*** Tasks ***/

// GetBusinessRuleTask ...
func (process Process) GetBusinessRuleTask(num int) *tasks.BusinessRuleTask {
	return &process.BusinessRuleTask[num]
}

// GetTask ...
func (process Process) GetTask(num int) *tasks.Task {
	return &process.Task[num]
}

// GetUserTask ...
func (process Process) GetUserTask(num int) *tasks.UserTask {
	return &process.UserTask[num]
}

// GetManualTask ...
func (process Process) GetManualTask(num int) *tasks.ManualTask {
	return &process.ManualTask[num]
}

// GetReceiveTask ...
func (process Process) GetReceiveTask(num int) *tasks.ReceiveTask {
	return &process.ReceiveTask[num]
}

// GetScriptTask ...
func (process Process) GetScriptTask(num int) *tasks.ScriptTask {
	return &process.ScriptTask[num]
}

// GetSendTask ...
func (process Process) GetSendTask(num int) *tasks.SendTask {
	return &process.SendTask[num]
}

// GetServiceTask ...
func (process Process) GetServiceTask(num int) *tasks.ServiceTask {
	return &process.ServiceTask[num]
}

// GetCallActivity ...
func (process Process) GetCallActivity(num int) *subprocesses.CallActivity {
	return &process.CallActivity[num]
}

// GetSubProcess ...
func (process Process) GetSubProcess(num int) *subprocesses.SubProcess {
	return &process.SubProcess[num]
}

// GetTransaction ...
func (process Process) GetTransaction(num int) *subprocesses.Transaction {
	return &process.Transaction[num]
}

// GetAdHocSubProcess ...
func (process Process) GetAdHocSubProcess(num int) *subprocesses.AdHocSubProcess {
	return &process.AdHocSubProcess[num]
}

/*** Gateways ***/

// GetExclusiveGateway
func (process Process) GetExclusiveGateway(num int) *gateways.ExclusiveGateway {
	return &process.ExclusiveGateway[num]
}

// GetInclsuiveGateway
func (process Process) GetInclusiveGateway(num int) *gateways.InclusiveGateway {
	return &process.InclusiveGateway[num]
}

// GetParallelGateway
func (process Process) GetParallelGateway(num int) *gateways.ParallelGateway {
	return &process.ParallelGateway[num]
}

// GetComplexGateway
func (process Process) GetComplexGateway(num int) *gateways.ComplexGateway {
	return &process.ComplexGateway[num]
}

// GetEventBasedGateway
func (process Process) GetEventBasedGateway(num int) *gateways.EventBasedGateway {
	return &process.EventBasedGateway[num]
}

/*** Marker ***/

// GetAssociation ...
func (process Process) GetAssociation(num int) *flow.Association {
	return &process.Association[num]
}

// GetSequenceFlow ...
func (process Process) GetSequenceFlow(num int) *flow.SequenceFlow {
	return &process.SequenceFlow[num]
}

/*** Data ***/

// GetDataObject ...
func (process Process) GetDataObject(num int) *data.DataObject {
	return &process.DataObject[num]
}
