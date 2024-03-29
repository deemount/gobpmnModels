package conditional

import impl "github.com/deemount/gobpmnTypes"

/*
 * @Base
 */

// ConditionalScriptFormat ...
type ConditionalScriptFormat interface {
	SetScriptFormat(format string)
	GetScriptFormat() impl.STR_PTR
}

// ConditionalScript ...
type ConditionalScript interface {
	SetScript(script string)
	GetScript() impl.STR_PTR
}

// ConditionalType ...
type ConditionalType interface {
	SetConditionType()
	GetConditionType() impl.STR_PTR
}

type ConditionalExpression interface {
	SetExpression(expression string)
	GetExpression() impl.STR_PTR
}

/*
 * @Repositories
 */

// CompletionConditionRepository ...
type CompletionConditionRepository interface{}

// ConditionRepository ...
type ConditionRepository interface {
	ConditionalScriptFormat
	ConditionalScript
	ConditionalType
}

// ConditionExpressionRepository ...
type ConditionExpressionRepository interface {
	ConditionalScriptFormat
	ConditionalScript
	ConditionalType
	ConditionalExpression
}
