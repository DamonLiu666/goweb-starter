package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type page struct {
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	List      interface{} `json:"list"`
}

func Error(c *gin.Context, code int, err error, msg string) {
	res := response{}
	if err != nil {
		res.Msg = err.Error()
	}
	if msg != "" {
		res.Msg = msg
	}

	res.Code = int32(code)
	c.Set("result", res)
	c.Set("status", code)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	res := response{}
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	res.Code = http.StatusOK
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, msg)
}
