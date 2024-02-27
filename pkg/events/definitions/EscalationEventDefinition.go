package definitions

import impl "github.com/deemount/gobpmnTypes"

// NewEscalationEventDefinition
func NewEscalationEventDefinition() EscalationEventDefinitionRepository {
	return &EscalationEventDefinition{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (esced *EscalationEventDefinition) SetID(typ string, suffix interface{}) {
	esced.ID = SetID(typ, suffix)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (esced EscalationEventDefinition) GetID() impl.STR_PTR {
	return &esced.ID
}
