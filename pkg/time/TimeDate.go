package time

import (
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewTimeDate ...
func NewTimeDate() TimeDateRepository {
	return &TimeDate{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetTimerDefinitionType ...
func (timedate *TimeDate) SetTimerDefinitionType() {
	timedate.TimerDefType = "bpmn:tFormalExpression"
}

// SetTimerDefinition ...
func (timedate *TimeDate) SetTimerDefinition(timerDefinition string) {
	timedate.TimerDef = timerDefinition
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetTimerDefinitionType ...
func (timedate TimeDate) GetTimerDefinitionType() impl.STR_PTR {
	return &timedate.TimerDefType
}

// GetTimerDefinition ...
func (timedate TimeDate) GetTimerDefinition() impl.STR_PTR {
	return &timedate.TimerDef
}
