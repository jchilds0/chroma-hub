/*

   Format json archive - :%!python -m json.tool

*/

package main

import (
	"flag"
	"log"

	"github.com/jchilds0/chroma-hub/chroma_hub"
)

var port = flag.Int("port", 9000, "graphics hub port")
var fileName = flag.String("file", "", "graphics hub archive")
var count = flag.Int("count", -1, "number of graphics hub requests to send (-1 for no upper limit)")

func main() {
    flag.Parse()

    if *fileName == "" {
        log.Fatal("Missing filename")
    }

    chroma_hub.StartHub(*port, *count, *fileName)
}

