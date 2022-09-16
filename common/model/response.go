package model

import "github.com/cloudwego/hertz/pkg/app"

/**
 * @Author: zze
 * @Date: 2022/9/15 14:04
 * @Desc: 响应模型
 */

type Response[T any] struct {
	C    *app.RequestContext `json:"-"`
	Code int                 `json:"code"`
	Msg  string              `json:"msg,omitempty"`
	Data *T                  `json:"data,omitempty"`
}

func (resp *Response[T]) JsonSelf() {
	resp.C.JSON(200, resp)
}

func (resp *Response[T]) JsonErr(err error) {
	resp.JsonErrWithCode(0, err)
}

func (resp *Response[T]) JsonOk(msg string, data *T) {
	resp.JsonOkWithCode(1, msg, data)
}

func (resp *Response[T]) JsonOkWithCode(code int, msg string, data *T) {
	resp.Code = code
	resp.Msg = msg
	resp.Data = data
	resp.JsonSelf()
}

func (resp *Response[T]) JsonErrWithCode(code int, err error) {
	resp.Code = code
	resp.Msg = err.Error()
	resp.Data = nil
	resp.JsonSelf()
}
