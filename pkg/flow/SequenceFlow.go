package flow

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/conditional"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewSequenceFlow ...
func NewSequenceFlow() SequenceFlowRepository {
	return &SequenceFlow{}
}

/*
 * Default Setters
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
	switch typ {
	case "activity":
		sequenceFlow.SourceRef = fmt.Sprintf("Activity_%s", sourceRef)
		break
	case "event":
		sequenceFlow.SourceRef = fmt.Sprintf("Event_%s", sourceRef)
		break
	// Notice: using %v instead %s for more flexible parameter values
	case "startevent":
		sequenceFlow.SourceRef = fmt.Sprintf("StartEvent_%v", sourceRef)
		break
	case "id":
		sequenceFlow.SourceRef = fmt.Sprintf("%s", sourceRef)
		break
	}
}

// SetTargetRef ...
func (sequenceFlow *SequenceFlow) SetTargetRef(typ string, targetRef interface{}) {
	switch typ {
	case "activity":
		sequenceFlow.TargetRef = fmt.Sprintf("Activity_%s", targetRef)
		break
	case "event":
		sequenceFlow.TargetRef = fmt.Sprintf("Event_%s", targetRef)
		break
	case "startevent":
		sequenceFlow.TargetRef = fmt.Sprintf("StartEvent_%s", targetRef)
		break
	case "id":
		sequenceFlow.TargetRef = fmt.Sprintf("%s", targetRef)
		break
	}
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (sequenceFlow *SequenceFlow) SetDocumentation() {
	sequenceFlow.Documentation = make([]attributes.Documentation, 1)
}

/*** Conditional ***/

// SetConditionExpression ...
func (sequenceFlow *SequenceFlow) SetConditionExpression() {
	sequenceFlow.ConditionExpression = make([]conditional.ConditionExpression, 1)
}

/*
 * Default Getters
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
func (sequenceFlow SequenceFlow) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &sequenceFlow.Documentation[0]
}

/*** Conditional ***/

// GetConditionExpression ...
func (sequenceFlow SequenceFlow) GetConditionExpression() *conditional.ConditionExpression {
	return &sequenceFlow.ConditionExpression[0]
}
