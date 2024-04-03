package gateways

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewInclusiveGateway ...
func NewInclusiveGateway() InclusiveGatewayRepository {
	return &InclusiveGateway{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (inclusiveGateway *InclusiveGateway) SetID(typ string, suffix interface{}) {
	inclusiveGateway.ID = SetID(typ, suffix)
}

// SetName ...
func (inclusiveGateway *InclusiveGateway) SetName(name string) {
	inclusiveGateway.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (inclusiveGateway *InclusiveGateway) SetDocumentation() {
	inclusiveGateway.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

// SetExtensionElements ...
func (inclusiveGateway *InclusiveGateway) SetExtensionElements() {
	inclusiveGateway.ExtensionElements = make(attributes.EXTENSION_ELEMENTS_SLC, 1)
}

/*** Marker ***/

// SetIncoming ...
func (inclusiveGateway *InclusiveGateway) SetIncoming(num int) {
	inclusiveGateway.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (inclusiveGateway *InclusiveGateway) SetOutgoing(num int) {
	inclusiveGateway.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (inclusiveGateway InclusiveGateway) GetID() impl.STR_PTR {
	return &inclusiveGateway.ID

}

// SetName ...
func (inclusiveGateway InclusiveGateway) GetName() impl.STR_PTR {
	return &inclusiveGateway.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (inclusiveGateway InclusiveGateway) GetDocumentation() *attributes.Documentation {
	return &inclusiveGateway.Documentation[0]
}

// GetExtensionElements ...
func (inclusiveGateway InclusiveGateway) GetExtensionElements() *attributes.ExtensionElements {
	return &inclusiveGateway.ExtensionElements[0]
}

/*** Marker ***/

// GetIncoming ...
func (inclusiveGateway InclusiveGateway) GetIncoming(num int) *marker.Incoming {
	return &inclusiveGateway.Incoming[num]
}

// GetOutgoing ...
func (inclusiveGateway InclusiveGateway) GetOutgoing(num int) *marker.Outgoing {
	return &inclusiveGateway.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (inclusiveGateway InclusiveGateway) String() string {
	return fmt.Sprintf("id=%v, name=%v", inclusiveGateway.ID, inclusiveGateway.Name)
}
