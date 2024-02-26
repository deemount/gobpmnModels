package flow

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
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
	dia.Documentation = make([]attributes.Documentation, 1)
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
func (dia DataInputAssociation) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &dia.Documentation[0]
}
