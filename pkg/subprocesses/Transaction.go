package subprocesses

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// NewTransaction ...
func NewTransaction() TransactionRepository {
	return &Transaction{}
}

/*
 * @Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (transaction *Transaction) SetID(typ string, suffix interface{}) {
	transaction.ID = SetID(typ, suffix)
}

// SetName ...
func (transaction *Transaction) SetName(name string) {
	transaction.Name = name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetDocumentation ...
func (transaction *Transaction) SetDocumentation() {
	transaction.Documentation = make(attributes.DOCUMENTATION_SLC, 1)
}

/*** Marker ***/

// SetIncoming ...
func (transaction *Transaction) SetIncoming(num int) {
	transaction.Incoming = make(marker.INCOMING_SLC, num)
}

// SetOutgoing ...
func (transaction *Transaction) SetOutgoing(num int) {
	transaction.Outgoing = make(marker.OUTGOING_SLC, num)
}

/*** Flow ***/

// SetSequenceFlow ...
func (transaction *Transaction) SetSequenceFlow(num int) {
	transaction.SequenceFlow = make(flow.SEQUENCE_FLOW_SLC, num)
}

/*
 * @Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (transaction Transaction) GetID() impl.STR_PTR {
	return &transaction.ID
}

// GetName ...
func (transaction Transaction) GetName() impl.STR_PTR {
	return &transaction.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (transaction Transaction) GetDocumentation() *attributes.Documentation {
	return &transaction.Documentation[0]
}

/*** Marker ***/

// GetIncoming ...
func (transaction Transaction) GetIncoming(num int) *marker.Incoming {
	return &transaction.Incoming[num]
}

// GetOutgoing ...
func (transaction Transaction) GetOutgoing(num int) *marker.Outgoing {
	return &transaction.Outgoing[num]
}

// GetSequenceFlow ...
func (transaction Transaction) GetSequenceFlow(num int) *flow.SequenceFlow {
	return &transaction.SequenceFlow[num]
}
