package conditional

import impl "github.com/deemount/gobpmnTypes"

// NewCondition ...
func NewCondition() ConditionRepository {
	return &Condition{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetConditionType ...
func (condition *Condition) SetConditionType() {
	condition.ConditionType = "bpmn:tFormalExpression"
}

// SetScriptFormat ...
func (condition *Condition) SetScriptFormat(format string) {
	condition.ScriptFormat = format
}

// SetScript ...
func (condition *Condition) SetScript(script string) {
	condition.Script = script
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetConditionType ...
func (condition Condition) GetConditionType() impl.STR_PTR {
	return &condition.ConditionType
}

// GetScriptFormat ...
func (condition Condition) GetScriptFormat() impl.STR_PTR {
	return &condition.ScriptFormat
}

// GetScript ...
func (condition Condition) GetScript() impl.STR_PTR {
	return &condition.Script
}
