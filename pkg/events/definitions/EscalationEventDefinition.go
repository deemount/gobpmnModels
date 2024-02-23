package definitions

import (
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewEscalationEventDefinition
func NewEscalationEventDefinition() EscalationEventDefinitionRepository {
	return &EscalationEventDefinition{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (esced *EscalationEventDefinition) SetID(typ string, suffix interface{}) {
	esced.ID = SetID(typ, suffix)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (esced EscalationEventDefinition) GetID() impl.STR_PTR {
	return &esced.ID
}
