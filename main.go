package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var listenPort = flag.String("listen.port",
	getEnv("LISTEN_PORT", "80"),
	"Listened port.")

func main() {
	//	package main, test action
	flag.Parse()
	startListen()
}

func startListen() {
	router := mux.NewRouter()
	router.HandleFunc("/helo", httpHeloHendler).Methods("GET")
	listenUrl := fmt.Sprintf("0.0.0.0:%s", *listenPort)
	log.Printf("Runing listener on %s", listenUrl)
	log.Fatal(http.ListenAndServe(listenUrl, router))
}

func getEnv(key string, defaultVal string) string {
	if envVal, ok := os.LookupEnv(key); ok {
		return envVal
	}
	return defaultVal
}

func httpHeloHendler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from host: %s", r.Host)
	metricLine := "Helo user! Code: 100"
	w.Write([]byte(metricLine))
}
