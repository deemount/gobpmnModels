package core

import (
	"encoding/xml"

	"github.com/deemount/gobpmnModels/pkg/collaboration"
	"github.com/deemount/gobpmnModels/pkg/events"
	"github.com/deemount/gobpmnModels/pkg/marker"
	"github.com/deemount/gobpmnModels/pkg/process"
)

// DefinitionsBaseElements ...
type DefinitionsBaseElements struct {
	Collaboration collaboration.COLLABORATION_SLC `xml:"bpmn:collaboration,omitempty" json:"collaboration"`
	Process       process.PROCESS_SLC             `xml:"bpmn:process,omitempty" json:"process"`
	Category      marker.CATEGORY_SLC             `xml:"bpmn:category,omitempty" json:"category,omitempty"`
	events.CoreEvents
}

// TDefinitionsBaseElements ...
type TDefinitionsBaseElements struct {
	Collaboration collaboration.TCOLLABORATION_SLC `xml:"collaboration,omitempty" json:"collaboration"`
	Process       process.TPROCESS_SLC             `xml:"process,omitempty" json:"process"`
	Category      marker.TCATEGORY_SLC             `xml:"category,omitempty" json:"category,omitempty"`
	events.TCoreEvents
}

// Definitions represents the root element
type Definitions struct {
	XMLName           xml.Name `xml:"bpmn:definitions" json:"-"`
	Bpmn              string   `xml:"xmlns:bpmn,attr" json:"-"`
	Xsd               string   `xml:"xmlns:xsd,attr,omitempty" json:"-"`
	Xsi               string   `xml:"xmlns:xsi,omitempty" json:"-"`
	XsiSchemaLocation string   `xml:"xsi:schemaLocation,attr,omitempty" json:"-"`
	DC                string   `xml:"xmlns:dc,attr,omitempty" json:"-"`
	OmgDC             string   `xml:"xmlns:omgdc,attr,omitempty" json:"-"`
	ID                string   `xml:"id,attr" json:"id"`
	TargetNamespace   string   `xml:"targetNamespace,attr" json:"-"`
	DefinitionsBaseElements
}

// TDefinitions ...
type TDefinitions struct {
	XMLName xml.Name `xml:"definitions" json:"-"`
	ID      string   `xml:"id,attr" json:"id"`
	TDefinitionsBaseElements
}
