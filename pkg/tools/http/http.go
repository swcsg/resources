package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxwell92/log"
)

var logger = log.Log

// ResData 统一的请求返回结构体
type ResData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//HTTPTool 结构体
type HTTPTool struct{}

// WriteOk 统一正常返回
func (h *HTTPTool) WriteOk(c *gin.Context, data interface{}) {
	res := ResData{
		Code:    0,
		Message: "ok",
		Data:    data,
	}
	c.JSON(http.StatusOK, res)
}

// WriteErr 统一错误返回
func (h *HTTPTool) WriteErr(c *gin.Context, err interface{}) {
	res := ResData{
		Code:    1,
		Message: "find a error",
		Data:    err,
	}
	c.JSON(http.StatusOK, res)
}

// JSONUnmarshal json序列化
func (h *HTTPTool) JSONUnmarshal(c *gin.Context, options interface{}) {
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Errorf("Tool GetOptionsFromBody ReadAll Error: error=%s, r.Body=%p", err, &c.Request.Body)
	}
	if len(reqBody) != 0 {
		err = json.Unmarshal(reqBody, options)
		if err != nil {
			logger.Errorf("Tool GetOptionsFromBody JSON Unmarshal Error: error=%s, reqBody=%p", err, reqBody)
		}
	}
}
