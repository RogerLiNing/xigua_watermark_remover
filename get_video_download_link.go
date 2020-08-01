package xigua_watermark_remover

import (
	"errors"
	"log"
	"net/http"
)

func getVideoDownloadLink(itemId string, videoId string) (string, error) {
	client := &http.Client{}
	url := "https://api.huoshan.com/hotsoon/item/video/_source/?video_id="+videoId+"&line=0&app_id=0&vquality=normal&watermark=0&sf=5&item_id="+itemId

	request, err := http.NewRequest("GET", url, nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("accept-encoding", "gzip, deflate, br")
	request.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	//var lastUrlQuery string
	var lastUrlQuery string

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {

		if len(via) > 10 {
			return errors.New("too many redirects")
		}
		lastUrlQuery = req.URL.String()

		return nil
	}

	response, err := client.Do(request)
	defer response.Body.Close()


	return lastUrlQuery, nil
}
