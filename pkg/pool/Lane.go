package pool

import impl "github.com/deemount/gobpmnTypes"

// NewLane ...
func NewLane() LaneRepository {
	return &Lane{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (lane *Lane) SetID(typ string, suffix interface{}) {
	lane.ID = SetID(typ, suffix)
}

/* Elements */

/** BPMN **/

// SetFlowNodeRef ...
func (lane *Lane) SetFlowNodeRef(num int) {
	lane.FlowNodeRef = make(FLOWNODEREF_SLC, num)
}

/*
 * @Getters
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
func (lane Lane) GetFlowNodeRef(num int) *FlowNodeRef {
	return &lane.FlowNodeRef[num]
}
