package flow

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

// NewMessageFlow ...
func NewMessageFlow() MessageFlowRepository {
	return &MessageFlow{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (messageFlow *MessageFlow) SetID(typ string, suffix interface{}) {
	messageFlow.ID = SetID(typ, suffix)
}

// SetName ...
func (messageFlow *MessageFlow) SetName(name string) {
	messageFlow.Name = name
}

// SetSourceRef ...
func (messageFlow *MessageFlow) SetSourceRef(typ string, sourceRef interface{}) {
	messageFlow.SourceRef = SetSourceTargetRef(typ, sourceRef)
}

// SetTargetRef ...
func (messageFlow *MessageFlow) SetTargetRef(typ string, targetRef interface{}) {
	messageFlow.TargetRef = SetSourceTargetRef(typ, targetRef)
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (messageFlow *MessageFlow) SetDocumentation() {
	messageFlow.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

// SetExtensionElements ...
func (messageFlow *MessageFlow) SetExtensionElements() {
	messageFlow.ExtensionElements = make(attributes.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (messageFlow MessageFlow) GetID() impl.STR_PTR {
	return &messageFlow.ID
}

// GetName ...
func (messageFlow MessageFlow) GetName() impl.STR_PTR {
	return &messageFlow.Name
}

// GetSourceRef ...
func (messageFlow MessageFlow) GetSourceRef() impl.STR_PTR {
	return &messageFlow.SourceRef
}

// GetTargetRef ...
func (messageFlow MessageFlow) GetTargetRef() impl.STR_PTR {
	return &messageFlow.TargetRef
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (messageFlow MessageFlow) GetDocumentation() *attributes.Documentation {
	return &messageFlow.Documentation[0]
}

// GetExtensionElements ...
func (messageFlow MessageFlow) GetExtensionElements() *attributes.ExtensionElements {
	return &messageFlow.ExtensionElements[0]
}
