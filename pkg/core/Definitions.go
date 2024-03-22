package core

import (
	"fmt"

	"github.com/deemount/gobpmnModels/internals/utils"
	"github.com/deemount/gobpmnModels/pkg/collaboration"
	"github.com/deemount/gobpmnModels/pkg/events"
	"github.com/deemount/gobpmnModels/pkg/events/elements"
	"github.com/deemount/gobpmnModels/pkg/marker"
	"github.com/deemount/gobpmnModels/pkg/process"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

var (
	schemaOMGBpmnModel = "http://www.omg.org/spec/BPMN/20100524/MODEL"
	schemaOMGDC        = "http://www.omg.org/spec/DD/20100524/DC"
	schemaBpmnIOSchema = "http://bpmn.io/schema/bpmn"
)

// NewDefinitions ...
func NewDefinitions() DefinitionsRepository {
	return &Definitions{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (definitions *Definitions) SetID(typ string, suffix interface{}) {
	definitions.ID = SetID(typ, suffix)
}

// SetBpmn ...
func (definitions *Definitions) SetBpmn() {
	definitions.Bpmn = schemaOMGBpmnModel
}

// SetDC ...
func (definitions *Definitions) SetDC() {
	definitions.DC = schemaOMGDC
}

// SetTargetNamespace ...
func (definitions *Definitions) SetTargetNamespace() {
	definitions.TargetNamespace = schemaBpmnIOSchema
}

/* Elements */

/** BPMN **/

// SetCollaboration ...
func (definitions *Definitions) SetCollaboration() {
	definitions.Collaboration = make(collaboration.COLLABORATION_SLC, 1)
}

// SetProcess ...
func (definitions *Definitions) SetProcess(num int) {
	if num == 0 {
		num = 1
	}
	definitions.Process = make(process.PROCESS_SLC, num)
}

// SetCategory ...
func (definitions *Definitions) SetCategory(num int) {
	definitions.Category = make(marker.CATEGORY_SLC, num)
}

/*** Events ***/

// SetMessage ...
func (definitions *Definitions) SetMessage(num int) {
	definitions.Message = make(events.MESSAGE_SLC, num)
}

// SetSignal ...
func (definitions *Definitions) SetSignal(num int) {
	definitions.Signal = make(events.SIGNAL_SLC, num)
}

/*
 * @Settings
 */

// SetDefaultAttributes ...
func (definitions *Definitions) SetDefaultAttributes() {
	definitionsHash := utils.GenerateHash()
	definitions.SetBpmn()
	definitions.SetDC()
	definitions.SetID("definitions", definitionsHash)
	definitions.SetTargetNamespace()
}

// SetMainElements ...
func (definitions *Definitions) SetMainElements(num int) {
	definitions.SetCollaboration()
	definitions.SetProcess(num)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (definitions Definitions) GetID() gobpmnTypes.STR_PTR {
	return &definitions.ID
}

/* Elements */

/** BPMN **/

// GetCollaboration ...
func (definitions Definitions) GetCollaboration() *collaboration.Collaboration {
	return &definitions.Collaboration[0]
}

// GetProcess ...
func (definitions Definitions) GetProcess(num int) *process.Process {
	return &definitions.Process[num]
}

// GetCategory ...
func (definitions Definitions) GetCategory(num int) *marker.Category {
	return &definitions.Category[num]
}

/*** Events ***/

// GetMessage ...
func (definitions Definitions) GetMessage(num int) *elements.Message {
	return &definitions.Message[num]
}

// GetSignal ...
func (definitions Definitions) GetSignal(num int) *elements.Signal {
	return &definitions.Signal[num]
}

/*
 * @String
 */

// String ...
func (definitions Definitions) String() string {
	return fmt.Sprintf("id=%v", definitions.ID)
}

// String ...
func (definitions TDefinitions) String() string {
	return fmt.Sprintf("id=%v", definitions.ID)
}
