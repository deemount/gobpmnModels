package gateways

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// ExclusiveGatewayRepository ...
func NewExclusiveGateway() ExclusiveGatewayRepository {
	return &ExclusiveGateway{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (exclusiveGateway *ExclusiveGateway) SetID(typ string, suffix interface{}) {
	exclusiveGateway.ID = SetID(typ, suffix)
}

// SetName ...
func (exclusiveGateway *ExclusiveGateway) SetName(name string) {
	exclusiveGateway.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (exclusiveGateway *ExclusiveGateway) SetDocumentation() {
	exclusiveGateway.Documentation = make([]attributes.Documentation, 1)
}

/*** Marker ***/

// SetIncoming ...
func (exclusiveGateway *ExclusiveGateway) SetIncoming(num int) {
	exclusiveGateway.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (exclusiveGateway *ExclusiveGateway) SetOutgoing(num int) {
	exclusiveGateway.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (exclusiveGateway ExclusiveGateway) GetID() impl.STR_PTR {
	return &exclusiveGateway.ID
}

// GetName ...
func (exclusiveGateway ExclusiveGateway) GetName() impl.STR_PTR {
	return &exclusiveGateway.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (exclusiveGateway ExclusiveGateway) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &exclusiveGateway.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (exclusiveGateway ExclusiveGateway) GetIncoming(num int) *marker.Incoming {
	return &exclusiveGateway.Incoming[num]
}

// GetOutgoing ...
func (exclusiveGateway ExclusiveGateway) GetOutgoing(num int) *marker.Outgoing {
	return &exclusiveGateway.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (exclusiveGateway ExclusiveGateway) String() string {
	return fmt.Sprintf("id=%v, name=%v", exclusiveGateway.ID, exclusiveGateway.Name)
}
