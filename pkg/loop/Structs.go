package loop

import (
	"github.com/deemount/gobpmnModels/pkg/conditional"
)

/*
 * @Elementary
 */

// LoopCardinality ...
type LoopCardinality struct{}

// TLoopCardinality ...
type TLoopCardinality struct{}

// MultiInstanceLoopCharacteristics ...
type MultiInstanceLoopCharacteristics struct {
	IsSequential        bool                              `xml:"isSequential,attr,omitempty" json:"isSequential,omitempty"`
	LoopCardinality     []LoopCardinality                 `xml:"bpmn:loopCardinality,omitempty" json:"loopCardinality,omitempty"`
	CompletionCondition []conditional.CompletionCondition `xml:"bpmn:completionCondition,omitempty" json:"completionCondition,omitempty"`
}

// TMultiInstanceLoopCharacteristics ...
type TMultiInstanceLoopCharacteristics struct {
	IsSequential        bool                              `xml:"isSequential,attr,omitempty" json:"isSequential,omitempty"`
	LoopCardinality     []LoopCardinality                 `xml:"loopCardinality,omitempty" json:"loopCardinality,omitempty"`
	CompletionCondition []conditional.CompletionCondition `xml:"completionCondition,omitempty" json:"completionCondition,omitempty"`
}

// ParticipantMultiplicity ...
type ParticipantMultiplicity struct{}

// TParticipantMultiplicity ...
type TParticipantMultiplicity struct{}

// StandardLoopCharacteristics ...
type StandardLoopCharacteristics struct{}

// TStandardLoopCharacteristics ...
type TStandardLoopCharacteristics struct{}
