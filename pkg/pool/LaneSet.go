package pool

import impl "github.com/deemount/gobpmnTypes"

// NewLaneSet ...
func NewLaneSet() LaneSetRepository {
	return &LaneSet{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ls *LaneSet) SetID(typ string, suffix interface{}) {
	ls.ID = SetID(typ, suffix)
}

/* Elements */

/** BPMN **/

// SetLane ...
func (ls *LaneSet) SetLane(num int) {
	ls.Lane = make(LANE_SLC, num)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ls LaneSet) GetID() impl.STR_PTR {
	return &ls.ID
}

/* Elements */

/** BPMN **/

// GetLane ...
func (ls LaneSet) GetLane(num int) LANE_PTR {
	return &ls.Lane[num]
}
