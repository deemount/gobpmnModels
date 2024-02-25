package definitions

import (
	"github.com/deemount/gobpmnModels/pkg/conditional"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewConditionalEventDefinition ...
func NewConditionalEventDefinition() ConditionalEventDefinitionRepository {
	return &ConditionalEventDefinition{}
}

/*
 * Default Setters
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
 * Default Getters
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
