package xigua_watermark_remover

import "regexp"

func ExtractItemId(url string) string  {
	// 提取item_id，目前看是19位整数，安全起见，取15位或者15位以上的整数
	//解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile(`(\d{15,})`)

	//根据规则提取关键信息
	result := reg.FindAllStringSubmatch(url,-1)
	var itemId string

	if len(result) > 0 {
		itemId = result[0][1]
	}

	return itemId
}
