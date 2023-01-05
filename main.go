package main

import (
	"fmt"
	"log"
	"net/http"
	redisdemo "redis-go-docker-demo/redis-demo"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Path("/ping").Methods(http.MethodGet).HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		rc := redisdemo.GetRedisClient()
		if pong, err := rc.GetPing(); err != nil {
			log.Println("Error while ping to redis server.", err)
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(fmt.Sprintf("%v", err)))
		} else {
			log.Println("Ping message from redis server.", pong)
			res.WriteHeader(http.StatusOK)
			res.Write([]byte("Hi there " + pong))
		}
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Panic(err)
	}

}
