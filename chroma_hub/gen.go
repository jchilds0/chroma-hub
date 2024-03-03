package chroma_hub

import "github.com/jchilds0/chroma-hub/db"

func GenerateTemplateHub(conf map[string]int, filename string) {
    hub := db.NewDataBase()
    hub.AddTemplate(0, "left_to_right", "", "left_to_right")
    
    var total int 
    for geo, count := range conf {
        for i := 0; i < count; i++ {
            hub.AddGeometry(0, total, geo)
            hub.AddAttr(0, total, "parent", "0")
            total++
        }
    }

    ExportArchive(hub, filename)
}
