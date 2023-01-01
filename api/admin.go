package api

import (
	"emcs-relay-go/db"
	"emcs-relay-go/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerAdmin(r *gin.RouterGroup) {
	//r.Use(SignHandler())
	r.POST("/login", login())
	r.POST("/updatePassword", updatePassword())
	r.POST("/upgrade", upgrade())
}

func updatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		updUser := UpdateUser{}
		err := c.ShouldBind(&updUser)
		if err != nil {
			ResError(c, "参数错误")
			return
		}
		db.UpdatePwd(updUser.UserName, updUser.Password)
		if err != nil {
			ResError(c, "更改密码失败")
			return
		}
		c.JSON(http.StatusOK, ResponseSuccess("OK"))
	}
}
func login() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestUser := LoginRequest{}
		err := c.BindJSON(&requestUser)

		if err != nil {
			c.JSON(http.StatusOK, ResponseError("参数错误"))
			utils.Log.Error(err.Error())
			return
		}
		err = db.UserExists(requestUser.UserName)
		if err != nil {
			c.JSON(http.StatusOK, ResponseError("用户不存在"))
			return
		}
		err = db.UserVerify(requestUser.UserName, requestUser.PassWord)
		if err != nil {
			c.JSON(http.StatusOK, ResponseError("密码错误"))
			return
		}
		response := ResponseSuccess(any("ok"))
		fmt.Println(response)
		c.JSON(http.StatusOK, response)
	}
}
func upgrade() gin.HandlerFunc {
	return func(c *gin.Context) {
		updUser := UpdateUser{}
		err := c.ShouldBind(&updUser)
		if err != nil {
			ResError(c, "参数错误")
			return
		}
		db.UpdatePwd(updUser.UserName, updUser.Password)
		if err != nil {
			ResError(c, "更改密码失败")
			return
		}
		c.JSON(http.StatusOK, ResponseSuccess("OK"))
	}
}
