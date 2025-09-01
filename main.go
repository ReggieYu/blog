package main

import (
	"blog/app/router"
	"blog/config/log"
	"blog/config/toml"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	log.InitLogger(toml.GetConfig().Log.Path, toml.GetConfig().Log.Level)
	log.Logger.Info("this is the first log")
	log.Logger.Info("config", log.Any("config", toml.GetConfig()))

	r := gin.Default()
	r.Use(router.Cors())
	r.Use(router.Recovery)
	router.InitRouter(r)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Logger.Error("server error", log.Any("server error", err))
	}
}
