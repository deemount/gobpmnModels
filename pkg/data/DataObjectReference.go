package data

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewDataObjectReference ...
func NewDataObjectReference() DataObjectReferenceRepository {
	return &DataObjectReference{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (dor *DataObjectReference) SetID(typ string, suffix interface{}) {
	dor.ID = SetID(typ, suffix)
}

// SetName ...
func (dor *DataObjectReference) SetName(name string) {
	dor.Name = name
}

// SetDataObjectRef ...
func (dor *DataObjectReference) SetDataObjectRef(suffix interface{}) {
	dor.Name = fmt.Sprintf("DataObject_%v", suffix)
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (dor *DataObjectReference) SetDocumentation() {
	dor.Documentation = make([]attributes.Documentation, 1)
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (dor DataObjectReference) GetID() impl.STR_PTR {
	return &dor.ID
}

// SetName ...
func (dor DataObjectReference) GetName() impl.STR_PTR {
	return &dor.Name
}

// SetDataObjectRef ...
func (dor DataObjectReference) GetDataObjectRef() impl.STR_PTR {
	return &dor.Name
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (dor DataObjectReference) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &dor.Documentation[0]
}
