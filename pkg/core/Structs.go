package core

import (
	"encoding/xml"

	"github.com/deemount/gobpmnModels/pkg/collaboration"
	"github.com/deemount/gobpmnModels/pkg/events"
	"github.com/deemount/gobpmnModels/pkg/marker"
	"github.com/deemount/gobpmnModels/pkg/process"
)

type (

	// DefinitionsBaseElements ...
	DefinitionsBaseElements struct {
		Collaboration collaboration.COLLABORATION_SLC `xml:"bpmn:collaboration,omitempty" json:"collaboration"`
		Process       process.PROCESS_SLC             `xml:"bpmn:process,omitempty" json:"process"`
		Category      marker.CATEGORY_SLC             `xml:"bpmn:category,omitempty" json:"category,omitempty"`
		events.CoreEvents
	}

	// TDefinitionsBaseElements ...
	TDefinitionsBaseElements struct {
		Collaboration collaboration.TCOLLABORATION_SLC `xml:"collaboration,omitempty" json:"collaboration"`
		Process       process.TPROCESS_SLC             `xml:"process,omitempty" json:"process"`
		Category      marker.TCATEGORY_SLC             `xml:"category,omitempty" json:"category,omitempty"`
		events.TCoreEvents
	}

	/*
	 * @Elementary
	 */

	// Definitions represents the root element
	Definitions struct {
		XMLName         xml.Name `xml:"bpmn:definitions" json:"-"`
		Bpmn            string   `xml:"xmlns:bpmn,attr" json:"-"`
		DC              string   `xml:"xmlns:dc,attr,omitempty" json:"-"`
		ID              string   `xml:"id,attr" json:"id"`
		TargetNamespace string   `xml:"targetNamespace,attr" json:"-"`
		DefinitionsBaseElements
	}

	// TDefinitions ...
	TDefinitions struct {
		XMLName xml.Name `xml:"definitions" json:"-"`
		ID      string   `xml:"id,attr" json:"id"`
		TDefinitionsBaseElements
	}
)
