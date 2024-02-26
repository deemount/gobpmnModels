package tasks

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

/*
 * @Base
 */

// TasksBaseAttributes ...
type TasksBaseAttributes interface {
	impl.IFBaseID
	impl.IFBaseName
}

// TasksMarkers ...
type TasksMarkers interface {
	SetIncoming(num int)
	GetIncoming(num int) marker.INCOMING_PTR
	SetOutgoing(num int)
	GetOutgoing(num int) marker.OUTGOING_PTR
}

// TasksBaseCoreElements ...
type TasksBaseCoreElements interface {
	attributes.AttributesBaseElements
}

// TasksBase ...
type TasksBase interface {
	TasksBaseAttributes
	TasksMarkers
	TasksBaseCoreElements
}

/*
 * @Repositories
 */

// TasksElementsRepository ...
type TasksElementsRepository interface {
	SetBusinessRuleTask(num int)
	GetBusinessRuleTask(num int) BUSINESS_RULE_TASK_PTR
	SetTask(num int)
	GetTask(num int) TASK_PTR
	SetUserTask(num int)
	GetUserTask(num int) USER_TASK_PTR
	SetManualTask(num int)
	GetManualTask(num int) MANUAL_TASK_PTR
	SetReceiveTask(num int)
	GetReceiveTask(num int) RECEIVE_TASK_PTR
	SetScriptTask(num int)
	GetScriptTask(num int) SCRIPT_TASK_PTR
	SetSendTask(num int)
	GetSendTask(num int) SEND_TASK_PTR
	SetServiceTask(num int)
	GetServiceTask(num int) SERVICE_TASK_PTR
}

// BusinessRuleTaskRepository ...
type BusinessRuleTaskRepository interface {
	TasksBase
	String() string
}

// ManualTaskRepository ...
type ManualTaskRepository interface {
	TasksBase
	String() string
}

// ReceiveTaskRepository ...
type ReceiveTaskRepository interface {
	TasksBase
	SetMessageRef(suffix string)
	GetMessageRef(suffix string) impl.STR_PTR
	String() string
}

// ScriptTaskRepository ...
type ScriptTaskRepository interface {
	TasksBase
	String() string
}

// SendTaskRepository ...
type SendTaskRepository interface {
	TasksBase
	String() string
}

// ServiceTaskRepository ...
type ServiceTaskRepository interface {
	TasksBase
	String() string
}

// TaskRepository ...
type TaskRepository interface {
	TasksBase
	SetDataInputAssociation(num int)
	GetDataInputAssociation(num int) flow.DATA_INPUT_ASSOCIATION_PTR
	String() string
}

// UserTaskRepository ...
type UserTaskRepository interface {
	TasksBase
	String() string
}
