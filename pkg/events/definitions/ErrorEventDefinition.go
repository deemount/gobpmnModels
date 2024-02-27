package definitions

import impl "github.com/deemount/gobpmnTypes"

// NewErrorEventDefinition ...
func NewErrorEventDefinition() ErrorEventDefinitionRepository {
	return &ErrorEventDefinition{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (eed *ErrorEventDefinition) SetID(typ string, suffix interface{}) {
	eed.ID = SetID(typ, suffix)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (eed ErrorEventDefinition) GetID() impl.STR_PTR {
	return &eed.ID
}
