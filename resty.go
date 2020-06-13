package gluaresty

import (
	"github.com/go-resty/resty/v2"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func Preload(L *lua.LState) {
	L.PreloadModule("resty", func(L *lua.LState) int {
		L.Push(luar.New(L, map[string]interface{}{
			"R": resty.New().R,
			"P": Parse,
		}))
		return 1
	})
}

type Response struct {
	Code   int
	Status string
	Body   string
}

func Parse(resp *resty.Response) *Response {
	if resp == nil {
		return nil
	}

	return &Response{
		Code:   resp.StatusCode(),
		Status: resp.Status(),
		Body:   string(resp.Body()),
	}
}
