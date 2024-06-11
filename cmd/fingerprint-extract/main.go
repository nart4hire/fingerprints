package main

import (
	"encoding/json"
	"os"

	"github.com/nart4hire/fingerprints/lib/extraction"
	"github.com/nart4hire/fingerprints/lib/helpers"
)

func main() {
	path := os.Args[1]
	_, m := helpers.LoadImage(path)
	result := extraction.DetectionResult(m)
	d, _ := json.Marshal(result)
	println(string(d))
}
