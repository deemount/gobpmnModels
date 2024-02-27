package definitions

import impl "github.com/deemount/gobpmnTypes"

// NewTerminateEventDefinition ...
func NewTerminateEventDefinition() TerminateEventDefinitionRepository {
	return &TerminateEventDefinition{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ted *TerminateEventDefinition) SetID(typ string, suffix interface{}) {
	ted.ID = SetID(typ, suffix)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ted TerminateEventDefinition) GetID() impl.STR_PTR {
	return &ted.ID
}
