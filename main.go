package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"github.com/gin-gonic/gin"
	"goCart/models"
	"goCart/pkg/setting"
	"goCart/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	setting.Setup()
	models.Setup()
}
func main() {
	gob.Register(map[interface{}]interface{}{})
	gob.Register(map[string]string{})
	gob.Register(models.Admin{})
	gob.Register(models.User{})
	gin.SetMode(setting.ServerSetting.RunMode)
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	//s.ListenAndServe()
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
