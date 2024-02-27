package definitions

import impl "github.com/deemount/gobpmnTypes"

// NewCancelEventDefinition ...
func NewCancelEventDefinition() CancelEventDefinitionRepository {
	return &CancelEventDefinition{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ced *CancelEventDefinition) SetID(typ string, suffix interface{}) {
	ced.ID = SetID(typ, suffix)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ced CancelEventDefinition) GetID() impl.STR_PTR {
	return &ced.ID
}
