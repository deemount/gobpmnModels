package artifacts

// NewTextAnnotation ...
func NewTextAnnotation() TextAnnotationRepository {
	return &TextAnnotation{}
}

/*
 * @Setters
 */

/* Elements */

/** BPMN **/

// SetText ...
func (textannotation *TextAnnotation) SetText() {
	textannotation.Text = make(TEXT_SLC, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetText ...
func (textannotation TextAnnotation) GetText() *Text {
	return &textannotation.Text[0]
}
