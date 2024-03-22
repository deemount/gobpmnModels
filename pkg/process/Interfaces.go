package process

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/data"
	"github.com/deemount/gobpmnModels/pkg/events"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/gateways"
	"github.com/deemount/gobpmnModels/pkg/pool"
	"github.com/deemount/gobpmnModels/pkg/subprocesses"
	"github.com/deemount/gobpmnModels/pkg/tasks"
	impl "github.com/deemount/gobpmnTypes"
)

/*
 * @Repositories
 */

// ProcessRepository ...
type ProcessRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	attributes.AttributesBaseElements
	events.ProcessEventsElementsRepository
	tasks.TasksElementsRepository
	gateways.GatewaysElementsRepository
	subprocesses.SubprocessesElementsRepository
	flow.FlowSequenceFlow

	SetIsExecutable(isExec bool)
	GetIsExecutable() impl.BOOL_PTR

	SetLaneSet()
	GetLaneSet() *pool.LaneSet

	SetDataObject(num int)
	GetDataObject(num int) *data.DataObject
}
