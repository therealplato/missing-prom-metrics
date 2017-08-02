package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/therealplato/missing-prom-metrics/metrics"
	"github.com/therealplato/missing-prom-metrics/snitch"
)

func main() {
	http.HandleFunc("/", HelloServer)
	sn := &snitch.Snitch{}
	go sn.Start()
	log.Println("starting service http server...")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		metrics.PromRequestStatusIncoming.WithLabelValues("myservice", strconv.Itoa(http.StatusMethodNotAllowed)).Inc()
		log.Println("incremented http_status_incoming_total metric")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	metrics.PromRequestStatusIncoming.WithLabelValues("myservice", strconv.Itoa(http.StatusOK)).Inc()
	io.WriteString(w, "hello, world!\n")
}
