package graphics

import (
	"fmt"
	"log"
)

type Template struct {
    Id            int
    Name          string
    Layer         int
    Geo           map[int]*Geometry
    Animate_on    string
    Animate_cont  string
    Animate_off   string
}

func NewTemplate(id int, anim_on, anim_cont, anim_off string) *Template {
    temp := &Template{
        Id: id, 
        Animate_on: anim_on, 
        Animate_cont: anim_cont,
        Animate_off: anim_off,
    }
    temp.Geo = make(map[int]*Geometry)

    return temp
}

func (temp *Template) String() string {
    first := true
    geometry := ""

    for _, geo := range temp.Geo {
        if first {
            geometry = geo.String()
            first = false 
            continue
        }

        geometry = fmt.Sprintf("%s,%s", geometry, geo.String())
    }
    
    return fmt.Sprintf("{'id': %d, 'name': '%s', 'layer': %d, 'num_geo': %d, 'anim_on': '%s', 'anim_cont': '%s', 'anim_off': '%s', 'geometry': [%s]}", 
        temp.Id, temp.Name, temp.Layer, len(temp.Geo), temp.Animate_on, temp.Animate_cont, temp.Animate_off, geometry);
}

func (temp *Template) AddGeometry(geo_id, parent int, geo_type string) {
    if temp.Geo[geo_id] != nil {
        log.Printf("Template %d already contains Geometry %d", temp.Id, geo_id)
        return
    }

    temp.Geo[geo_id] = NewGeometry(geo_id, parent, geo_type)    
}

func (temp *Template) AddAttr(geo_id int, name, attr string) {
    if temp.Geo[geo_id] == nil {
        log.Printf("Template %d does not contain Geometry %d", temp.Id, geo_id)
        return
    }

    temp.Geo[geo_id].AddAttr(name, attr)
}
