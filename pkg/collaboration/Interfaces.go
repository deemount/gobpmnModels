package collaboration

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/loop"
)

/*
 * @Repositories
 */

// CollaborationRepository ...
type CollaborationRepository interface {
	impl.IFBaseID
	attributes.AttributesBaseElements

	SetParticipant(num int)
	GetParticipant(num int) PARTICIPANT_PTR

	SetMessageFlow(num int)
	GetMessageFlow(num int) flow.MESSAGE_FLOW_PTR
}

// ParticipantRepository ...
type ParticipantRepository interface {
	impl.IFBaseID
	impl.IFBaseName

	SetProcessRef(typ string, suffix string)
	GetProcessRef() impl.STR_PTR

	SetDocumentation()
	GetDocumentation() attributes.DOCUMENTATION_PTR

	SetParticipantMultiplicity()
	GetParticipantMultiplicity() *loop.ParticipantMultiplicity
}
