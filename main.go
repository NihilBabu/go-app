package main

import (
	"context"
	"flag"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/NihilBabu/go-app/handler"
	"github.com/NihilBabu/go-app/storage/mysql"
	"github.com/NihilBabu/go-app/util"
)

func init() {

}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "wait for existing connections to finish")
	flag.Parse()

	svc, err := mysql.New("root", "pass", "go")
	if err != nil {
		log.Fatalln(err)
	}
	defer svc.Close()

	util.InitialMigrate(svc)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 120,
		Handler:      handler.New(svc),
	}

	go func() {
		log.Println("app is starting up")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
