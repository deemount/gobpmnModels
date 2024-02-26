package impl

// CoreID ...
type CoreID struct {
	ID string `xml:"id,attr,omitempty" json:"id" csv:"ID"`
}

// CoreInnerID ...
type CoreInnerID struct {
	ID string `xml:",innerxml" json:"id" csv:"ID"`
}

// BaseAttributes ...
type BaseAttributes struct {
	ID   string `xml:"id,attr,omitempty" json:"id" csv:"ID"`
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" csv:"NAME"`
}
