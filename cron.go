package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	_ = c.AddFunc("0/5 * * * * *", func() {
		log.Println("task.GetData()...")
	})

	c.Start()
	t1 := time.NewTimer(time.Second * 10)
	select {
	case <-t1.C:
		t1.Reset(time.Minute * 10)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT)
	<-quit
	fmt.Println("Ready.............")
	time.Sleep(6 * time.Second)
	log.Println("Shutdown Server ...")
}
