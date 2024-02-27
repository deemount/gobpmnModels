package artifacts

import impl "github.com/deemount/gobpmnTypes"

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
