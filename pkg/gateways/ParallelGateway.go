package gateways

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewParallelGateway ...
func NewParallelGateway() ParallelGatewayRepository {
	return &ParallelGateway{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (parallelGateway *ParallelGateway) SetID(typ string, suffix interface{}) {
	parallelGateway.ID = SetID(typ, suffix)
}

// SetName ...
func (parallelGateway *ParallelGateway) SetName(name string) {
	parallelGateway.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (parallelGateway *ParallelGateway) SetDocumentation() {
	parallelGateway.Documentation = make([]attributes.Documentation, 1)
}

/*** Marker ***/

// SetIncoming ...
func (parallelGateway *ParallelGateway) SetIncoming(num int) {
	parallelGateway.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (parallelGateway *ParallelGateway) SetOutgoing(num int) {
	parallelGateway.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (parallelGateway ParallelGateway) GetID() impl.STR_PTR {
	return &parallelGateway.ID
}

// GetName ...
func (parallelGateway ParallelGateway) GetName() impl.STR_PTR {
	return &parallelGateway.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (parallelGateway ParallelGateway) GetDocumentation() *attributes.Documentation {
	return &parallelGateway.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (parallelGateway ParallelGateway) GetIncoming(num int) *marker.Incoming {
	return &parallelGateway.Incoming[num]
}

// GetOutgoing ...
func (parallelGateway ParallelGateway) GetOutgoing(num int) *marker.Outgoing {
	return &parallelGateway.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (parallelGateway ParallelGateway) String() string {
	return fmt.Sprintf("id=%v, name=%v", parallelGateway.ID, parallelGateway.Name)
}
