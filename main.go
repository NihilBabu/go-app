package main

import (
	"context"
	"flag"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/NihilBabu/go-app/storage"
	"github.com/NihilBabu/go-app/storage/mysql"
	"github.com/NihilBabu/go-app/util"
	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	log.Println("Initializing app")
	//testing for database connection
	util.InitialMigrate()

	router = GetAllRoutes()

	mys := mysql.Mysql{}

	a, _ := storage.Service.Save(&mys, "hlo")

	log.Println(a)

	// a, _ := storage.Service.Save("hlo")

	// log.Println(a)
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// r := GetAllRoutes()

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	go func() {
		log.Println("app is starting up")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
