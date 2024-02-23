package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

func NewServiceTask() ServiceTaskRepository {
	return &ServiceTask{}
}

/*
 * Default Setters
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

/*** Make Elements ***/

// SetDocumentation ...
func (serviceTask *ServiceTask) SetDocumentation() {
	serviceTask.Documentation = make([]attributes.Documentation, 1)
}

// SetIncoming ...
func (serviceTask *ServiceTask) SetIncoming(num int) {
	serviceTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (serviceTask *ServiceTask) SetOutgoing(num int) {
	serviceTask.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
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

// GetDocumentation ...
func (serviceTask ServiceTask) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &serviceTask.Documentation[0]
}

// GetIncoming ...
func (serviceTask ServiceTask) GetIncoming(num int) *marker.Incoming {
	return &serviceTask.Incoming[num]
}

// GetOutgoing ...
func (serviceTask ServiceTask) GetOutgoing(num int) *marker.Outgoing {
	return &serviceTask.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (serviceTask ServiceTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", serviceTask.ID, serviceTask.Name)
}
