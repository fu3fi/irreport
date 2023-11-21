package mitre

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

func MitreByHash(hash string) map[string]map[string]string {
	url := fmt.Sprintf("https://www.virustotal.com/api/v3/files/%s/behaviour_mitre_trees", hash)

	vtClient := resty.New()

	resp, _ := vtClient.R().
		SetHeader("x-apikey", VIRUSTOTAL_API_KEY).
		Get(url)

	sandboxes := gjson.Get(resp.String(), "data").Map()

	result := make(map[string]map[string]string)
	for sandboxName, data := range sandboxes {
		curs := data.Map()

		result[sandboxName] = make(map[string]string)
		for _, technique := range curs["tactics"].Array() {
			result[sandboxName]["Description"] = translate(technique.Map()["description"].String(), "", "")
			result[sandboxName]["Link"] = technique.Map()["link"].String()
			result[sandboxName]["Id"] = technique.Map()["id"].String()
		}
	}

	return result
}
