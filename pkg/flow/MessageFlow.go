package flow

import (
	"fmt"
	"log"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewMessageFlow ...
func NewMessageFlow() MessageFlowRepository {
	return &MessageFlow{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (messageFlow *MessageFlow) SetID(typ string, suffix interface{}) {
	messageFlow.ID = SetID(typ, suffix)
}

// SetName ...
func (messageFlow *MessageFlow) SetName(name string) {
	messageFlow.Name = name
}

// SetSourceRef ...
func (messageFlow *MessageFlow) SetSourceRef(typ string, sourceRef interface{}) {
	switch typ {
	case "activity":
		messageFlow.SourceRef = fmt.Sprintf("Activity_%s", sourceRef)
		break
	case "event":
		messageFlow.SourceRef = fmt.Sprintf("Event_%s", sourceRef)
		break
	case "id":
		messageFlow.SourceRef = fmt.Sprintf("%s", sourceRef)
		break
	case "participant":
		messageFlow.SourceRef = fmt.Sprintf("Participant_%s", sourceRef)
		break
	}
}

// SetTargetRef ...
func (messageFlow *MessageFlow) SetTargetRef(typ string, targetRef interface{}) {
	switch typ {
	case "activity":
		messageFlow.TargetRef = fmt.Sprintf("Activity_%s", targetRef)
		break
	case "event":
		messageFlow.TargetRef = fmt.Sprintf("Event_%s", targetRef)
		break
	case "id":
		messageFlow.TargetRef = fmt.Sprintf("%s", targetRef)
		break
	case "participant":
		messageFlow.TargetRef = fmt.Sprintf("Participant_%s", targetRef)
		break
	default:
		log.Panic("no typ set in target ref for message flow")
	}
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (messageFlow *MessageFlow) SetDocumentation() {
	messageFlow.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (messageFlow MessageFlow) GetID() impl.STR_PTR {
	return &messageFlow.ID
}

// GetName ...
func (messageFlow MessageFlow) GetName() impl.STR_PTR {
	return &messageFlow.Name
}

// GetSourceRef ...
func (messageFlow MessageFlow) GetSourceRef() impl.STR_PTR {
	return &messageFlow.SourceRef
}

// GetTargetRef ...
func (messageFlow MessageFlow) GetTargetRef() impl.STR_PTR {
	return &messageFlow.TargetRef
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (messageFlow MessageFlow) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &messageFlow.Documentation[0]
}
