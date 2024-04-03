package attributes

import impl "github.com/deemount/gobpmnTypes"

// Attributes ...
type Attributes struct {
	Documentation     DOCUMENTATION_SLC      `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements EXTENSION_ELEMENTS_SLC `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// TAttributes ...
type TAttributes struct {
	Documentation     DOCUMENTATION_SLC       `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements TEXTENSION_ELEMENTS_SLC `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// Documentation ...
type Documentation struct {
	Data string `xml:",innerxml" json:"documentation,omitempty"`
}

// ExtensionElements ...
type ExtensionElements struct {
}

// TExtensionElements ...
type TExtensionElements struct {
}

// Property ...
type Property struct {
	impl.BaseAttributes
}
