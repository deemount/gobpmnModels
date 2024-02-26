package process

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/data"
	"github.com/deemount/gobpmnModels/pkg/events"
	"github.com/deemount/gobpmnModels/pkg/events/elements"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/gateways"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/pool"
	"github.com/deemount/gobpmnModels/pkg/subprocesses"
	"github.com/deemount/gobpmnModels/pkg/tasks"
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

/** Pool **/

// SetLaneSet ...
func (process *Process) SetLaneSet() {
	process.LaneSet = make([]pool.LaneSet, 1)
}

/*** Events ***/

// SetStartEvent ...
func (process *Process) SetStartEvent(num int) {
	process.StartEvent = make([]elements.StartEvent, num)
}

// SetBoundaryEvent ...
func (process *Process) SetBoundaryEvent(num int) {
	process.BoundaryEvent = make([]elements.BoundaryEvent, num)
}

// SetEndEvent ...
func (process *Process) SetEndEvent(num int) {
	process.EndEvent = make(events.END_EVENT_SLC, num)
}

// SetIntermedCatchEvent ...
func (process *Process) SetIntermediateCatchEvent(num int) {
	process.IntermediateCatchEvent = make([]elements.IntermediateCatchEvent, num)
}

// SetIntermedThrowEvent ...
func (process *Process) SetIntermediateThrowEvent(num int) {
	process.IntermediateThrowEvent = make([]elements.IntermediateThrowEvent, num)
}

/*** Tasks ***/

// SetBusinessRuleTask ...
func (process *Process) SetBusinessRuleTask(num int) {
	process.BusinessRuleTask = make([]tasks.BusinessRuleTask, num)
}

// SetTask ...
func (process *Process) SetTask(num int) {
	process.Task = make([]tasks.Task, num)
}

// SetUserTask ...
func (process *Process) SetUserTask(num int) {
	process.UserTask = make([]tasks.UserTask, num)
}

// SetManualTask ...
func (process *Process) SetManualTask(num int) {
	process.ManualTask = make([]tasks.ManualTask, num)
}

// SetReceiveTask ...
func (process *Process) SetReceiveTask(num int) {
	process.ReceiveTask = make([]tasks.ReceiveTask, num)
}

// SetScriptTask ...
func (process *Process) SetScriptTask(num int) {
	process.ScriptTask = make([]tasks.ScriptTask, num)
}

// SetSendTask ...
func (process *Process) SetSendTask(num int) {
	process.SendTask = make([]tasks.SendTask, num)
}

// SetServiceTask ...
func (process *Process) SetServiceTask(num int) {
	process.ServiceTask = make([]tasks.ServiceTask, num)
}

/*** Subprocesses ***/

// SetCallActivity ...
func (process *Process) SetCallActivity(num int) {
	process.CallActivity = make([]subprocesses.CallActivity, num)
}

// SetSubProcess ...
func (process *Process) SetSubProcess(num int) {
	process.SubProcess = make([]subprocesses.SubProcess, num)
}

// SetTransaction ...
func (process *Process) SetTransaction(num int) {
	process.Transaction = make([]subprocesses.Transaction, num)
}

// SetAdHocSubProcess ...
func (process *Process) SetAdHocSubProcess(num int) {
	process.AdHocSubProcess = make([]subprocesses.AdHocSubProcess, num)
}

/*** Gateways ***/

// SetExclusiveGateway
func (process *Process) SetExclusiveGateway(num int) {
	process.ExclusiveGateway = make([]gateways.ExclusiveGateway, num)
}

// SetInclsuiveGateway
func (process *Process) SetInclusiveGateway(num int) {
	process.InclusiveGateway = make([]gateways.InclusiveGateway, num)
}

// SetParallelGateway
func (process *Process) SetParallelGateway(num int) {
	process.ParallelGateway = make([]gateways.ParallelGateway, num)
}

// SetComplexGateway
func (process *Process) SetComplexGateway(num int) {
	process.ComplexGateway = make([]gateways.ComplexGateway, num)
}

// SetEventBasedGateway
func (process *Process) SetEventBasedGateway(num int) {
	process.EventBasedGateway = make([]gateways.EventBasedGateway, num)
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
func (process Process) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &process.Documentation[0]
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
func (process Process) GetEndEvent(num int) events.END_EVENT_PTR {
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
func (process Process) GetBusinessRuleTask(num int) tasks.BUSINESS_RULE_TASK_PTR {
	return &process.BusinessRuleTask[num]
}

// GetTask ...
func (process Process) GetTask(num int) tasks.TASK_PTR {
	return &process.Task[num]
}

// GetUserTask ...
func (process Process) GetUserTask(num int) tasks.USER_TASK_PTR {
	return &process.UserTask[num]
}

// GetManualTask ...
func (process Process) GetManualTask(num int) tasks.MANUAL_TASK_PTR {
	return &process.ManualTask[num]
}

// GetReceiveTask ...
func (process Process) GetReceiveTask(num int) tasks.RECEIVE_TASK_PTR {
	return &process.ReceiveTask[num]
}

// GetScriptTask ...
func (process Process) GetScriptTask(num int) tasks.SCRIPT_TASK_PTR {
	return &process.ScriptTask[num]
}

// GetSendTask ...
func (process Process) GetSendTask(num int) tasks.SEND_TASK_PTR {
	return &process.SendTask[num]
}

// GetServiceTask ...
func (process Process) GetServiceTask(num int) tasks.SERVICE_TASK_PTR {
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
func (process Process) GetExclusiveGateway(num int) gateways.EXCLUSIVE_GATEWAY_PTR {
	return &process.ExclusiveGateway[num]
}

// GetInclsuiveGateway
func (process Process) GetInclusiveGateway(num int) gateways.INCLUSIVE_GATEWAY_PTR {
	return &process.InclusiveGateway[num]
}

// GetParallelGateway
func (process Process) GetParallelGateway(num int) gateways.PARALLEL_GATEWAY_PTR {
	return &process.ParallelGateway[num]
}

// GetComplexGateway
func (process Process) GetComplexGateway(num int) gateways.COMPLEX_GATEWAY_PTR {
	return &process.ComplexGateway[num]
}

// GetEventBasedGateway
func (process Process) GetEventBasedGateway(num int) gateways.EVENT_BASED_GATEWAYS_PTR {
	return &process.EventBasedGateway[num]
}

/*** Marker ***/

// GetAssociation ...
func (process Process) GetAssociation(num int) flow.ASSOCIATION_PTR {
	return &process.Association[num]
}

// GetSequenceFlow ...
func (process Process) GetSequenceFlow(num int) flow.SEQUENCE_FLOW_PTR {
	return &process.SequenceFlow[num]
}

/*** Data ***/

// GetDataObject ...
func (process Process) GetDataObject(num int) *data.DataObject {
	return &process.DataObject[num]
}
