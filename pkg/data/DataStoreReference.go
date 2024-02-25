package data

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewDataStoreReference ...
func NewDataStoreReference() DataStoreReferenceRepository {
	return &DataStoreReference{}
}

/*
 * Default Setters
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
	dsr.Documentation = make([]attributes.Documentation, 1)
}

/*
 * Default Getters
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
