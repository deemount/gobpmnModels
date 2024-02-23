package gateways

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
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

/*** Make Elements ***/

/*** Attributes ***/

// SetDocumentation ...
func (eventBasedGateway *EventBasedGateway) SetDocumentation() {
	eventBasedGateway.Documentation = make([]attributes.Documentation, 1)
}

/*** Marker ***/

// SetIncoming ...
func (eventBasedGateway *EventBasedGateway) SetIncoming(num int) {
	eventBasedGateway.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (eventBasedGateway *EventBasedGateway) SetOutgoing(num int) {
	eventBasedGateway.Outgoing = make([]marker.Outgoing, num)
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

/*** Attributes ***/

// GetDocumentation ...
func (eventBasedGateway *EventBasedGateway) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &eventBasedGateway.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (eventBasedGateway EventBasedGateway) GetIncoming(num int) marker.INCOMING_PTR {
	return &eventBasedGateway.Incoming[num]
}

// GetOutgoing ...
func (eventBasedGateway EventBasedGateway) GetOutgoing(num int) marker.OUTGOING_PTR {
	return &eventBasedGateway.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (eventBasedGateway EventBasedGateway) String() string {
	return fmt.Sprintf("id=%v, name=%v", eventBasedGateway.ID, eventBasedGateway.Name)
}