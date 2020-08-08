package main

import (
	"log"
	"net/http"
	"os"

	"github.com/NihilBabu/micro/homepage"
	"github.com/NihilBabu/micro/server"
	"github.com/NihilBabu/micro/storage/mysql"
)

var (
	CertFile    = os.Getenv("CERT_FILE")
	KeyFile     = os.Getenv("KEY_FILE")
	ServiceAddr = os.Getenv("SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout, "micro-app ", log.LstdFlags|log.Lshortfile)

	db, err := mysql.New("hi", "hlo", "hey")
	if err != nil {
		logger.Fatalln(err)
	}

	h := homepage.New(logger, db)
	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux, ServiceAddr)
	logger.Printf("Server is starting up in %v\n", ServiceAddr)
	err = srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
