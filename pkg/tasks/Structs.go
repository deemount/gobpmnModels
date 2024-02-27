package tasks

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/flow"
	"github.com/deemount/gobpmnModels/pkg/marker"
	impl "github.com/deemount/gobpmnTypes"
)

// Tasks ...
type Tasks struct {
	BusinessRuleTask BUSINESS_RULE_TASK_SLC `xml:"bpmn:businessRuleTask,omitempty" json:"businessRuleTask,omitempty" csv:"-"`
	Task             TASK_SLC               `xml:"bpmn:task,omitempty" json:"task,omitempty" csv:"-"`
	UserTask         USER_TASK_SLC          `xml:"bpmn:userTask,omitempty" json:"userTask,omitempty" csv:"-"`
	ManualTask       MANUAL_TASK_SLC        `xml:"bpmn:manualTask,omitempty" json:"manualTask,omitempty" csv:"-"`
	ReceiveTask      RECEIVE_TASK_SLC       `xml:"bpmn:receiveTask,omitempty" json:"receiveTask,omitempty" csv:"-"`
	ScriptTask       SCRIPT_TASK_SLC        `xml:"bpmn:scriptTask,omitempty" json:"scriptTask,omitempty" csv:"-"`
	SendTask         SEND_TASK_SLC          `xml:"bpmn:sendTask,omitempty" json:"sendTask,omitempty" csv:"-"`
	ServiceTask      SERVICE_TASK_SLC       `xml:"bpmn:serviceTask,omitempty" json:"serviceTask,omitempty" csv:"-"`
}

// TTasks ...
type TTasks struct {
	BusinessRuleTask BUSINESS_RULE_TASK_SLC `xml:"businessRuleTask,omitempty" json:"businessRuleTask,omitempty" csv:"-"`
	Task             TASK_SLC               `xml:"task,omitempty" json:"task,omitempty" csv:"-"`
	UserTask         USER_TASK_SLC          `xml:"userTask,omitempty" json:"userTask,omitempty" csv:"-"`
	ManualTask       MANUAL_TASK_SLC        `xml:"manualTask,omitempty" json:"manualTask,omitempty" csv:"-"`
	ReceiveTask      RECEIVE_TASK_SLC       `xml:"receiveTask,omitempty" json:"receiveTask,omitempty" csv:"-"`
	ScriptTask       SCRIPT_TASK_SLC        `xml:"scriptTask,omitempty" json:"scriptTask,omitempty" csv:"-"`
	SendTask         SEND_TASK_SLC          `xml:"sendTask,omitempty" json:"sendTask,omitempty" csv:"-"`
	ServiceTask      SERVICE_TASK_SLC       `xml:"serviceTask,omitempty" json:"serviceTask,omitempty" csv:"-"`
}

/*
 * @Elementary
 */

// BusinessRuleTask ...
type BusinessRuleTask struct {
	impl.BaseAttributes
	attributes.Attributes
	marker.IncomingOutgoing
}

// TBusinessRuleTask ...
type TBusinessRuleTask struct {
	impl.BaseAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
}

// ManualTask ...
type ManualTask struct {
	impl.BaseAttributes
	attributes.Attributes
	marker.IncomingOutgoing
}

// TManualTask ...
type TManualTask struct {
	impl.BaseAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
}

// ReceiveTask ...
type ReceiveTask struct {
	impl.BaseAttributes
	attributes.Attributes
	marker.IncomingOutgoing
	MessageRef string `xml:"messageRef,attr,omitempty" json:"messageRef,omitempty"`
}

// TReceiveTask ...
type TReceiveTask struct {
	impl.BaseAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
	MessageRef string `xml:"messageRef,attr,omitempty" json:"messageRef,omitempty"`
}

// ScriptTask ...
type ScriptTask struct {
	impl.BaseAttributes
	attributes.Attributes
	marker.IncomingOutgoing
}

// TScriptTask ...
type TScriptTask struct {
	impl.BaseAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
}

// SendTask ...
type SendTask struct {
	impl.BaseAttributes
	attributes.Attributes
	marker.IncomingOutgoing
}

// TSendTask ...
type TSendTask struct {
	impl.BaseAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
}

// ServiceTask ...
type ServiceTask struct {
	impl.BaseAttributes
	attributes.Attributes
	marker.IncomingOutgoing
}

// TServiceTask ...
type TServiceTask struct {
	impl.BaseAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
}

// Task ...
type Task struct {
	impl.BaseAttributes
	attributes.Attributes
	marker.IncomingOutgoing
	Property             attributes.PROPERTY_SLC         `xml:"bpmn:property,omitempty" json:"property,omitempty"`
	DataInputAssociation flow.DATA_INPUT_ASSOCIATION_SLC `xml:"bpmn:dataInputAssociation,omitempty" json:"dataInputAssociation,omitempty"`
}

// TTask ...
type TTask struct {
	impl.BaseAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
	Property             attributes.PROPERTY_SLC         `xml:"property,omitempty" json:"property,omitempty"`
	DataInputAssociation flow.DATA_INPUT_ASSOCIATION_SLC `xml:"dataInputAssociation,omitempty" json:"dataInputAssociation,omitempty"`
}

// UserTask ...
type UserTask struct {
	impl.BaseAttributes
	attributes.Attributes
	marker.IncomingOutgoing
}

// TUserTask ...
type TUserTask struct {
	impl.BaseAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
}
