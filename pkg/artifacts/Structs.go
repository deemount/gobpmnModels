package artifacts

/*
 * @Elementary
 */

// TextAnnotation ...
type TextAnnotation struct {
	Text TEXT_SLC `xml:"bpmn:text,omitempty" json:"text,omitempty"`
}

// TTextAnnotation ...
type TTextAnnotation struct {
	Text TEXT_SLC `xml:"text,omitempty" json:"text,omitempty"`
}

// Text ...
type Text struct {
	Text string `xml:",innerxml" json:"text,omitempty"`
}
