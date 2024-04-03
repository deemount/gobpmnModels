package collaboration

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/loop"
	impl "github.com/deemount/gobpmnTypes"
)

/*
 * @Repositories
 */

// CollaborationRepository ...
type CollaborationRepository interface {
	impl.IFBaseID
	attributes.AttributesBaseElements

	SetParticipant(num int)
	GetParticipant(num int) *Participant

	SetMessageFlow(num int)
	GetMessageFlow(num int) *flow.MessageFlow
}

// ParticipantRepository ...
type ParticipantRepository interface {
	impl.IFBaseID
	impl.IFBaseName

	SetDocumentation()
	GetDocumentation() *attributes.Documentation

	SetProcessRef(typ string, suffix string)
	GetProcessRef() impl.STR_PTR

	SetParticipantMultiplicity()
	GetParticipantMultiplicity() *loop.ParticipantMultiplicity
}
