package elements

import (
	"fmt"

	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewMessage ...
func NewMessage() MessageRepository {
	return &Message{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (msg *Message) SetID(typ string, suffix interface{}) {
	msg.ID = SetID(typ, suffix)
}

// SetName ...
func (msg *Message) SetName(suffix string) {
	msg.Name = fmt.Sprintf("Message_%s", suffix)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (msg Message) GetID() impl.STR_PTR {
	return &msg.ID
}

// GetName ...
func (msg Message) GetName() impl.STR_PTR {
	return &msg.Name
}

/*
 * Default String
 */

// String ...
func (msg Message) String() string {
	return fmt.Sprintf("id=%v, name=%v", msg.ID, msg.Name)
}

// String ...
func (msg TMessage) String() string {
	return fmt.Sprintf("id=%v, name=%v", msg.ID, msg.Name)
}
