package tasks

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
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
	GetIncoming(num int) *marker.Incoming
	SetOutgoing(num int)
	GetOutgoing(num int) *marker.Outgoing
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
	GetBusinessRuleTask(num int) *BusinessRuleTask
	SetTask(num int)
	GetTask(num int) *Task
	SetUserTask(num int)
	GetUserTask(num int) *UserTask
	SetManualTask(num int)
	GetManualTask(num int) *ManualTask
	SetReceiveTask(num int)
	GetReceiveTask(num int) *ReceiveTask
	SetScriptTask(num int)
	GetScriptTask(num int) *ScriptTask
	SetSendTask(num int)
	GetSendTask(num int) *SendTask
	SetServiceTask(num int)
	GetServiceTask(num int) *ServiceTask
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
	GetDataInputAssociation(num int) *flow.DataInputAssociation
	String() string
}

// UserTaskRepository ...
type UserTaskRepository interface {
	TasksBase
	String() string
}
