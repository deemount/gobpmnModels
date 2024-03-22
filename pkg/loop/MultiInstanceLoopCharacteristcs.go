package loop

import (
	"github.com/deemount/gobpmnModels/pkg/conditional"
	impl "github.com/deemount/gobpmnTypes"
)

// NewMultiInstanceLoopCharacteristics ...
func NewMultiInstanceLoopCharacteristics() MultiInstanceLoopCharacteristicsRepository {
	return &MultiInstanceLoopCharacteristics{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetIsSequential ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetIsSequential(is bool) {
	multiInstanceLoopCharacteristics.IsSequential = is
}

/* Elements */

/** BPMN **/

// SetLoopCardinality ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetLoopCardinality() {
	multiInstanceLoopCharacteristics.LoopCardinality = make(LOOP_CARDINALITY_SLICE, 1)
}

/*** Conditional ***/

// SetCompletionCondition ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetCompletionCondition() {
	multiInstanceLoopCharacteristics.CompletionCondition = make([]conditional.CompletionCondition, 1)
}

/*
 * @Getters
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
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetLoopCardinality() *LoopCardinality {
	return &multiInstanceLoopCharacteristics.LoopCardinality[0]
}

/*** Conditional ***/

// GetCompletionCondition ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCompletionCondition() *conditional.CompletionCondition {
	return &multiInstanceLoopCharacteristics.CompletionCondition[0]
}
