package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"judgeMore/biz/model/model"
	"judgeMore/pkg/errno"
	"strconv"
	"time"
)

type Base struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type Response struct {
	Base `json:"base"`
}

// 数据类型多样-用interface
func SendResponse(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, data)
}

func BuildBaseResp(err errno.ErrNo) *model.BaseResp {
	return &model.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

func SendFailResponse(c *app.RequestContext, err errno.ErrNo) {
	baseResp := Base{
		Code: err.ErrorCode,
		Msg:  err.Error(),
	}
	response := Response{
		Base: baseResp,
	}
	c.JSON(consts.StatusBadRequest, response)
}

func ChangeFormat(timeStr string) string {
	timestamp, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return ""
	}
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
