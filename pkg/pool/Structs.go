package pool

import "github.com/deemount/gobpmnModels/pkg/impl"

/*
 * @Elementary
 */

// FlowNodeRef ...
type FlowNodeRef struct {
	impl.CoreInnerID
}

// Lane ...
type Lane struct {
	impl.CoreID
	FlowNodeRef []FlowNodeRef `xml:"bpmn:flowNodeRef,omitempty" json:"flowNodeRef,omitempty"`
}

// TLane ...
type TLane struct {
	impl.CoreID
	FlowNodeRef []FlowNodeRef `xml:"flowNodeRef,omitempty" json:"flowNodeRef,omitempty"`
}

// LaneSet ...
type LaneSet struct {
	impl.CoreID
	Lane LANE_SLC `xml:"bpmn:lane,omitempty" json:"lane,omitempty"`
}

// TLaneSet ...
type TLaneSet struct {
	impl.CoreID
	Lane TLANE_SLC `xml:"lane,omitempty" json:"lane,omitempty"`
}
