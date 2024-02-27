package definitions

import (
	"github.com/deemount/gobpmnModels/pkg/conditional"
	impl "github.com/deemount/gobpmnTypes"
)

// NewConditionalEventDefinition ...
func NewConditionalEventDefinition() ConditionalEventDefinitionRepository {
	return &ConditionalEventDefinition{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (conditionalEventDefinition *ConditionalEventDefinition) SetID(typ string, suffix interface{}) {
	conditionalEventDefinition.ID = SetID(typ, suffix)
}

/* Elements */

/** BPMN **/

// SetCondition ...
func (conditionalEventDefinition *ConditionalEventDefinition) SetCondition() {
	conditionalEventDefinition.Condition = make([]conditional.Condition, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (conditionalEventDefinition ConditionalEventDefinition) GetID() impl.STR_PTR {
	return &conditionalEventDefinition.ID
}

/* Elements */

/** BPMN **/

// GetCondition ...
func (conditionalEventDefinition ConditionalEventDefinition) GetCondition() *conditional.Condition {
	return &conditionalEventDefinition.Condition[0]
}
