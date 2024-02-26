package chroma_hub

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/jchilds0/chroma-hub/db"
)

var hub = db.NewDataBase()

func StartHub(port, count int, fileName string) {
    ln, err := net.Listen("tcp", ":" + strconv.Itoa(port))
    if err != nil {
        log.Fatalf("Error creating server (%s)", err)
    }
    defer ln.Close()

    err = ImportArchive(fileName)
    if err != nil {
        log.Printf("Error importing archive (%s)", err)
    }

    upperLimit := (count != -1)
    count = 2 * count

    for {
        if upperLimit && count == 0 {
            break
        }

        conn, err := ln.Accept()
        if err != nil {
            log.Printf("Error accepting connection (%s)", err)
        }

        handleConnection(conn)
        if upperLimit {
            count--
        }
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

