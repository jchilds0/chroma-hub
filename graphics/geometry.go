package graphics

import (
	"fmt"
)

type Geometry struct {
    Id          int
    Parent      int
    GeoType     string 
    Attrs       []*Attr
}

func NewGeometry(id, parent int, geo_type string) *Geometry {
    geo := &Geometry{Id: id, Parent: parent, GeoType: geo_type}
    geo.Attrs = make([]*Attr, 0, 10)

    return geo
}

func (geo *Geometry) String() (str string) {
    first := true
    attributes := ""

    for _, attr := range geo.Attrs {
        if first {
            attributes = attr.String()
            first = false 
            continue
        }

        attributes = fmt.Sprintf("%s,%s", attributes, attr.String())
    }
    
    return fmt.Sprintf("{'id': %d, 'type': '%s', 'parent': %d, 'attr': [%s]}", geo.Id, geo.GeoType, geo.Parent, attributes);
}

func (geo *Geometry) AddAttr(name, value string) {
    geo.Attrs = append(geo.Attrs, NewAttr(name, value))
}
