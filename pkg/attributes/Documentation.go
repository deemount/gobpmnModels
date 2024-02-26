package attributes

// NewDocumenation ...
func NewDocumentation() DocumentationRepository {
	return &Documentation{}
}

/*
 * @Setters
 */

/* Content */

/** BPMN **/

// SetData ...
func (documentation *Documentation) SetData(data string) {
	documentation.Data = data
}

/*
 * @Getters
 */

/* Content */

/** BPMN **/

// GetData ...
func (documentation Documentation) GetData() *string {
	return &documentation.Data
}
