package artifacts

import "github.com/deemount/gobpmnModels/pkg/impl"

// NewText ...
func NewText() TextRepository {
	return &Text{}
}

/*
 * @Setters
 */

/* Content */

/** BPMN **/

// SetText ...
func (txt *Text) SetContent(text string) {
	txt.Text = text
}

/*
 * @Getters
 */

/* Content */

/** BPMN **/

// GetText ...
func (txt Text) GetContent() impl.STR_PTR {
	return &txt.Text
}
