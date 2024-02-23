package core

import (
	"fmt"

	"github.com/deemount/gobpmnModels/internals/utils"
	"github.com/deemount/gobpmnModels/pkg/collaboration"
	"github.com/deemount/gobpmnModels/pkg/events"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
	"github.com/deemount/gobpmnModels/pkg/process"
)

var (
	schemaBpmnModel           = "http://www.omg.org/spec/BPMN/20100524/MODEL"
	schemaBpmn20              = "http://www.omg.org/spec/BPMN/2.0/20100501/BPMN20.xsd"
	schemaBpmn20_HTTPS        = "https://www.omg.org/spec/BPMN/20100501/BPMN20.xsd"
	schemaBpmnDI              = "http://www.omg.org/spec/BPMN/20100524/DI"
	schemaOMGDI               = "http://www.omg.org/spec/DD/20100524/DI"
	schemaOMGDC               = "http://www.omg.org/spec/DD/20100524/DC"
	schemaBpmnIOSchema        = "http://bpmn.io/schema/bpmn"
	schemaBpmnIOBioColor      = "http://bpmn.io/schema/bpmn/biocolor/1.0"
	schemaW3XmlSchema         = "http://www.w3.org/2001/XMLSchema"
	schemaW3XmlSchemaInstance = "http://www.w3.org/2001/XMLSchema-instance"
)

// NewDefinitions ...
func NewDefinitions() DefinitionsRepository {
	return &Definitions{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetBpmn ...
func (definitions *Definitions) SetBpmn() {
	definitions.Bpmn = schemaBpmnModel
}

// SetBpmnDI ...
func (definitions *Definitions) SetBpmnDI() {
	definitions.BpmnDI = schemaBpmnDI
}

// SetOmgDI ...
func (definitions *Definitions) SetOmgDI() {
	definitions.OmgDI = schemaOMGDI
}

// SetDC ...
func (definitions *Definitions) SetDC() {
	definitions.DC = schemaOMGDC
}

// SetOmgDC ...
func (definitions *Definitions) SetOmgDC() {
	definitions.OmgDC = schemaOMGDC
}

// SetBioc ...
func (definitions *Definitions) SetBioc() {
	definitions.Bioc = schemaBpmnIOBioColor
}

// SetXSD ...
func (definitions *Definitions) SetXSD() {
	definitions.Xsd = schemaW3XmlSchema
}

// SetXSI ...
func (definitions *Definitions) SetXSI() {
	definitions.Xsi = schemaW3XmlSchemaInstance
}

// SetXsiSchemaLocation ...
func (definitions *Definitions) SetXsiSchemaLocation() {
	definitions.XsiSchemaLocation = fmt.Sprintf("%s %s", schemaBpmnModel, schemaBpmn20)
}

// SetXsiSchemaLocation ...
func (definitions *Definitions) SetXsiSchemaLocationHTTPS() {
	definitions.XsiSchemaLocation = fmt.Sprintf("%s %s", schemaBpmnModel, schemaBpmn20_HTTPS)
}

// SetID ...
func (definitions *Definitions) SetID(typ string, suffix interface{}) {
	switch typ {
	case "definitions":
		definitions.ID = fmt.Sprintf("Definitions_%v", suffix)
		break
	case "id":
		definitions.ID = fmt.Sprintf("%s", suffix)
	}
}

// SetTargetNamespace ...
func (definitions *Definitions) SetTargetNamespace() {
	definitions.TargetNamespace = schemaBpmnIOSchema
}

/*** Make Elements ***/

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
 * Default Settings
 */

// SetDefinitionsAttributes ...
func (definitions *Definitions) SetDefaultAttributes() {
	definitionsHash := utils.GenerateHash()
	definitions.SetBpmn()
	definitions.SetBpmnDI()
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
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (definitions Definitions) GetID() impl.STR_PTR {
	return &definitions.ID
}

/* Elements */

/** BPMN **/

// GetCollaboration ...
func (definitions Definitions) GetCollaboration() collaboration.COLLABORATION_PTR {
	return &definitions.Collaboration[0]
}

// GetProcess ...
func (definitions Definitions) GetProcess(num int) process.PROCESS_PTR {
	return &definitions.Process[num]
}

// GetCategory ...
func (definitions Definitions) GetCategory(num int) marker.CATEGORY_PTR {
	return &definitions.Category[num]
}

/*** Events ***/

// GetMessage ...
func (definitions Definitions) GetMessage(num int) events.MESSAGE_PTR {
	return &definitions.Message[num]
}

// GetSignal ...
func (definitions Definitions) GetSignal(num int) events.SIGNAL_PTR {
	return &definitions.Signal[num]
}

/*
 * Default String
 */

// String ...
func (definitions Definitions) String() string {
	return fmt.Sprintf("id=%v", definitions.ID)
}

// String ...
func (definitions TDefinitions) String() string {
	return fmt.Sprintf("id=%v", definitions.ID)
}