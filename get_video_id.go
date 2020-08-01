package xigua_watermark_remover

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type JsonData struct {
	Data Data `json:"data"`
}

type Data struct {
	VideoId string `json:"video_id"`
}


func GetVideoId(itemId string) (string, error)  {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://m.365yg.com/i" + itemId + "/info/", nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	response, err := client.Do(request)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
		return "", err
	}


	jsonByteData := []byte(string(body))

	jsonData := JsonData{}
	json.Unmarshal(jsonByteData, &jsonData)

	return jsonData.Data.VideoId, nil
}
