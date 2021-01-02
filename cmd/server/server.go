package main

import (
	"github.com/emilgelman/gocrud/pkg/db"
	"github.com/emilgelman/gocrud/pkg/domain"
	"github.com/emilgelman/gocrud/pkg/transport"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initialize()
}

func initialize() {
	var database db.Db
	database = db.NewMemory()
	loadData(database)
	http := transport.NewHTTPServer(&database)
	grpc:= transport.NewGRPCServer(&database)
	transports:= []transport.Transport{http, grpc}
	for _,t := range transports {
		go func(t transport.Transport) {
			t.Serve()
		}(t)
	}

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	// Waiting for SIGINT (pkill -2)
	<-stop
	log.Printf("Exiting")
}

func loadData(database db.Db) {
	database.Create("1", domain.Article{Id: "1", Title: "title1", Content: "content1"})
	database.Create("2", domain.Article{Id: "2", Title: "title2", Content: "content2"})
}