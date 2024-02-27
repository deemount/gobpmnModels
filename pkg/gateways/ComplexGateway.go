package gateways

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewComplexGateway ...
func NewComplexGateway() ComplexGatewayRepository {
	return &ComplexGateway{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (complexGateway *ComplexGateway) SetID(typ string, suffix interface{}) {
	complexGateway.ID = SetID(typ, suffix)
}

// SetName ...
func (complexGateway *ComplexGateway) SetName(name string) {
	complexGateway.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (complexGateway *ComplexGateway) SetDocumentation() {
	complexGateway.Documentation = make([]attributes.Documentation, 1)
}

/*** Marker ***/

// SetIncoming ...
func (complexGateway *ComplexGateway) SetIncoming(num int) {
	complexGateway.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (complexGateway *ComplexGateway) SetOutgoing(num int) {
	complexGateway.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (complexGateway ComplexGateway) GetID() impl.STR_PTR {
	return &complexGateway.ID
}

// SetName ...
func (complexGateway ComplexGateway) GetName() impl.STR_PTR {
	return &complexGateway.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (complexGateway ComplexGateway) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &complexGateway.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (complexGateway ComplexGateway) GetIncoming(num int) marker.INCOMING_PTR {
	return &complexGateway.Incoming[num]
}

// GetOutgoing ...
func (complexGateway ComplexGateway) GetOutgoing(num int) marker.OUTGOING_PTR {
	return &complexGateway.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (complexGateway ComplexGateway) String() string {
	return fmt.Sprintf("id=%v, name=%v", complexGateway.ID, complexGateway.Name)
}
