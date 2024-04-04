package elements

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events/definitions"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

/*
 * @Elementary
 */

// BoundaryEvent ...
type BoundaryEvent struct {
	impl.BaseAttributes
	definitions.Boundaries
	AttachedToRef  string              `xml:"attachedToRef,attr,omitempty" json:"attachedToRef,omitempty"`
	CancelActivity bool                `xml:"cancelActivity,attr,omitempty" json:"cancelActivity,omitempty"` // maybe @deprecated in new modeler
	Outgoing       marker.OUTGOING_SLC `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TBoundaryEvent ...
type TBoundaryEvent struct {
	impl.BaseAttributes
	definitions.TBoundaries
	AttachedToRef  string              `xml:"attachedToRef,attr,omitempty" json:"attachedToRef,omitempty"`
	CancelActivity bool                `xml:"cancelActivity,attr,omitempty" json:"cancelActivity,omitempty"` // maybe @deprecated in new modeler
	Outgoing       marker.OUTGOING_SLC `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

// EndEvent ...
type EndEvent struct {
	impl.BaseAttributes
	attributes.Attributes
	definitions.EndEvent
	Incoming marker.INCOMING_SLC `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
}

// TEndEvent ...
type TEndEvent struct {
	impl.BaseAttributes
	attributes.TAttributes
	definitions.TEndEvent
	Incoming marker.INCOMING_SLC `xml:"incoming,omitempty" json:"incoming,omitempty"`
}

// IntermediateCatchEvent ...
type IntermediateCatchEvent struct {
	impl.BaseAttributes
	attributes.Attributes
	marker.IncomingOutgoing
	definitions.IntermediateCatchEvent
}

// TIntermediateCatchEvent ...
type TIntermediateCatchEvent struct {
	impl.BaseAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
	definitions.TIntermediateCatchEvent
}

// IntermediateThrowEvent ...
type IntermediateThrowEvent struct {
	impl.BaseAttributes
	attributes.Attributes
	marker.IncomingOutgoing
	definitions.IntermediateThrowEvent
}

// TIntermediateThrowEvent ...
type TIntermediateThrowEvent struct {
	impl.BaseAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
	definitions.TIntermediateThrowEvent
}

// StartEvent ...
type StartEvent struct {
	impl.BaseAttributes
	attributes.Attributes
	definitions.StartEvent
	IsInterrupting bool                `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
	Outgoing       marker.OUTGOING_SLC `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TStartEvent ...
type TStartEvent struct {
	impl.BaseAttributes
	attributes.TAttributes
	definitions.TStartEvent
	IsInterrupting bool                `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
	Outgoing       marker.OUTGOING_SLC `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

// Message ...
type Message struct {
	impl.BaseAttributes
}

// TMessage ...
type TMessage struct {
	impl.BaseAttributes
}

// Signal ...
type Signal struct {
	impl.BaseAttributes
}

// TSignal ...
type TSignal struct {
	impl.BaseAttributes
}

// Error ...
type Error struct {
	impl.BaseAttributes
	ErrorCode string `xml:"errorCode,attr,omitempty" json:"errorCode,omitempty"`
}
