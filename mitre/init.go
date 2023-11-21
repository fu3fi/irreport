package mitre

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

var VIRUSTOTAL_API_KEY string = "0a8f1a3ad8bd7a41f95a2e4719a96aa6e20911d176dca73deeaccaf424d8af03"
var VIRUSTOTAL_API_LIMIT_PER_SECOND = 4
var TRANSLATER_URL = "http://192.168.169.231:5667"

func init() {
	checkData := map[string]string{
		"q":      "Hello World",
		"source": "en",
		"target": "ru",
	}

	checkClient := resty.New()
	check, _ := checkClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(First(json.Marshal(checkData))).
		Post(TRANSLATER_URL + "/translate")

	if gjson.Get(check.String(), "translatedText").String() != "Привет мир" {
		fmt.Println("Translator is not working!")
		os.Exit(1)
	}
}
