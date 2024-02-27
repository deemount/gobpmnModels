package time

import impl "github.com/deemount/gobpmnTypes"

// NewTimeDuration ...
func NewTimeDuration() TimeDurationRepository {
	return &TimeDuration{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetTimerDefinitionType ...
func (timeduration *TimeDuration) SetTimerDefinitionType() {
	timeduration.TimerDefType = "bpmn:tFormalExpression"
}

// SetTimerDefinition ...
func (timeduration *TimeDuration) SetTimerDefinition(timerDefinition string) {
	timeduration.TimerDef = timerDefinition
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetTimerDefinitionType ...
func (timeduration TimeDuration) GetTimerDefinitionType() impl.STR_PTR {
	return &timeduration.TimerDefType
}

// GetTimerDefinition ...
func (timeduration TimeDuration) GetTimerDefinition() impl.STR_PTR {
	return &timeduration.TimerDef
}
