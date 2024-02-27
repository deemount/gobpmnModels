package time

import impl "github.com/deemount/gobpmnTypes"

// NewTimeCycle ...
func NewTimeCycle() TimeCycleRepository {
	return &TimeCycle{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetTimerDefinitionType ...
func (timecycle *TimeCycle) SetTimerDefinitionType() {
	timecycle.TimerDefType = "bpmn:tFormalExpression"
}

// SetTimerDefinition ...
func (timecycle *TimeCycle) SetTimerDefinition(timerDefinition string) {
	timecycle.TimerDef = timerDefinition
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetTimerDefinitionType ...
func (timecycle TimeCycle) GetTimerDefinitionType() impl.STR_PTR {
	return &timecycle.TimerDefType
}

// GetTimerDefinition ...
func (timecycle TimeCycle) GetTimerDefinition() impl.STR_PTR {
	return &timecycle.TimerDef
}
