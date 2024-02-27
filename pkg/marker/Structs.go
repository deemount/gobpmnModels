package marker

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
)

/*
 * @Elementary
 */

// IncomingOutgoing ...
type IncomingOutgoing struct {
	Incoming INCOMING_SLC `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing OUTGOING_SLC `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TIncomingOutgoing ...
type TIncomingOutgoing struct {
	Incoming INCOMING_SLC `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing OUTGOING_SLC `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

// Category ...
type Category struct {
	impl.CoreID
	CategoryValue CATEGORY_VALUE_SLC `xml:"bpmn:categoryValue,omitempty" json:"categoryValue,omitempty"`
	attributes.Attributes
}

// TCategory ...
type TCategory struct {
	impl.CoreID
	CategoryValue CATEGORY_VALUE_SLC `xml:"categoryValue,omitempty" json:"categoryValue,omitempty"`
	attributes.TAttributes
}

// CategoryValue ...
type CategoryValue struct {
	impl.CoreID
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}

// Group ...
type Group struct {
	impl.CoreID
	CategoryValueRef string `xml:"categoryValueRef,attr,omitempty" json:"categoryValueRef,omitempty"`
	attributes.Attributes
}

// TGroup ...
type TGroup struct {
	impl.CoreID
	CategoryValueRef string `xml:"categoryValueRef,attr,omitempty" json:"categoryValueRef,omitempty"`
	attributes.TAttributes
}

// Incoming ...
type Incoming struct {
	Flow string `xml:",innerxml" json:"flow"`
}

// Outgoing ...
type Outgoing struct {
	Flow string `xml:",innerxml" json:"flow"`
}
