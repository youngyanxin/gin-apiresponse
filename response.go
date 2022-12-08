package gin_apiresponse

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Res struct {
}

//func New() *Res {
//	return &Res{}
//}

// api接口响应内容
type Response struct {
	Code      int         `json:"code"`       //业务状态值
	Msg       string      `json:"msg"`        //业务信息
	Data      interface{} `json:"data"`       //数据
	RequestID string      `json:"request_id"` //请求id
	Success   bool        `json:"success"`    //接口是否请求成功,true成功、false失败
	Redirect  string      `json:"redirect"`   //重定向连接
}

func MsgReturn(msgCode int, data interface{}) *Response {
	result := &Response{
		Code:      msgCode,
		Msg:       GetCodeMsg(msgCode),
		Data:      data,
		RequestID: CreateRequestID(),
		Success:   true,
	}
	return result
}

func CreateRequestID() string {
	u1 := uuid.NewV4()
	fmt.Println(u1)
	str := fmt.Sprintf("%x", u1)
	fmt.Println(str)
	return str
}

func JsonExit(ctx *gin.Context, res *Response) {
	ctx.JSON(200, res)
	return
}

// 参数错误
func ParamsErr(ctx *gin.Context, msgCode int) {
	result := MsgReturn(msgCode, "")
	JsonExit(ctx, result)
}

// 成功
func Success(ctx *gin.Context, msgCode int) {
	result := MsgReturn(msgCode, "")
	JsonExit(ctx, result)
}

// 成功返回数据
func SuccessWithData(ctx *gin.Context, msgCode int, data interface{}) {
	result := MsgReturn(msgCode, data)
	JsonExit(ctx, result)
}

// 失败
func Fail(ctx *gin.Context, msgCode int, err error) {
	result := MsgReturn(msgCode, err.Error())
	JsonExit(ctx, result)
}

// 404 未找到
func NotFound(ctx *gin.Context) {
	result := MsgReturn(404, "")
	JsonExit(ctx, result)
}

// 认证失败
func AuthErr(ctx *gin.Context) {
	result := MsgReturn(-1, "")
	JsonExit(ctx, result)
}
