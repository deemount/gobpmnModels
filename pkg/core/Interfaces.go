package core

import (
	"github.com/deemount/gobpmnModels/pkg/collaboration"
	"github.com/deemount/gobpmnModels/pkg/events"
	"github.com/deemount/gobpmnModels/pkg/marker"
	"github.com/deemount/gobpmnModels/pkg/process"
	impl "github.com/deemount/gobpmnTypes"
)

/*
 * @Base
 */

// DefinitionsElements ...
type DefinitionsElements interface {
	SetCollaboration()
	GetCollaboration() *collaboration.Collaboration
	SetProcess(num int)
	GetProcess(num int) *process.Process
	SetCategory(num int)
	GetCategory(num int) *marker.Category
	events.CoreEventsElementsRepository
}

/*
 * @Repositories
 */

// DefinitionsRepository ...
type DefinitionsRepository interface {
	impl.IFBaseID
	DefinitionsElements
	SetBpmn()
	SetDC()
	SetTargetNamespace()
	SetMainElements(num int)
	SetDefaultAttributes()
}
