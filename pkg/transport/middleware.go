package transport

import (
	"log"
	"net/http"
	"time"
)

func HandleWithLog(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		start:= time.Now()
		log.Printf("started %s %s", request.Method, request.URL.Path)
		defer log.Printf("ended %s %s, took %v millis", request.Method, request.URL.Path, time.Since(start).Milliseconds())
		f(writer,request)
	}
}

