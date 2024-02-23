package elements

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewSignal ...
func NewSignal() SignalRepository {
	return &Signal{}
}

/*
 * Default Setters
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
 * Default Getters
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
 * Default String
 */

// String ...
func (signal Signal) String() string {
	return fmt.Sprintf("id=%v, name=%v", signal.ID, signal.Name)
}

// String ...
func (signal TSignal) String() string {
	return fmt.Sprintf("id=%v, name=%v", signal.ID, signal.Name)
}
