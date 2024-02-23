package pool

import (
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// FlowNodeRefRepository ...
type FlowNodeRefRepository interface {
	impl.IFBaseID
}

// LaneRepository ...
type LaneRepository interface {
	impl.IFBaseID

	SetFlowNodeRef(num int)
	GetFlowNodeRef(num int) FLOWNODEREF_PTR
}

// LaneSetRepository ...
type LaneSetRepository interface {
	impl.IFBaseID
	SetLane(num int)
	GetLane(num int) LANE_PTR
}
