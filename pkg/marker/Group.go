package marker

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewGroup ...
func NewGroup() GroupRepository {
	return &Group{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (group *Group) SetID(typ string, suffix interface{}) {
	group.ID = SetID(typ, suffix)
}

// SetCategoryValRef ...
func (group *Group) SetCategoryValueRef(suffix string) {
	group.CategoryValueRef = fmt.Sprintf("_%s", suffix)
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (group *Group) SetDocumentation() {
	group.Documentation = make([]attributes.Documentation, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (group Group) GetID() impl.STR_PTR {
	return &group.ID

}

// GetCategoryValRef ...
func (group Group) GetCategoryValueRef() *string {
	return &group.CategoryValueRef
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (group Group) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &group.Documentation[0]
}