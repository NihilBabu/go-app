package main

import (
	"github.com/NihilBabu/micro/homepage"
	"github.com/NihilBabu/micro/server"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
)

var (
	CertFile    = os.Getenv("CERT_FILE")
	KeyFile     = os.Getenv("KEY_FILE")
	ServiceAddr = os.Getenv("SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout,"micro-app ",log.LstdFlags | log.Lshortfile )

	db,err := sqlx.Open("postgres","")
	if err != nil {
		logger.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		logger.Fatalln(err)
	}

	h := homepage.NewHandlers(logger,db)
	mux := http.NewServeMux()
	h.SetupRoutes(mux)


	srv := server.New(mux, ServiceAddr)
	logger.Printf("Server is starting up in %v\n",ServiceAddr)
	err = srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}

