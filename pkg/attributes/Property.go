package attributes

import impl "github.com/deemount/gobpmnTypes"

// NewProperty ...
func NewProperty() PropertyRepository {
	return &Property{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (prop *Property) SetID(typ string, suffix interface{}) {
	prop.ID = SetID(typ, suffix)
}

// SetName ...
// Notice: maybe put a placeholder like name="__targetRef_placeholder" in it
func (prop *Property) SetName(name string) {
	prop.Name = name
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (prop Property) GetID() impl.STR_PTR {
	return &prop.ID

}

// GetName ...
func (prop Property) GetName() impl.STR_PTR {
	return &prop.Name
}
