package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/maxwell92/log"
	"io/ioutil"
	"net/http"
)

var logger = log.Log

type ResData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HttpTool struct{}

func (h *HttpTool) WriteOk(c *gin.Context, data interface{}) {
	res := ResData{
		Code:    0,
		Message: "ok",
		Data:    data,
	}
	c.JSON(http.StatusOK, res)
}

func (h *HttpTool) WriteErr(c *gin.Context, err interface{}) {
	res := ResData{
		Code:    1,
		Message: "find a error",
		Data:    err,
	}
	c.JSON(http.StatusOK, res)
}

func (h *HttpTool) JsonUnmarshal(c *gin.Context, options interface{}) {
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
