package main

import (
	"log"
	"os"

	"github.com/NihilBabu/micro/handler"
	"github.com/NihilBabu/micro/server"
	"github.com/NihilBabu/micro/storage/mysql"
	"github.com/gorilla/mux"
)

var (
	CertFile    = os.Getenv("CERT_FILE")
	KeyFile     = os.Getenv("KEY_FILE")
	ServiceAddr = os.Getenv("SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout, "micro-app ", log.LstdFlags|log.Lshortfile)

	db, err := mysql.New("root", "password", "go","127.0.0.1:3306")
	if err != nil {
		logger.Fatalln(err)
	}

	h := handler.New(logger, db)
	mux := mux.NewRouter()
	h.SetupRoutes(mux)

	srv := server.New(mux, ServiceAddr)
	logger.Printf("Server is starting up in %v\n", ServiceAddr)
	err = srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
