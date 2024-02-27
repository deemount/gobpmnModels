package marker

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

// NewCategory ...
func NewCategory() CategoryRepository {
	return &Category{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (category *Category) SetID(typ string, suffix interface{}) {
	category.ID = SetID(typ, suffix)
}

/* Elements */

/** BPMN **/

// SetCategoryValue...
func (category *Category) SetCategoryValue() {
	category.CategoryValue = make(CATEGORY_VALUE_SLC, 1)
}

/*** Attributes ***/

// SetDocumentation ...
func (category *Category) SetDocumentation() {
	category.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (category Category) GetID() impl.STR_PTR {
	return &category.ID
}

/* Elements */

/** BPMN **/

// SetCategoryValue...
func (category Category) GetCategoryValue() CATEGORY_VALUE_PTR {
	return &category.CategoryValue[0]
}

/*** Attributes ***/

// GetDocumentation ...
func (category Category) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &category.Documentation[0]
}
