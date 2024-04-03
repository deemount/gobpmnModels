package flow

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

// NewDataInputAssociation ...
func NewDataInputAssociation() DataInputAssociationRepository {
	return &DataInputAssociation{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (dia *DataInputAssociation) SetID(typ string, suffix interface{}) {
	dia.ID = SetID(typ, suffix)
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (dia *DataInputAssociation) SetDocumentation() {
	dia.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

// SetExtensionElements ...
func (dia *DataInputAssociation) SetExtensionElements() {
	dia.ExtensionElements = make(attributes.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (dia DataInputAssociation) GetID() impl.STR_PTR {
	return &dia.ID
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (dia DataInputAssociation) GetDocumentation() *attributes.Documentation {
	return &dia.Documentation[0]
}

// GetExtensionElements ...
func (dia DataInputAssociation) GetExtensionElements() *attributes.ExtensionElements {
	return &dia.ExtensionElements[0]
}
