package gin_apiresponse

var CodeMsg = map[int]string{

	0:  "未知错误%%0",
	-1: "认证失败%%0",
	1:  "操作成功%%0",
	2:  "操作失败%%0",

	9999: "签名错误",

	404: "接口不存在%%0",
	500: "服务器内部错误%%0",
}

func GetCodeMsg(code int) string {
	if msg, ok := CodeMsg[code]; ok {
		return msg
	}
	return CodeMsg[0]
}
