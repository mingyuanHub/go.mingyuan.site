package common

import (
	"encoding/base64"
	"net/url"
	"strings"
)

//urlEncode
func UrlEscape(str string) string{
	return url.QueryEscape(str)
}


//urlUncode
func UrlUnescape(str string) string {
	unescape, err := url.QueryUnescape(str)
	if err != nil {
		return ""
	}
	return unescape
}

//base64解码
func Base64URLDecode(data string) ([]byte, error){
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing)
	res, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//安全Base64编码
func Base64UrlSafeEncode(source []byte) string {
	// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
	bytearr := base64.StdEncoding.EncodeToString(source)
	safeurl := strings.Replace(string(bytearr), "/", "_", -1)
	safeurl = strings.Replace(safeurl, "+", "-", -1)
	safeurl = strings.Replace(safeurl, "=", "", -1)
	return safeurl
}