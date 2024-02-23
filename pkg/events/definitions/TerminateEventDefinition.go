package definitions

import (
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewTerminateEventDefinition ...
func NewTerminateEventDefinition() TerminateEventDefinitionRepository {
	return &TerminateEventDefinition{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ted *TerminateEventDefinition) SetID(typ string, suffix interface{}) {
	ted.ID = SetID(typ, suffix)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ted TerminateEventDefinition) GetID() impl.STR_PTR {
	return &ted.ID
}
