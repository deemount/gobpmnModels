package flow

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/conditional"
	impl "github.com/deemount/gobpmnTypes"
)

// NewSequenceFlow ...
func NewSequenceFlow() SequenceFlowRepository {
	return &SequenceFlow{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (sequenceFlow *SequenceFlow) SetID(typ string, suffix interface{}) {
	sequenceFlow.ID = SetID(typ, suffix)
}

// SetName ...
func (sequenceFlow *SequenceFlow) SetName(name string) {
	sequenceFlow.Name = name
}

// SetSourceRef ...
func (sequenceFlow *SequenceFlow) SetSourceRef(typ string, sourceRef interface{}) {
	sequenceFlow.SourceRef = SetSourceTargetRef(typ, sourceRef)
}

// SetTargetRef ...
func (sequenceFlow *SequenceFlow) SetTargetRef(typ string, targetRef interface{}) {
	sequenceFlow.TargetRef = SetSourceTargetRef(typ, targetRef)
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (sequenceFlow *SequenceFlow) SetDocumentation() {
	sequenceFlow.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

// SetExtensionElements ...
func (sequenceFlow *SequenceFlow) SetExtensionElements() {
	sequenceFlow.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*** Conditional ***/

// SetConditionExpression ...
func (sequenceFlow *SequenceFlow) SetConditionExpression() {
	sequenceFlow.ConditionExpression = make([]conditional.ConditionExpression, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (sequenceFlow SequenceFlow) GetID() impl.STR_PTR {
	return &sequenceFlow.ID
}

// GetName ...
func (sequenceFlow SequenceFlow) GetName() impl.STR_PTR {
	return &sequenceFlow.Name
}

// GetSourceRef ...
func (sequenceFlow SequenceFlow) GetSourceRef() impl.STR_PTR {
	return &sequenceFlow.SourceRef
}

// GetTargetRef ...
func (sequenceFlow SequenceFlow) GetTargetRef() impl.STR_PTR {
	return &sequenceFlow.TargetRef
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (sequenceFlow SequenceFlow) GetDocumentation() *attributes.Documentation {
	return &sequenceFlow.Documentation[0]
}

/*** Conditional ***/

// GetConditionExpression ...
func (sequenceFlow SequenceFlow) GetConditionExpression() *conditional.ConditionExpression {
	return &sequenceFlow.ConditionExpression[0]
}

// GetExtensionElements ...
func (sequenceFlow SequenceFlow) GetExtensionElements() *attributes.ExtensionElements {
	return &sequenceFlow.ExtensionElements[0]
}
