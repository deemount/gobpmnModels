package data

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

// NewDataStoreReference ...
func NewDataStoreReference() DataStoreReferenceRepository {
	return &DataStoreReference{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (dsr *DataStoreReference) SetID(typ string, suffix interface{}) {
	dsr.ID = SetID(typ, suffix)
}

// SetName ...
func (dsr *DataStoreReference) SetName(name string) {
	dsr.Name = name
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (dsr *DataStoreReference) SetDocumentation() {
	dsr.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (dsr DataStoreReference) GetID() impl.STR_PTR {
	return &dsr.ID
}

// GetName ...
func (dsr DataStoreReference) GetName() impl.STR_PTR {
	return &dsr.Name
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (dsr DataStoreReference) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &dsr.Documentation[0]
}
