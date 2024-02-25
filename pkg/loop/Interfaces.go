package loop

import (
	"github.com/deemount/gobpmnModels/pkg/conditional"
	"github.com/deemount/gobpmnModels/pkg/impl"
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
	GetLoopCardinality() LOOP_CARDINALITY_PTR

	SetCompletionCondition()
	GetCompletionCondition() *conditional.CompletionCondition
}

// ParticipantMultiplicityRepository ...
type ParticipantMultiplicityRepository interface{}

// StandardLoopCharacteristicsRepository ...
type StandardLoopCharacteristicsRepository interface{}
