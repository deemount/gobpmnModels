package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewBusinessRuleTask ...
func NewBusinessRuleTask() BusinessRuleTaskRepository {
	return &BusinessRuleTask{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (businessRuleTask *BusinessRuleTask) SetID(typ string, suffix interface{}) {
	businessRuleTask.ID = SetID(typ, suffix)
}

// SetName ...
func (businessRuleTask *BusinessRuleTask) SetName(name string) {
	businessRuleTask.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (businessRuleTask *BusinessRuleTask) SetDocumentation() {
	businessRuleTask.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*** Marker ***/

// SetIncoming ...
func (businessRuleTask *BusinessRuleTask) SetIncoming(num int) {
	businessRuleTask.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (businessRuleTask *BusinessRuleTask) SetOutgoing(num int) {
	businessRuleTask.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (businessRuleTask BusinessRuleTask) GetID() impl.STR_PTR {
	return &businessRuleTask.ID
}

// GetName ...
func (businessRuleTask BusinessRuleTask) GetName() impl.STR_PTR {
	return &businessRuleTask.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (businessRuleTask BusinessRuleTask) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &businessRuleTask.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (businessRuleTask BusinessRuleTask) GetIncoming(num int) *marker.Incoming {
	return &businessRuleTask.Incoming[num]
}

// GetOutgoing ...
func (businessRuleTask BusinessRuleTask) GetOutgoing(num int) *marker.Outgoing {
	return &businessRuleTask.Outgoing[num]
}

/*
 * @String
 */

// String ...
func (businessRuleTask BusinessRuleTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", businessRuleTask.ID, businessRuleTask.Name)
}
