package artifacts

import impl "github.com/deemount/gobpmnTypes"

/*
 * @Repositories
 */

// TextRepository ...
type TextRepository interface {
	SetContent(text string)
	GetContent() impl.STR_PTR
}

// TextAnnotationRepository ...
type TextAnnotationRepository interface {
	SetText()
	GetText() *Text
}
