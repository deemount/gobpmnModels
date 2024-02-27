package definitions

import (
	"fmt"

	impl "github.com/deemount/gobpmnTypes"
)

// NewSignalEventDefinition ...
func NewSignalEventDefinition() SignalEventDefinitionRepository {
	return &SignalEventDefinition{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (sed *SignalEventDefinition) SetID(typ string, suffix interface{}) {
	sed.ID = SetID(typ, suffix)
}

// SetSignalRef ...
func (sed *SignalEventDefinition) SetSignalRef(suffix string) {
	sed.SignalRef = fmt.Sprintf("Signal_%s", suffix)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (sed SignalEventDefinition) GetID() impl.STR_PTR {
	return &sed.ID
}

// GetSignalRef ...
func (sed SignalEventDefinition) GetSignalRef() impl.STR_PTR {
	return &sed.SignalRef
}
