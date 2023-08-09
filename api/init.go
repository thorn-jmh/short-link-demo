package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-svc-tpl/api/route"
	"net/http"
	"time"
)

type WebServerCfg struct {
	Port         int      `mapstructure:"Port"`
	WriteTimeout int      `mapstructure:"WriteTimeout"`
	ReadTimeout  int      `mapstructure:"ReadTimeout"`
	AllowOrigins []string `mapstructure:"AllowOrigins"`
}

func StartServer() error {
	var cfg WebServerCfg
	if err := viper.Sub("WebServer").UnmarshalExact(&cfg); err != nil {
		return err
	}

	e := gin.Default()
	e.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length"},
		AllowCredentials: true,
	}))
	route.SetupRouter(e.Group("/"))

	if viper.GetString("App.RunLevel") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        e,
		ReadTimeout:    time.Second * time.Duration(cfg.ReadTimeout),
		WriteTimeout:   time.Second * time.Duration(cfg.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
}
