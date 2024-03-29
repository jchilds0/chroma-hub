package chroma_hub

import "github.com/jchilds0/chroma-hub/db"

func GenerateTemplateHub(geo []string, geo_count []int, filename string) {
    hub := db.NewDataBase()
    hub.AddTemplate(0, "left_to_right", "", "left_to_right")
    
    var total int 
    for i := range geo {
        for j := 0; j < geo_count[i]; j++ {
            hub.AddGeometry(0, total, geo[i])
            hub.AddAttr(0, total, "parent", "0")
            total++
        }
    }

    ExportArchive(hub, filename)
}
