package shelf

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"strings"
)

func decode(encoding string, data []byte) ([]byte, error) {
	encoding = strings.ToLower(encoding)
	switch encoding {
	case "gbk":
		return simplifiedchinese.GBK.NewDecoder().Bytes(data)
	case "gb18030":
		return simplifiedchinese.GB18030.NewDecoder().Bytes(data)
	case "gb2312":
		return simplifiedchinese.HZGB2312.NewDecoder().Bytes(data)
	case "utf-8", "":
		return data, nil
	default:
		return data, NewUnsupportedEncodingError(nil, encoding)
	}
}

func encode(encoding string, data []byte) ([]byte, error) {
	encoding = strings.ToLower(encoding)
	switch encoding {
	case "gbk":
		return simplifiedchinese.GBK.NewEncoder().Bytes(data)
	case "gb18030":
		return simplifiedchinese.GB18030.NewEncoder().Bytes(data)
	case "gb2312":
		return simplifiedchinese.HZGB2312.NewEncoder().Bytes(data)
	case "utf-8", "":
		return data, nil
	default:
		return data, NewUnsupportedEncodingError(nil, encoding)
	}
}
