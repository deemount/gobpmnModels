package marker

import "github.com/deemount/gobpmnModels/pkg/impl"

// NewCategoryValue ...
func NewCategoryValue() CategoryValueRepository {
	return &CategoryValue{}
}

/*
 * @Setters
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
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID
func (categoryValue CategoryValue) GetID() impl.STR_PTR {
	return &categoryValue.ID
}

// GetValue
func (categoryValue CategoryValue) GetValue() impl.STR_PTR {
	return &categoryValue.Value
}
