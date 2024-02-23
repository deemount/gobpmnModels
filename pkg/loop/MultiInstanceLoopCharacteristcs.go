package loop

import (
	"github.com/deemount/gobpmnModels/pkg/conditional"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewMultiInstanceLoopCharacteristics ...
func NewMultiInstanceLoopCharacteristics() MultiInstanceLoopCharacteristicsRepository {
	return &MultiInstanceLoopCharacteristics{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetIsSequential ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetIsSequential(is bool) {
	multiInstanceLoopCharacteristics.IsSequential = is
}

/*** Make Elements ***/

/** BPMN **/

// SetLoopCardinality ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetLoopCardinality() {
	multiInstanceLoopCharacteristics.LoopCardinality = make([]LoopCardinality, 1)
}

// SetCompletionCondition ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetCompletionCondition() {
	multiInstanceLoopCharacteristics.CompletionCondition = make([]conditional.CompletionCondition, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetIsSequential ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetIsSequential() impl.BOOL_PTR {
	return &multiInstanceLoopCharacteristics.IsSequential
}

/* Elements */

/** BPMN **/

// GetLoopCardinality ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetLoopCardinality() LOOP_CARDINALITY_PTR {
	return &multiInstanceLoopCharacteristics.LoopCardinality[0]
}

// GetCompletionCondition ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCompletionCondition() *conditional.CompletionCondition {
	return &multiInstanceLoopCharacteristics.CompletionCondition[0]
}
