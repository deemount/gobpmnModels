package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewServiceTask ...
func NewServiceTask() ServiceTaskRepository {
	return &ServiceTask{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (serviceTask *ServiceTask) SetID(typ string, suffix interface{}) {
	serviceTask.ID = SetID(typ, suffix)
}

// SetName ...
func (serviceTask *ServiceTask) SetName(name string) {
	serviceTask.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (serviceTask *ServiceTask) SetDocumentation() {
	serviceTask.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*** Marker ***/

// SetIncoming ...
func (serviceTask *ServiceTask) SetIncoming(num int) {
	serviceTask.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (serviceTask *ServiceTask) SetOutgoing(num int) {
	serviceTask.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (serviceTask ServiceTask) GetID() impl.STR_PTR {
	return &serviceTask.ID
}

// GetName ...
func (serviceTask ServiceTask) GetName() impl.STR_PTR {
	return &serviceTask.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (serviceTask ServiceTask) GetDocumentation() *attributes.Documentation {
	return &serviceTask.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (serviceTask ServiceTask) GetIncoming(num int) *marker.Incoming {
	return &serviceTask.Incoming[num]
}

// GetOutgoing ...
func (serviceTask ServiceTask) GetOutgoing(num int) *marker.Outgoing {
	return &serviceTask.Outgoing[num]
}

/*
 * @String
 */

// String ...
func (serviceTask ServiceTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", serviceTask.ID, serviceTask.Name)
}
