package definitions

import (
	"github.com/deemount/gobpmnModels/pkg/conditional"
	"github.com/deemount/gobpmnModels/pkg/time"
	impl "github.com/deemount/gobpmnTypes"
)

// Boundaries ...
type Boundaries struct {
	CancelEventDefinition      CANCEL_EVENT_DEF_SLC      `xml:"bpmn:cancelEventDefinition,omitempty" json:"cancelEventDefinition,omitempty"`
	MessageEventDefinition     MESSAGE_EVENT_DEF_SLC     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDefinition       TIMER_EVENT_DEF_SLC       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	EscalationEventDefinition  ESCALATION_EVENT_DEF_SLC  `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	ConditionalEventDefinition CONDITIONAL_EVENT_DEF_SLC `xml:"bpmn:conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	ErrorEventDefinition       ERROR_EVENT_DEF_SLC       `xml:"bpmn:errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition      SIGNAL_EVENT_DEF_SLC      `xml:"bpmn:signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	CompensateEventDefinition  COMPENSATE_EVENT_DEF_SLC  `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
}

// TBoundaries ...
type TBoundaries struct {
	CancelEventDefinition      CANCEL_EVENT_DEF_SLC       `xml:"cancelEventDefinition,omitempty" json:"cancelEventDefinition,omitempty"`
	MessageEventDefinition     MESSAGE_EVENT_DEF_SLC      `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDefinition       TTIMER_EVENT_DEF_SLC       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	EscalationEventDefinition  ESCALATION_EVENT_DEF_SLC   `xml:"escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	ConditionalEventDefinition TCONDITIONAL_EVENT_DEF_SLC `xml:"conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	ErrorEventDefinition       ERROR_EVENT_DEF_SLC        `xml:"errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition      SIGNAL_EVENT_DEF_SLC       `xml:"signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	CompensateEventDefinition  COMPENSATE_EVENT_DEF_SLC   `xml:"compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
}

// EndEvent ...
type EndEvent struct {
	CompensateEventDefinition COMPENSATE_EVENT_DEF_SLC `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	EscalationEventDefinition ESCALATION_EVENT_DEF_SLC `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    MESSAGE_EVENT_DEF_SLC    `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	ErrorEventDefinition      ERROR_EVENT_DEF_SLC      `xml:"bpmn:errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition     SIGNAL_EVENT_DEF_SLC     `xml:"bpmn:signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	TerminateEventDefinition  TERMINATE_EVENT_DEF_SLC  `xml:"bpmn:terminateEventDefinition,omitempty" json:"terminateEventDefinition,omitempty"`
}

// TEndEvent ...
type TEndEvent struct {
	CompensateEventDefinition TCOMPENSATE_EVENT_DEF_SLC `xml:"compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	EscalationEventDefinition ESCALATION_EVENT_DEF_SLC  `xml:"escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    MESSAGE_EVENT_DEF_SLC     `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	ErrorEventDefinition      ERROR_EVENT_DEF_SLC       `xml:"errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition     SIGNAL_EVENT_DEF_SLC      `xml:"signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	TerminateEventDefinition  TERMINATE_EVENT_DEF_SLC   `xml:"terminateEventDefinition,omitempty" json:"terminateEventDefinition,omitempty"`
}

// IntermediateCatchEvent ...
type IntermediateCatchEvent struct {
	MessageEventDefinition     MESSAGE_EVENT_DEF_SLC     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	LinkEventDefinition        LINK_EVENT_DEF_SLC        `xml:"bpmn:linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	ConditionalEventDefinition CONDITIONAL_EVENT_DEF_SLC `xml:"bpmn:conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	TimerEventDefinition       TIMER_EVENT_DEF_SLC       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
}

// TIntermediateCatchEvent ...
type TIntermediateCatchEvent struct {
	ConditionalEventDefinition TCONDITIONAL_EVENT_DEF_SLC `xml:"conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	LinkEventDefinition        LINK_EVENT_DEF_SLC         `xml:"linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	TimerEventDefinition       TTIMER_EVENT_DEF_SLC       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	MessageEventDefinition     MESSAGE_EVENT_DEF_SLC      `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
}

// IntermediateThrowEvent ...
type IntermediateThrowEvent struct {
	MessageEventDefinition    MESSAGE_EVENT_DEF_SLC    `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	LinkEventDefinition       LINK_EVENT_DEF_SLC       `xml:"bpmn:linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	CompensateEventDefinition COMPENSATE_EVENT_DEF_SLC `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	EscalationEventDefinition ESCALATION_EVENT_DEF_SLC `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	SignalEventDefinition     SIGNAL_EVENT_DEF_SLC     `xml:"bpmn:signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
}

// TIntermediateThrowEvent ...
type TIntermediateThrowEvent struct {
	MessageEventDefinition    MESSAGE_EVENT_DEF_SLC     `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	LinkEventDefinition       LINK_EVENT_DEF_SLC        `xml:"linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	CompensateEventDefinition TCOMPENSATE_EVENT_DEF_SLC `xml:"compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	EscalationEventDefinition ESCALATION_EVENT_DEF_SLC  `xml:"escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	SignalEventDefinition     SIGNAL_EVENT_DEF_SLC      `xml:"signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
}

// StartEvent ...
type StartEvent struct {
	ConditionalEventDef    CONDITIONAL_EVENT_DEF_SLC `xml:"bpmn:conditionalEventDefintion,omitempty" json:"conditionalEventDefinition,omitempty"`
	MessageEventDefinition MESSAGE_EVENT_DEF_SLC     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDef          TIMER_EVENT_DEF_SLC       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
}

// TStartEvent ...
type TStartEvent struct {
	ConditionalEventDef    TCONDITIONAL_EVENT_DEF_SLC `xml:"conditionalEventDefintion,omitempty" json:"conditionalEventDefinition,omitempty"`
	MessageEventDefinition MESSAGE_EVENT_DEF_SLC      `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDef          TTIMER_EVENT_DEF_SLC       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
}

/*
 * @Elementary
 */

// CancelEventDefinition ...
type CancelEventDefinition struct {
	impl.CoreID
}

// CompensateEventDefinition ...
type CompensateEventDefinition struct{}

// TCompensateEventDefinition ...
type TCompensateEventDefinition struct{}

// ConditionalEventDefinition ...
type ConditionalEventDefinition struct {
	impl.CoreID
	Condition []conditional.Condition `xml:"bpmn:condition,omitempty" json:"condition,omitempty"`
}

// TConditionalEventDefinition ...
type TConditionalEventDefinition struct {
	impl.CoreID
	Condition []conditional.Condition `xml:"condition,omitempty" json:"condition,omitempty"`
}

// ErrorEventDefinition ...
type ErrorEventDefinition struct {
	impl.CoreID
}

// EscalationEventDefinition ...
type EscalationEventDefinition struct {
	impl.CoreID
}

// LinkEventDefinition ...
type LinkEventDefinition struct {
	impl.CoreID
}

// MessageEventDefinition ...
type MessageEventDefinition struct {
	impl.CoreID
	MsgRef string `xml:"messageRef,attr,omitempty" json:"messageRef,omitempty"`
}

// SignalEventDefinition ...
type SignalEventDefinition struct {
	impl.CoreID
	SignalRef string `xml:"signalRef,attr,omitempty" json:"signalRef,omitempty"`
}

// TerminateEventDefinition ...
type TerminateEventDefinition struct {
	impl.CoreID
}

// TimerEventDefinition ...
type TimerEventDefinition struct {
	impl.CoreID
	TimeDate     time.TIME_DATE_SLC     `xml:"bpmn:timeDate,omitempty" json:"timeDate,omitempty"`
	TimeDuration time.TIME_DURATION_SLC `xml:"bpmn:timeDuration,omitempty" json:"timeDuration,omitempty"`
}

// TTimerEventDefinition ...
type TTimerEventDefinition struct {
	impl.CoreID
	TimeDate     time.TIME_DATE_SLC     `xml:"timeDate,omitempty" json:"timeDate,omitempty"`
	TimeDuration time.TIME_DURATION_SLC `xml:"timeDuration,omitempty" json:"timeDuration,omitempty"`
}
