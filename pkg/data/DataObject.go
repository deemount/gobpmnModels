package data

import impl "github.com/deemount/gobpmnTypes"

// NewDataObject ...
func NewDataObject() DataObjectRepository {
	return &DataObject{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (do *DataObject) SetID(typ string, suffix interface{}) {
	do.ID = SetID(typ, suffix)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (do DataObject) GetID() impl.STR_PTR {
	return &do.ID
}
