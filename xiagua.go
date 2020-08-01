package xigua_watermark_remover

import "fmt"

func WatermarkRemover(url string)(string, error)  {

	itemId, err := GetItemId(url)

	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	videoId, err := GetVideoId(itemId)

	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	videoLink, err := getVideoDownloadLink(itemId, videoId)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	return videoLink, err
}
