package events

// ProcessEvents ...
type ProcessEvents struct {
	StartEvent             START_EVENT_SLC              `xml:"bpmn:startEvent,omitempty" json:"startEvent,omitempty"`
	BoundaryEvent          BOUNDARY_EVENT_SLC           `xml:"bpmn:boundaryEvent,omitempty" json:"boundaryEvent,omitempty"`
	EndEvent               END_EVENT_SLC                `xml:"bpmn:endEvent,omitempty" json:"endEvent,omitempty"`
	IntermediateCatchEvent INTERMEDIATE_CATCH_EVENT_SLC `xml:"bpmn:intermediateCatchEvent,omitempty" json:"intermediateCatchEvent,omitempty"`
	IntermediateThrowEvent INTERMEDIATE_THROW_EVENT_SLC `xml:"bpmn:intermediateThrowEvent,omitempty" json:"intermediateThrowEvent,omitempty"`
}

// TProcessEvents ...
type TProcessEvents struct {
	StartEvent             TSTART_EVENT_SLC              `xml:"startEvent,omitempty" json:"startEvent,omitempty"`
	BoundaryEvent          TBOUNDARY_EVENT_SLC           `xml:"boundaryEvent,omitempty" json:"boundaryEvent,omitempty"`
	EndEvent               TEND_EVENT_SLC                `xml:"endEvent,omitempty" json:"endEvent,omitempty"`
	IntermediateCatchEvent TINTERMEDIATE_CATCH_EVENT_SLC `xml:"intermediateCatchEvent,omitempty" json:"intermediateCatchEvent,omitempty"`
	IntermediateThrowEvent TINTERMEDIATE_THROW_EVENT_SLC `xml:"intermediateThrowEvent,omitempty" json:"intermediateThrowEvent,omitempty"`
}

// CoreEvents ...
type CoreEvents struct {
	Message MESSAGE_SLC `xml:"bpmn:message,omitempty" json:"message,omitempty"`
	Signal  SIGNAL_SLC  `xml:"bpmn:signal,omitempty" json:"signal,omitempty"`
}

// TCoreEvents ...
type TCoreEvents struct {
	Message TMESSAGE_SLC `xml:"message,omitempty" json:"message,omitempty"`
	Signal  TSIGNAL_SLC  `xml:"signal,omitempty" json:"signal,omitempty"`
}
