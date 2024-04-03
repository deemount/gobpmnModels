package gateways

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// EventBasedGatewayRepository ...
func NewEventBasedGateway() EventBasedGatewayRepository {
	return &EventBasedGateway{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (eventBasedGateway *EventBasedGateway) SetID(typ string, suffix interface{}) {
	eventBasedGateway.ID = SetID(typ, suffix)
}

// SetName ...
func (eventBasedGateway *EventBasedGateway) SetName(name string) {
	eventBasedGateway.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (eventBasedGateway *EventBasedGateway) SetDocumentation() {
	eventBasedGateway.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

// SetExtensionElements ...
func (eventBasedGateway *EventBasedGateway) SetExtensionElements() {
	eventBasedGateway.ExtensionElements = make(attributes.EXTENSION_ELEMENTS_SLC, 1)
}

/*** Marker ***/

// SetIncoming ...
func (eventBasedGateway *EventBasedGateway) SetIncoming(num int) {
	eventBasedGateway.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (eventBasedGateway *EventBasedGateway) SetOutgoing(num int) {
	eventBasedGateway.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (eventBasedGateway EventBasedGateway) GetID() impl.STR_PTR {
	return &eventBasedGateway.ID
}

// GetName ...
func (eventBasedGateway EventBasedGateway) GetName() impl.STR_PTR {
	return &eventBasedGateway.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (eventBasedGateway *EventBasedGateway) GetDocumentation() *attributes.Documentation {
	return &eventBasedGateway.Documentation[0]
}

// GetExtensionElements ...
func (eventBasedGateway EventBasedGateway) GetExtensionElements() *attributes.ExtensionElements {
	return &eventBasedGateway.ExtensionElements[0]
}

/*** Marker ***/

// GetIncoming ...
func (eventBasedGateway EventBasedGateway) GetIncoming(num int) *marker.Incoming {
	return &eventBasedGateway.Incoming[num]
}

// GetOutgoing ...
func (eventBasedGateway EventBasedGateway) GetOutgoing(num int) *marker.Outgoing {
	return &eventBasedGateway.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (eventBasedGateway EventBasedGateway) String() string {
	return fmt.Sprintf("id=%v, name=%v", eventBasedGateway.ID, eventBasedGateway.Name)
}
