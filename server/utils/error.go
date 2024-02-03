package utils

const (
	MSG_OK        = "0"
	MSG_DBERR     = "7001"
	MSG_UNKNOWERR = "7005"
)

var recodeText = map[string]string{
	MSG_DBERR:     "数据库错误",
	MSG_OK:        "成功",
	MSG_UNKNOWERR: "未知错误",
}

func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[MSG_UNKNOWERR]
}
