package artifacts

import "github.com/deemount/gobpmnModels/pkg/impl"

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
	GetText() TEXT_PTR
}
