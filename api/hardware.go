package api

import (
	"emcs-relay-go/api/entity"
	"emcs-relay-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerHardwareManager(r *gin.RouterGroup) {
	r.POST("/info", hardwareInfo())
	r.POST("/cpu", cpu())
	r.POST("/memory", memory())
	r.POST("/disk", disk())
	r.POST("/storage", storage())
}
func hardwareInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		cpu := entity.Hardware{Name: "CPU", Total: "100", Used: "ss", Proportion: ""}
		memory := entity.Hardware{Name: "Memory", Total: "100", Used: "ss", Proportion: ""}
		utils.Log.Info(cpu)
		utils.Log.Info(memory)
		c.JSON(http.StatusOK, "result")
	}
}
func cpu() gin.HandlerFunc {
	return func(c *gin.Context) {
		cpu := entity.Hardware{Name: "CPU", Total: "100", Used: "ss", Proportion: ""}
		memory := entity.Hardware{Name: "Memory", Total: "100", Used: "ss", Proportion: ""}
		utils.Log.Info(cpu)
		utils.Log.Info(memory)
		c.JSON(http.StatusOK, "result")
	}
}
func memory() gin.HandlerFunc {
	return func(c *gin.Context) {
		cpu := entity.Hardware{Name: "CPU", Total: "100", Used: "ss", Proportion: ""}
		memory := entity.Hardware{Name: "Memory", Total: "100", Used: "ss", Proportion: ""}
		utils.Log.Info(cpu)
		utils.Log.Info(memory)
		c.JSON(http.StatusOK, "result")
	}
}
func disk() gin.HandlerFunc {
	return func(c *gin.Context) {
		cpu := entity.Hardware{Name: "CPU", Total: "100", Used: "ss", Proportion: ""}
		memory := entity.Hardware{Name: "Memory", Total: "100", Used: "ss", Proportion: ""}
		utils.Log.Info(cpu)
		utils.Log.Info(memory)
		c.JSON(http.StatusOK, "result")
	}
}
func storage() gin.HandlerFunc {
	return func(c *gin.Context) {
		cpu := entity.Hardware{Name: "CPU", Total: "100", Used: "ss", Proportion: ""}
		memory := entity.Hardware{Name: "Memory", Total: "100", Used: "ss", Proportion: ""}
		utils.Log.Info(cpu)
		utils.Log.Info(memory)
		c.JSON(http.StatusOK, "result")
	}
}
