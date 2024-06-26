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

// SetExtensionElements ...
func (category *Category) SetExtensionElements() {
	category.ExtensionElements = make(attributes.EXTENSION_ELEMENTS_SLC, 1)
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
func (category Category) GetCategoryValue() *CategoryValue {
	return &category.CategoryValue[0]
}

/*** Attributes ***/

// GetDocumentation ...
func (category Category) GetDocumentation() *attributes.Documentation {
	return &category.Documentation[0]
}

// GetExtensionElements ...
func (category Category) GetExtensionElements() *attributes.ExtensionElements {
	return &category.ExtensionElements[0]
}
