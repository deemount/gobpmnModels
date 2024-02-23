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
 * Default Setters
 */

/* Attributes */

// SetID ...
func (dia *DataInputAssociation) SetID(typ string, suffix interface{}) {
	dia.ID = SetID(typ, suffix)
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (dia *DataInputAssociation) SetDocumentation() {
	dia.Documentation = make([]attributes.Documentation, 1)
}

/*
 * Default Getters
 */

/* Attributes */

// GetID ...
func (dia DataInputAssociation) GetID() impl.STR_PTR {
	return &dia.ID
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (dia DataInputAssociation) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &dia.Documentation[0]
}
