package main

import (
	"fmt"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func heartbeat(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func version(ctx *gin.Context) {
	ctx.String(http.StatusOK, viper.GetString("server.version"))
}

func LoadConfiguration() {
	viper.SetConfigFile("local.yaml")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configuration")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config file failed to load: %s", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file has been changed:", e.Name)
	})
}

func main() {
	LoadConfiguration()

	server := gin.Default()
	server.GET("/heartbeat", heartbeat)
	server.GET("/version", version)

	address := viper.GetString("server.host.address")
	port := viper.GetInt("server.host.port")
	server.Run(fmt.Sprintf("%s:%d", address, port))
}
