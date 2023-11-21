package mitre

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Miner(path string) map[string]interface{} {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	result := make(map[string]interface{})

	for _, file := range files {

		if file.Name() == filepath.Base(os.Args[0]) || file.Name() == "result.fu3fi.txt" {
			continue
		}

		fmt.Println("Start for", file.Name())

		cursor, err := os.Open(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close()

		if file.IsDir() {
			continue
		}

		result[file.Name()] = make(map[string]interface{})

		result[file.Name()].(map[string]interface{})["hashes"] = map[string]string{
			"md5":    fmt.Sprintf("%x", First(hashSum(file.Name(), md5.New))),
			"sha1":   fmt.Sprintf("%x", First(hashSum(file.Name(), sha1.New))),
			"sha256": fmt.Sprintf("%x", First(hashSum(file.Name(), sha256.New))),
		}

		result[file.Name()].(map[string]interface{})["sandboxes"] = make(map[string]interface{})
		hash := First(hashSum(file.Name(), sha256.New))
		for sandboxName, info := range MitreByHash(fmt.Sprintf("%x", hash)) {
			result[file.Name()].(map[string]interface{})["sandboxes"].(map[string]interface{})[sandboxName] = info
		}

		time.Sleep(time.Second * time.Duration(60/VIRUSTOTAL_API_LIMIT_PER_SECOND+5))
	}

	return result
}
