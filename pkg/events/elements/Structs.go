package elements

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/events/definitions"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

/*
 * @Elementary
 */

// BoundaryEvent ...
type BoundaryEvent struct {
	impl.BaseAttributes
	definitions.Boundaries
	AttachedToRef  string            `xml:"attachedToRef,attr,omitempty" json:"attachedToRef,omitempty"`
	CancelActivity bool              `xml:"cancelActivity,attr,omitempty" json:"cancelActivity,omitempty"` // maybe @deprecated in new modeler
	Outgoing       []marker.Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TBoundaryEvent ...
type TBoundaryEvent struct {
	impl.BaseAttributes
	definitions.TBoundaries
	AttachedToRef  string            `xml:"attachedToRef,attr,omitempty" json:"attachedToRef,omitempty"`
	CancelActivity bool              `xml:"cancelActivity,attr,omitempty" json:"cancelActivity,omitempty"` // maybe @deprecated in new modeler
	Outgoing       []marker.Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

// EndEvent ...
type EndEvent struct {
	impl.BaseAttributes
	attributes.Attributes
	definitions.EndEvent
	Incoming []marker.Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
}

// TEndEvent ...
type TEndEvent struct {
	impl.BaseAttributes
	attributes.TAttributes
	definitions.TEndEvent
	Incoming []marker.Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
}

// IntermediateCatchEvent ...
type IntermediateCatchEvent struct {
	impl.BaseAttributes
	attributes.Attributes
	marker.TIncomingOutgoing
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
	IsInterrupting bool              `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
	Outgoing       []marker.Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TStartEvent ...
type TStartEvent struct {
	impl.BaseAttributes
	attributes.TAttributes
	definitions.TStartEvent
	IsInterrupting bool              `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
	FormKey        string            `xml:"formKey,attr,omitempty" json:"formKey,omitempty"`
	FormRef        string            `xml:"formRef,attr,omitempty" json:"formRef,omitempty"`
	FormRefBind    string            `xml:"formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	FormRefVersion string            `xml:"formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	Initiator      string            `xml:"initiator,attr,omitempty" json:"init,omitempty"`
	Outgoing       []marker.Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
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
