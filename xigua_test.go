package xigua_watermark_remover

import "testing"

func TestAvailableLink(t *testing.T)  {
	/*
		处理以下多种格式
		URL1：https://www.ixigua.com/6509242953687368199/?app=video_article
		URL2：https://m.ixigua.com/group/6792545386817913348
		URL3：https://v.ixigua.com/J2wEgNH/
		URL4: 其他域名，属于分享域名，最终都是跳转https://m.ixigua.com/group/6792545386817913348 这种格式，只需要获取ID就行
	*/
	link := "https://www.ixigua.com/6509242953687368199/?app=video_article"

	url, err := WatermarkRemover(link)
	t.Log("测试链接", link)

	if err != nil {
		t.Fail()
		t.Log("返回视频链接，测试通过")
	}

	if len(url) > 0 {
		t.Log("测试通过")
	}

}