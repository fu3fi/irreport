package main

import (
	"encoding/json"
	"fmt"
	"fu3fi/cir/work/irreport/mitre"
	"os"
)

func main() {
	result := mitre.Miner(".")

	resultFile, _ := os.Create("./result.fu3fi.txt")
	defer resultFile.Close()

	resultFile.WriteString(fmt.Sprintf("%s", mitre.First(json.MarshalIndent(result, "", "    "))))
}
