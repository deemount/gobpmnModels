package process

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/data"
	"github.com/deemount/gobpmnModels/pkg/events"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/gateways"
	"github.com/deemount/gobpmnModels/pkg/marker"
	"github.com/deemount/gobpmnModels/pkg/pool"
	"github.com/deemount/gobpmnModels/pkg/subprocesses"
	"github.com/deemount/gobpmnModels/pkg/tasks"
	impl "github.com/deemount/gobpmnTypes"
)

/*
 * @Elementary
 */

// Process ...
type Process struct {
	impl.BaseAttributes
	attributes.Attributes
	events.ProcessEvents
	tasks.Tasks
	subprocesses.Subprocesses
	gateways.Gateways
	// Base
	IsExecutable bool `xml:"isExecutable,attr" json:"isExecutable,omitempty"`
	// Pool
	LaneSet []pool.LaneSet `xml:"bpmn:laneSet,omitempty" json:"laneSet,omitempty"`
	// Flow
	Association  flow.ASSOCIATION_SLC   `xml:"bpmn:association,omitempty" json:"association,omitempty"`
	SequenceFlow flow.SEQUENCE_FLOW_SLC `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
	// Marker
	Group []marker.Group `xml:"bpmn:group,omitempty" json:"group,omitempty"`
	// Data
	DataObject []data.DataObject `xml:"bpmn:dataObject,omitempty" json:"dataObject,omitempty"`
}

// TProcess ...
type TProcess struct {
	impl.BaseAttributes
	attributes.TAttributes
	events.TProcessEvents
	tasks.TTasks
	subprocesses.TSubprocesses
	gateways.TGateways
	// Base
	IsExecutable bool `xml:"isExecutable,attr" json:"isExecutable,omitempty"`
	// Pool
	LaneSet []pool.LaneSet `xml:"laneSet,omitempty" json:"laneSet,omitempty"`
	// Flow
	Association  flow.TASSOCIATION_SLC   `xml:"association,omitempty" json:"association,omitempty"`
	SequenceFlow flow.TSEQUENCE_FLOW_SLC `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
	// Marker
	Group []marker.TGroup `xml:"group,omitempty" json:"group,omitempty" csv:"-"`
	// Data
	DataObject []data.DataObject `xml:"dataObject,omitempty" json:"dataObject,omitempty"`
}
