package chroma_hub 

// func baseArchive() {
//     // Teal Box
//     temp_id := 0
//     hub.AddTemplate(temp_id, "left_to_right", "", "left_to_right")
//
//     geo_id := 0
//     hub.AddGeometry(temp_id, geo_id, 0, "rect")
//     hub.AddAttr(temp_id, geo_id, "color", "2 132 130 255")
//     hub.AddAttr(temp_id, geo_id, "rounding", "20")
//
//     geo_id = 1
//     hub.AddGeometry(temp_id, geo_id, 0, "circle")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//
//     geo_id = 2
//     hub.AddGeometry(temp_id, geo_id, 0, "text")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//     hub.AddAttr(temp_id, geo_id, "scale", "1.0")
//
//     geo_id = 3
//     hub.AddGeometry(temp_id, geo_id, 0, "text")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//     hub.AddAttr(temp_id, geo_id, "scale", "1.0")
//
//     // Clock Box
//     temp_id = 1
//     hub.AddTemplate(temp_id, "left_to_right", "clock_tick", "left_to_right")
//
//     // bg
//     geo_id = 0
//     hub.AddGeometry(temp_id, geo_id, 0, "rect")
//     hub.AddAttr(temp_id, geo_id, "color", "5 68 94 255")
//     hub.AddAttr(temp_id, geo_id, "rounding", "10")
//
//     // circle
//     geo_id = 1
//     hub.AddGeometry(temp_id, geo_id, 0, "circle")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//
//     // left split
//     geo_id = 2
//     hub.AddGeometry(temp_id, geo_id, 0, "rect")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//
//     // team 1
//     geo_id = 3
//     hub.AddGeometry(temp_id, geo_id, 2, "text")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//     hub.AddAttr(temp_id, geo_id, "scale", "1.0")
//     
//     // score 1
//     geo_id = 4
//     hub.AddGeometry(temp_id, geo_id, 2, "text")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//     hub.AddAttr(temp_id, geo_id, "scale", "1.0")
//
//     // mid split
//     geo_id = 5
//     hub.AddGeometry(temp_id, geo_id, 2, "rect")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//
//     // team 2
//     geo_id = 6
//     hub.AddGeometry(temp_id, geo_id, 5, "text")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//     hub.AddAttr(temp_id, geo_id, "scale", "1.0")
//
//     // score 2
//     geo_id = 7
//     hub.AddGeometry(temp_id, geo_id, 5, "text")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//     hub.AddAttr(temp_id, geo_id, "scale", "1.0")
//
//     // right split
//     geo_id = 8
//     hub.AddGeometry(temp_id, geo_id, 5, "rect")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//
//     // clock
//     geo_id = 9
//     hub.AddGeometry(temp_id, geo_id, 8, "text")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//     hub.AddAttr(temp_id, geo_id, "scale", "1.0")
//
//     // clock text
//     geo_id = 10
//     hub.AddGeometry(temp_id, geo_id, 8, "text")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//     hub.AddAttr(temp_id, geo_id, "scale", "1.0")
//
//     // White circle
//     temp_id = 2
//     hub.AddTemplate(temp_id, "", "", "")
//
//     geo_id = 0
//     hub.AddGeometry(temp_id, geo_id, 0, "circle")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//
//     // Graph 
//     temp_id = 3
//     hub.AddTemplate(temp_id, "left_to_right", "", "left_to_right")

//     geo_id = 0
//     hub.AddGeometry(temp_id, geo_id, 0, "rect")
//     hub.AddAttr(temp_id, geo_id, "color", "2 132 130 255")
//     hub.AddAttr(temp_id, geo_id, "rounding", "20")
//
//     geo_id = 1
//     hub.AddGeometry(temp_id, geo_id, 0, "graph")
//     hub.AddAttr(temp_id, geo_id, "graph_type", "step")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//
//     geo_id = 2
//     hub.AddGeometry(temp_id, geo_id, 0, "text")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//     hub.AddAttr(temp_id, geo_id, "scale", "1.0")
//
//     // Ticker 
//     temp_id = 4
//     hub.AddTemplate(temp_id, "up", "", "up")
//
//     geo_id = 0
//     hub.AddGeometry(temp_id, geo_id, 0, "rect")
//     hub.AddAttr(temp_id, geo_id, "color", "2 132 130 255")
//
//     geo_id = 1
//     hub.AddGeometry(temp_id, geo_id, 0, "rect")
//     hub.AddAttr(temp_id, geo_id, "color", "5 68 94 255")
//     hub.AddAttr(temp_id, geo_id, "rounding", "50")
//
//     geo_id = 2
//     hub.AddGeometry(temp_id, geo_id, 0, "text")
//     hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
//     hub.AddAttr(temp_id, geo_id, "scale", "1.0")
// }
