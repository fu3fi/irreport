package mitre

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

func translate(text string, src string, dst string) string {
	client := resty.New()
	if src == "" {
		src = "en"
	}
	if dst == "" {
		dst = "ru"
	}

	data := map[string]string{
		"q":      text,
		"source": src,
		"target": dst,
	}
	dataJ, _ := json.Marshal(data)
	resp, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(dataJ)).
		Post(TRANSLATER_URL + "/translate")

	return gjson.Get(resp.String(), "translatedText").String()
}
