package main

import (
	"log"

	"github.com/Marcus-hayes/media-lookup/cmd"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
}

func main() {
	cmd.Execute()
}
