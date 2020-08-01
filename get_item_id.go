package xigua_watermark_remover

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func GetItemId(url string) (string, error)  {

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var lastUrlQuery string


	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {

		if len(via) > 10 {
			return errors.New("too many redirects")
		}
		//lastUrlQuery = req.URL.RequestURI()

		lastUrlQuery = req.URL.String()
		return nil
	}

	response, err := client.Do(request)
	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
		return "", err
	}


	itemId := ExtractItemId(lastUrlQuery)


	return itemId, nil

}
