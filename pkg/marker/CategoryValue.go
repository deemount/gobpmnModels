package marker

import (
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewCategoryValue ...
func NewCategoryValue() CategoryValueRepository {
	return &CategoryValue{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID
func (categoryValue *CategoryValue) SetID(typ string, suffix interface{}) {
	categoryValue.ID = SetID(typ, suffix)
}

// SetValue
func (categoryValue *CategoryValue) SetValue(value string) {
	categoryValue.Value = value
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID
func (categoryValue CategoryValue) GetID() impl.STR_PTR {
	return &categoryValue.ID
}

// GetValue
func (categoryValue CategoryValue) GetValue() *string {
	return &categoryValue.Value
}
