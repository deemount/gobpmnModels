package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

// NewBusinessRuleTask ...
func NewBusinessRuleTask() BusinessRuleTaskRepository {
	return &BusinessRuleTask{}
}

/*
 * Default Setters
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
	businessRuleTask.Documentation = make([]attributes.Documentation, 1)
}

/*** Marker ***/

// SetIncoming ...
func (businessRuleTask *BusinessRuleTask) SetIncoming(num int) {
	businessRuleTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (businessRuleTask *BusinessRuleTask) SetOutgoing(num int) {
	businessRuleTask.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
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
 * Default String
 */

// String ...
func (businessRuleTask BusinessRuleTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", businessRuleTask.ID, businessRuleTask.Name)
}
