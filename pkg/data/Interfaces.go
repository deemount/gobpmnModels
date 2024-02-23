package data

import (
	"github.com/deemount/gobpmnModels/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// DataObjectRepository ...
type DataObjectRepository interface {
	impl.IFBaseID
}

// DataObjectReferenceRepository ...
type DataObjectReferenceRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	attributes.AttributesBaseElements
	SetDataObjectRef(suffix interface{})
	GetDataObjectRef() impl.STR_PTR
}

// DataStoreRepository ...
type DataStoreReferenceRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	attributes.AttributesBaseElements
}
