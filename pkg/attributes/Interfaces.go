package attributes

import (
	"github.com/deemount/gobpmnModels/pkg/impl"
)

type AttributesBaseElements interface {
	SetDocumentation()
	GetDocumentation() DOCUMENTATION_PTR
}

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
