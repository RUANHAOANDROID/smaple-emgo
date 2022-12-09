package api

import (
	"emcs-relay-go/db"
	"emcs-relay-go/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

func HandlerConfigManager(r *gin.RouterGroup) {

	r.POST("/get", getConfig())
	r.POST("/getConfig", getConfig())
	r.POST("/getMyConfig", getMyConfig())
	r.POST("/saveMyConfig", saveMyConfig())
}

func getMyConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := db.MyConfig{}
		err := db.GetMyConfig(&config)
		if err != nil {
			utils.Log.Info(err.Error())
			c.JSON(http.StatusOK, ResponseError(err.Error()))
			return
		}
		c.JSON(http.StatusOK, ResponseSuccess(config.JsonContent))
	}
}

// {"data":{"serialNum":"3Y32225000630212"},"sign":"E430EA624BDC994DFA29BF20DF7FD244","timestamp":1670554957167}

func getConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := WebGetConfig{}
		err := c.BindJSON(&request)
		if err != nil {
			ResError(c, "参数错误")
			return
		}
		wgc := request
		result, err := rpcGetConf(wgc.ConfigUrl, wgc.Id, time.Now().Local().Unix()/100)
		if err != nil {
			ResError(c, "获取配置失败"+err.Error())
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
func saveMyConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)

		if err != nil {
			utils.Log.Error(err.Error())
			c.JSON(http.StatusOK, ResponseError(err.Error()))
			return
		}
		jsonString := string(body)
		err = db.SaveMyConfig(&db.MyConfig{JsonContent: jsonString})
		if err != nil {
			utils.Log.Error(err.Error())
			c.JSON(http.StatusOK, ResponseError(err.Error()))
			return
		}
		c.JSON(http.StatusOK, ResponseSuccess("ok"))
		err = db.DeleteMyOtherConfig()
	}
}
