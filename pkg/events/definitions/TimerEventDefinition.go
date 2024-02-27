package definitions

import (
	"github.com/deemount/gobpmnModels/pkg/time"
	impl "github.com/deemount/gobpmnTypes"
)

// NewTimerEventDefinition ...
func NewTimerEventDefinition() TimerEventDefinitionRepository {
	return &TimerEventDefinition{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ted *TimerEventDefinition) SetID(typ string, suffix interface{}) {
	ted.ID = SetID(typ, suffix)
}

/* Elements */

/** BPMN **/

// SetTimeDuration ...
func (ted *TimerEventDefinition) SetTimeDate() {
	ted.TimeDate = make([]time.TimeDate, 1)
}

// SetTimeDuration ...
func (ted *TimerEventDefinition) SetTimeDuration() {
	ted.TimeDuration = make([]time.TimeDuration, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ted TimerEventDefinition) GetID() impl.STR_PTR {
	return &ted.ID
}

/* Elements */

/** BPMN **/

// GetTimeDuration ...
func (ted TimerEventDefinition) GetTimeDate() *time.TimeDate {
	return &ted.TimeDate[0]
}

// GetTimeDuration ...
func (ted TimerEventDefinition) GetTimeDuration() *time.TimeDuration {
	return &ted.TimeDuration[0]
}
