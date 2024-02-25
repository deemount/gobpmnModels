package data

import (
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewDataObject ...
func NewDataObject() DataObjectRepository {
	return &DataObject{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (do *DataObject) SetID(typ string, suffix interface{}) {
	do.ID = SetID(typ, suffix)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (do DataObject) GetID() impl.STR_PTR {
	return &do.ID
}
