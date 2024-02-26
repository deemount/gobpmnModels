package collaboration

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/loop"
)

/*
 * @Elementary
 */

// Collaboration ...
type Collaboration struct {
	impl.CoreID
	attributes.Attributes
	Participant PARTICIPANT_SLC       `xml:"bpmn:participant" json:"participant,omitempty"`
	MessageFlow flow.MESSAGE_FLOW_SLC `xml:"bpmn:messageFlow,omitempty" json:"messageFlow,omitempty"`
}

// TCollaboration ...
type TCollaboration struct {
	impl.CoreID
	attributes.TAttributes
	Participant TPARTICIPANT_SLC       `xml:"participant" json:"participant,omitempty"`
	MessageFlow flow.TMESSAGE_FLOW_SLC `xml:"messageFlow,omitempty" json:"messageFlow,omitempty"`
}

// Participant ...
type Participant struct {
	impl.BaseAttributes
	ProcessRef              string                         `xml:"processRef,attr,omitempty" json:"processRef,omitempty" csv:"PROCESS_REF"`
	Documentation           attributes.DOCUMENTATION_SLC   `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty" csv:"DOCUMENTATION"`
	ParticipantMultiplicity []loop.ParticipantMultiplicity `xml:"bpmn:participantMultiplicity,omitempty" json:"participantMultiplicity,omitempty" csv:"PARTICIPANT_MULTIPLICITY"`
}

// TParticipant ...
type TParticipant struct {
	impl.BaseAttributes
	ProcessRef              string                          `xml:"processRef,attr" json:"processRef,omitempty"`
	Documentation           attributes.DOCUMENTATION_SLC    `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ParticipantMultiplicity []loop.TParticipantMultiplicity `xml:"participantMultiplicity,omitempty" json:"participantMultiplicity,omitempty"`
}
