package loop

import (
	"github.com/deemount/gobpmnModels/pkg/conditional"
	impl "github.com/deemount/gobpmnTypes"
)

type LoopBaseID interface{}
type LoopBaseName interface{}
type LoopCamundaBase interface{}
type LoopBaseCoreElements interface{}
type LoopBase interface{}

/*
 * @Repositories
 */

// LoopCardinalityRepository ...
type LoopCardinalityRepository interface{}

// MultiInstanceLoopCharacteristicsRepository ...
type MultiInstanceLoopCharacteristicsRepository interface {
	SetIsSequential(is bool)
	GetIsSequential() impl.BOOL_PTR

	SetLoopCardinality()
	GetLoopCardinality() *LoopCardinality

	SetCompletionCondition()
	GetCompletionCondition() *conditional.CompletionCondition
}

// ParticipantMultiplicityRepository ...
type ParticipantMultiplicityRepository interface{}

// StandardLoopCharacteristicsRepository ...
type StandardLoopCharacteristicsRepository interface{}
