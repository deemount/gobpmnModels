package pool

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewLane ...
func NewLane() LaneRepository {
	return &Lane{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (lane *Lane) SetID(typ string, suffix interface{}) {
	switch typ {
	case "lane":
		lane.ID = fmt.Sprintf("Lane_%v", suffix)
		break
	case "id":
		lane.ID = fmt.Sprintf("%s", suffix)
	}
}

/* Elements */

/** BPMN **/

// SetFlowNodeRef ...
func (lane *Lane) SetFlowNodeRef(num int) {
	lane.FlowNodeRef = make([]FlowNodeRef, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (lane Lane) GetID() impl.STR_PTR {
	return &lane.ID
}

/* Elements */

/** BPMN **/

// GetFlowNodeRef ...
func (lane Lane) GetFlowNodeRef(num int) FLOWNODEREF_PTR {
	return &lane.FlowNodeRef[num]
}
