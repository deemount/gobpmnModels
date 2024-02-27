package elements

import (
	"fmt"

	impl "github.com/deemount/gobpmnTypes"
)

// NewSignal ...
func NewSignal() SignalRepository {
	return &Signal{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (signal *Signal) SetID(typ string, suffix interface{}) {
	signal.ID = SetID(typ, suffix)
}

// SetName ...
func (signal *Signal) SetName(suffix string) {
	signal.Name = fmt.Sprintf("Signal_%s", suffix)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (signal Signal) GetID() impl.STR_PTR {
	return &signal.ID
}

// GetName ...
func (signal Signal) GetName() impl.STR_PTR {
	return &signal.Name
}

/*
 * @String
 */

// String ...
func (signal Signal) String() string {
	return fmt.Sprintf("id=%v, name=%v", signal.ID, signal.Name)
}

// String ...
func (signal TSignal) String() string {
	return fmt.Sprintf("id=%v, name=%v", signal.ID, signal.Name)
}
