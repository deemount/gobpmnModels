package tasks

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
	"github.com/deemount/gobpmnModels/pkg/marker"
)

// NewScriptTask ...
func NewScriptTask() ScriptTaskRepository {
	return &ScriptTask{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (scriptTask *ScriptTask) SetID(typ string, suffix interface{}) {
	scriptTask.ID = SetID(typ, suffix)
}

// SetName ...
func (scriptTask *ScriptTask) SetName(name string) {
	scriptTask.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (scriptTask *ScriptTask) SetDocumentation() {
	scriptTask.Documentation = make([]attributes.Documentation, 1)
}

/*** Marker ***/

// SetIncoming ...
func (scriptTask *ScriptTask) SetIncoming(num int) {
	scriptTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (scriptTask *ScriptTask) SetOutgoing(num int) {
	scriptTask.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (scriptTask ScriptTask) GetID() impl.STR_PTR {
	return &scriptTask.ID
}

// GetName ...
func (scriptTask ScriptTask) GetName() impl.STR_PTR {
	return &scriptTask.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (scriptTask ScriptTask) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &scriptTask.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (scriptTask ScriptTask) GetIncoming(num int) *marker.Incoming {
	return &scriptTask.Incoming[num]
}

// GetOutgoing ...
func (scriptTask ScriptTask) GetOutgoing(num int) *marker.Outgoing {
	return &scriptTask.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (scriptTask ScriptTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", scriptTask.ID, scriptTask.Name)
}
