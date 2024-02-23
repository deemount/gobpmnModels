package collaboration

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewCollaboration ...
func NewCollaboration() CollaborationRepository {
	return &Collaboration{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (collaboration *Collaboration) SetID(typ string, suffix interface{}) {
	collaboration.ID = SetID(typ, suffix)
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (collaboration *Collaboration) SetDocumentation() {
	collaboration.Documentation = make([]attributes.Documentation, 1)
}

// SetParticipant ...
func (collaboration *Collaboration) SetParticipant(num int) {
	collaboration.Participant = make(PARTICIPANT_SLC, num)
}

// SetMessageFlow ...
func (collaboration *Collaboration) SetMessageFlow(num int) {
	collaboration.MessageFlow = make([]flow.MessageFlow, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (collaboration Collaboration) GetID() impl.STR_PTR {
	return &collaboration.ID
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (collaboration Collaboration) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &collaboration.Documentation[0]
}

// GetParticipant ...
func (collaboration Collaboration) GetParticipant(num int) PARTICIPANT_PTR {
	return &collaboration.Participant[num]
}

// GetMessageFlow ...
func (collaboration Collaboration) GetMessageFlow(num int) *flow.MessageFlow {
	return &collaboration.MessageFlow[num]
}
