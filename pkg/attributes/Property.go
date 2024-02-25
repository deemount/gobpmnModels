package attributes

import (
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewProperty ...
func NewProperty() PropertyRepository {
	return &Property{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (prop *Property) SetID(typ string, suffix interface{}) {
	prop.ID = SetID(typ, suffix)
}

// SetName ...
// Notice: maybe put a placeholder like name="__targetRef_placeholder" in it
func (prop *Property) SetName(name string) {
	prop.Name = name
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (prop Property) GetID() impl.STR_PTR {
	return &prop.ID

}

// GetName ...
func (prop Property) GetName() impl.STR_PTR {
	return &prop.Name
}
