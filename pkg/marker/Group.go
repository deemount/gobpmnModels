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
 * @Setters
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

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (group *Group) SetDocumentation() {
	group.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (group Group) GetID() impl.STR_PTR {
	return &group.ID

}

// GetCategoryValRef ...
func (group Group) GetCategoryValueRef() impl.STR_PTR {
	return &group.CategoryValueRef
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (group Group) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &group.Documentation[0]
}
