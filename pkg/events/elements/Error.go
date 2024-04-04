package elements

import impl "github.com/deemount/gobpmnTypes"

// NewError ...
func NewError() Error {
	return Error{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (error *Error) SetID(id string) {
	error.ID = id
}

// SetName ...
func (error *Error) SetName(name string) {
	error.Name = name
}

// SetErrorCode ...
func (error *Error) SetErrorCode(errorCode string) {
	error.ErrorCode = errorCode
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (error Error) GetID() impl.STR_PTR {
	return &error.ID
}

// GetName ...
func (error Error) GetName() impl.STR_PTR {
	return &error.Name
}

// GetErrorCode ...
func (error Error) GetErrorCode() impl.STR_PTR {
	return &error.ErrorCode
}
