package server

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/sstallion/go-hid"
	"lianlinux/core"
	"net/http"
	"strconv"
)

type rootJSON struct {
	SupportedDevices []string          `json:"SupportedDevices"`
	Devices          []*hid.DeviceInfo `json:"Devices"`
}

type responseJSON struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
}

func root(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rootJSON{
		SupportedDevices: []string{"a100"},
		Devices:          core.Devs,
	})
}

func lightMode(c *gin.Context) {
	deviceNumber, err := strconv.Atoi(c.Query("dev"))
	if err != nil {
		c.IndentedJSON(http.StatusOK, responseJSON{
			Status:  "error",
			Message: fmt.Sprintf("%v", err),
		})
		return
	}
	log.Info(deviceNumber)

	newMode := c.Query("mode")
	if newMode != "" {
		core.SetLightMode(*core.Devs[deviceNumber], newMode)
		c.IndentedJSON(http.StatusOK, responseJSON{
			Status:  "ok",
			Message: fmt.Sprintf("Set device %d mode to %s", deviceNumber, newMode),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, responseJSON{
		Status:  "error",
		Message: "Mode is empty",
	})
}

func Listen(port int) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", root)
	router.GET("/mode", lightMode)

	log.Infof("Listening on port %d", port)
	err := router.Run("localhost:" + strconv.Itoa(port))
	if err != nil {
		return
	}
}
