package definitions

import impl "github.com/deemount/gobpmnTypes"

// NewLinkEventDefinition ...
func NewLinkEventDefinition() LinkEventDefinitionRepository {
	return &LinkEventDefinition{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (led *LinkEventDefinition) SetID(typ string, suffix interface{}) {
	led.ID = SetID(typ, suffix)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (led LinkEventDefinition) GetID() impl.STR_PTR {
	return &led.ID
}
