package gateways

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

/*
 * @Base
 */

// GatewayID ...
type GatewayID interface {
	SetID(typ string, suffix interface{})
	GetID() impl.STR_PTR
}

// GatewayName ...
type GatewayName interface {
	SetName(name string)
	GetName() impl.STR_PTR
}

// GatewayBase ...
type GatewayBase interface {
	GatewayID
	GatewayName
	marker.MarkerIncomingOutgoing
	attributes.AttributesBaseElements
}

/*
 * @Repositories
 */

// GatewaysElementsRepository ...
type GatewaysElementsRepository interface {
	SetExclusiveGateway(num int)
	GetExclusiveGateway(num int) *ExclusiveGateway
	SetInclusiveGateway(num int)
	GetInclusiveGateway(num int) *InclusiveGateway
	SetParallelGateway(num int)
	GetParallelGateway(num int) *ParallelGateway
	SetComplexGateway(num int)
	GetComplexGateway(num int) *ComplexGateway
	SetEventBasedGateway(num int)
	GetEventBasedGateway(num int) *EventBasedGateway
}

// ComplexGatewayRepository ...
type ComplexGatewayRepository interface {
	GatewayBase
}

// EventBasedGatewayRepository ...
type EventBasedGatewayRepository interface {
	GatewayBase
}

// ExclusiveGatewayRepository ...
type ExclusiveGatewayRepository interface {
	GatewayBase
}

// InclusiveGatewayRepository ...
type InclusiveGatewayRepository interface {
	GatewayBase
}

// ParallelGatewayRepository ...
type ParallelGatewayRepository interface {
	GatewayBase
}
