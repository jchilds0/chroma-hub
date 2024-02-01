package main

import (
	"chroma-hub/db"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

var port = flag.Int("port", 9000, "graphics hub port")
var fileName = flag.String("file", "", "graphics hub archive")
var hub = db.NewDataBase()

func main() {
    flag.Parse()

    ln, err := net.Listen("tcp", ":" + strconv.Itoa(*port))
    if err != nil {
        log.Fatalf("Error creating server (%s)", err)
    }

    if *fileName != "" {
        err = ImportArchive(*fileName)
        if err != nil {
            log.Printf("Error importing archive (%s)", err)
        }
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Printf("Error accepting connection (%s)", err)
        }

        handleConnection(conn)
    }
}

/*
    Send graphics hub to client using the following grammar

    S -> {'num_temp': num, 'templates': [T]}
    T -> {'id': num, 'num_geo': num, 'geometry': [G]} | T, T 
    G -> {'id': num, 'type': string, 'parent': num, 'attr': [A]} | G, G 
    A -> {'name': string, 'value': string} | A, A

 */

func handleConnection(conn net.Conn) {
    _, err := conn.Write([]byte(hub.String()))
    if err != nil {
        log.Printf("Error sending hub (%s)", err)
    }
}

func ImportArchive(fileName string) error {
    buf, err := os.ReadFile(fileName)
    if err != nil {
        return err
    }

    err = json.Unmarshal(buf, hub)
    if err != nil {
        return err
    }

    fmt.Printf("Imported Hub\n")
    return nil
}

func ExportArchive(fileName string) {
    file, err := os.Create(fileName)
    if err != nil {
        log.Fatalf("Couldn't open file (%s)", err)
    }
    defer file.Close()

    buf, err := json.Marshal(hub)
    if err != nil {
        log.Printf("Error encoding hub (%s)", err)
    }

    _, err = file.Write(buf)
    if err != nil {
        log.Printf("Error writing hub (%s)", err)
    }
    
    fmt.Printf("Exported Hub\n")
}

func BaseArchive() {
    // Teal Box
    temp_id := 0
    hub.AddTemplate(temp_id, "left_to_right", "", "left_to_right")

    geo_id := 0
    hub.AddGeometry(temp_id, geo_id, 0, "rect")
    hub.AddAttr(temp_id, geo_id, "color", "2 132 130 255")
    hub.AddAttr(temp_id, geo_id, "rounding", "20")

    geo_id = 1
    hub.AddGeometry(temp_id, geo_id, 0, "circle")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")

    geo_id = 2
    hub.AddGeometry(temp_id, geo_id, 0, "text")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
    hub.AddAttr(temp_id, geo_id, "scale", "1.0")

    geo_id = 3
    hub.AddGeometry(temp_id, geo_id, 0, "text")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
    hub.AddAttr(temp_id, geo_id, "scale", "1.0")

    // Clock Box
    temp_id = 1
    hub.AddTemplate(temp_id, "left_to_right", "clock_tick", "left_to_right")

    // bg
    geo_id = 0
    hub.AddGeometry(temp_id, geo_id, 0, "rect")
    hub.AddAttr(temp_id, geo_id, "color", "5 68 94 255")
    hub.AddAttr(temp_id, geo_id, "rounding", "10")

    // circle
    geo_id = 1
    hub.AddGeometry(temp_id, geo_id, 0, "circle")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")

    // left split
    geo_id = 2
    hub.AddGeometry(temp_id, geo_id, 0, "rect")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")

    // team 1
    geo_id = 3
    hub.AddGeometry(temp_id, geo_id, 2, "text")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
    hub.AddAttr(temp_id, geo_id, "scale", "1.0")
    
    // score 1
    geo_id = 4
    hub.AddGeometry(temp_id, geo_id, 2, "text")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
    hub.AddAttr(temp_id, geo_id, "scale", "1.0")

    // mid split
    geo_id = 5
    hub.AddGeometry(temp_id, geo_id, 2, "rect")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")

    // team 2
    geo_id = 6
    hub.AddGeometry(temp_id, geo_id, 5, "text")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
    hub.AddAttr(temp_id, geo_id, "scale", "1.0")

    // score 2
    geo_id = 7
    hub.AddGeometry(temp_id, geo_id, 5, "text")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
    hub.AddAttr(temp_id, geo_id, "scale", "1.0")

    // right split
    geo_id = 8
    hub.AddGeometry(temp_id, geo_id, 5, "rect")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")

    // clock
    geo_id = 9
    hub.AddGeometry(temp_id, geo_id, 8, "text")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
    hub.AddAttr(temp_id, geo_id, "scale", "1.0")

    // clock text
    geo_id = 10
    hub.AddGeometry(temp_id, geo_id, 8, "text")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
    hub.AddAttr(temp_id, geo_id, "scale", "1.0")

    // White circle
    temp_id = 2
    hub.AddTemplate(temp_id, "", "", "")

    geo_id = 0
    hub.AddGeometry(temp_id, geo_id, 0, "circle")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")

    // Graph 
    temp_id = 3
    hub.AddTemplate(temp_id, "left_to_right", "", "left_to_right")

    geo_id = 0
    hub.AddGeometry(temp_id, geo_id, 0, "rect")
    hub.AddAttr(temp_id, geo_id, "color", "2 132 130 255")
    hub.AddAttr(temp_id, geo_id, "rounding", "20")

    geo_id = 1
    hub.AddGeometry(temp_id, geo_id, 0, "graph")
    hub.AddAttr(temp_id, geo_id, "graph_type", "step")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")

    geo_id = 2
    hub.AddGeometry(temp_id, geo_id, 0, "text")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
    hub.AddAttr(temp_id, geo_id, "scale", "1.0")

    // Ticker 
    temp_id = 4
    hub.AddTemplate(temp_id, "up", "", "up")

    geo_id = 0
    hub.AddGeometry(temp_id, geo_id, 0, "rect")
    hub.AddAttr(temp_id, geo_id, "color", "2 132 130 255")

    geo_id = 1
    hub.AddGeometry(temp_id, geo_id, 0, "rect")
    hub.AddAttr(temp_id, geo_id, "color", "5 68 94 255")
    hub.AddAttr(temp_id, geo_id, "rounding", "50")

    geo_id = 2
    hub.AddGeometry(temp_id, geo_id, 0, "text")
    hub.AddAttr(temp_id, geo_id, "color", "255 255 255 255")
    hub.AddAttr(temp_id, geo_id, "scale", "1.0")
}
