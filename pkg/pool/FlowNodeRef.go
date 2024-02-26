package pool

import (
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewFlowNodeRef ...
func NewFlowNodeRef() FlowNodeRefRepository {
	return &FlowNodeRef{}
}

/*
 * @Setters
 */

/* Content */

// SetID ...
func (fnr *FlowNodeRef) SetID(typ string, suffix interface{}) {
	fnr.ID = SetID(typ, suffix)
}

/*
 * @Getters
 */

/* Content */

// GetID ...
func (fnr FlowNodeRef) GetID() impl.STR_PTR {
	return &fnr.ID
}
