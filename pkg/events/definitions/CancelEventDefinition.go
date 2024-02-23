package definitions

import (
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewCancelEventDefinition ...
func NewCancelEventDefinition() CancelEventDefinitionRepository {
	return &CancelEventDefinition{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ced *CancelEventDefinition) SetID(typ string, suffix interface{}) {
	ced.ID = SetID(typ, suffix)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ced CancelEventDefinition) GetID() impl.STR_PTR {
	return &ced.ID
}
