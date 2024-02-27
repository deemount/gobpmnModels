package time

import impl "github.com/deemount/gobpmnTypes"

/*
 * @Base
 */

// TimeBaseDefintionType ...
type TimeBaseDefintionType interface {
	SetTimerDefinitionType()
	GetTimerDefinitionType() impl.STR_PTR
}

// TimeBaseDefinition ...
type TimeBaseDefinition interface {
	SetTimerDefinition(timerDefinition string)
	GetTimerDefinition() impl.STR_PTR
}

// TimeBaseCoreElements ...
type TimeBaseCoreElements interface {
	TimeBaseDefintionType
	TimeBaseDefinition
}

// TimeBase ...
type TimeBase interface {
	TimeBaseCoreElements
}

/*
 * @Repositories
 */

// TimeCycle ...
type TimeCycleRepository interface {
	TimeBase
}

// TimeDateRepository ...
type TimeDateRepository interface {
	TimeBase
}

// TimeDurationRepository ...
type TimeDurationRepository interface {
	TimeBase
}
