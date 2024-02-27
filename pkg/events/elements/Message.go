package elements

import (
	"fmt"

	impl "github.com/deemount/gobpmnTypes"
)

// NewMessage ...
func NewMessage() MessageRepository {
	return &Message{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (message *Message) SetID(typ string, suffix interface{}) {
	message.ID = SetID(typ, suffix)
}

// SetName ...
func (message *Message) SetName(suffix string) {
	message.Name = fmt.Sprintf("Message_%s", suffix)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (message Message) GetID() impl.STR_PTR {
	return &message.ID
}

// GetName ...
func (message Message) GetName() impl.STR_PTR {
	return &message.Name
}

/*
 * @String
 */

// String ...
func (message Message) String() string {
	return fmt.Sprintf("id=%v, name=%v", message.ID, message.Name)
}

// String ...
func (message TMessage) String() string {
	return fmt.Sprintf("id=%v, name=%v", message.ID, message.Name)
}
