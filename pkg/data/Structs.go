package data

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// DataObject ...
type DataObject struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

// DataObjectReference ...
type DataObjectReference struct {
	impl.BaseAttributes
	attributes.Attributes
	DataObjectRef string `xml:"dataObjectRef,attr,omitempty" json:"dataObjectRef"`
}

// TDataObjectReference ...
type TDataObjectReference struct {
	impl.BaseAttributes
	attributes.TAttributes
	DataObjectRef string `xml:"dataObjectRef,attr,omitempty" json:"dataObjectRef"`
}

// DataStoreReference ...
type DataStoreReference struct {
	impl.BaseAttributes
	attributes.Attributes
}

// TDataStoreReference ...
type TDataStoreReference struct {
	impl.BaseAttributes
	attributes.TAttributes
}
