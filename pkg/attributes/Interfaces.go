package attributes

import impl "github.com/deemount/gobpmnTypes"

/*
 * @Base
 */

// AttributesBaseElements ...
type AttributesBaseElements interface {
	SetDocumentation()
	GetDocumentation() DOCUMENTATION_PTR
}

/*
 * @Repositories
 */

// DocumentationRepository ...
type DocumentationRepository interface {
	SetData(data string)
	GetData() *string
}

// PropertyRepository ...
type PropertyRepository interface {
	impl.IFBaseID
	impl.IFBaseName
}
