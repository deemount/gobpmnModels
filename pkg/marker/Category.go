package marker

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewCategory ...
func NewCategory() CategoryRepository {
	return &Category{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (category *Category) SetID(typ string, suffix interface{}) {
	switch typ {
	case "category":
		category.ID = fmt.Sprintf("Category_%s", suffix)
		break
	case "id":
		category.ID = fmt.Sprintf("%s", suffix)
		break
	}
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
	category.Documentation = make([]attributes.Documentation, 1)
}

/**
 * Default Getters
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
