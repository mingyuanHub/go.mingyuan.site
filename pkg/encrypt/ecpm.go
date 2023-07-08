package encrypt

import (
	"strconv"

	"mingyuanHub/mingyuan.site/pkg/common"
)

//ecpm加密
func EcpmEncrypt(price interface{}, key string) string{
	var priceStr string
	if price == nil {
		return priceStr
	}

	switch price.(type) {
	case float64:
		ft := price.(float64)
		priceStr = common.Float64ToString(ft)
	case string:
		priceStr = price.(string)
	}

	return common.Base64UrlSafeEncode(AesEncrypt(priceStr, key))
}

//ecpm解密
func EcpmDecrypt(price string, key string) (string, error){
	_, err := strconv.ParseFloat(price, 64)
	if err == nil {
		return price, nil
	}
	priceByte, err := common.Base64URLDecode(price)
	if err != nil {
		return "", err
	}
	return AesDecrypt(priceByte, key)
}