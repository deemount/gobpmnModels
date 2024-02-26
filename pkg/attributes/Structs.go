package attributes

import "github.com/deemount/gobpmnModels/pkg/impl"

// Attributes ...
type Attributes struct {
	Documentation DOCUMENTATION_SLC `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
}

// TAttributes ...
type TAttributes struct {
	Documentation DOCUMENTATION_SLC `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

// Documentation ...
type Documentation struct {
	Data string `xml:",innerxml" json:"documentation,omitempty"`
}

// Property ...
type Property struct {
	impl.BaseAttributes
}
