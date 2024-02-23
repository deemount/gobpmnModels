package flow

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewAssociation ...
func NewAssociation() AssociationRepository {
	return &Association{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (association *Association) SetID(typ string, suffix interface{}) {
	association.ID = SetID(typ, suffix)
}

// SetSourceRef ...
func (association *Association) SetSourceRef(typ string, sourceRef interface{}) {
	association.SourceRef = SetSourceTargetRef(typ, sourceRef)
}

// SetTargetRef ...
func (association *Association) SetTargetRef(typ string, targetRef interface{}) {
	association.TargetRef = SetSourceTargetRef(typ, targetRef)
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (association *Association) SetDocumentation() {
	association.Documentation = make([]attributes.Documentation, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (association Association) GetID() impl.STR_PTR {
	return &association.ID
}

// GetSourceRef ...
func (association Association) GetSourceRef() impl.STR_PTR {
	return &association.SourceRef
}

// GetTargetRef ...
func (association Association) GetTargetRef() impl.STR_PTR {
	return &association.TargetRef
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (association Association) GetDocumentation() attributes.DOCUMENTATION_PTR {
	return &association.Documentation[0]
}
