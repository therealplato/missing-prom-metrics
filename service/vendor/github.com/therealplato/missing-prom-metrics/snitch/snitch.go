package snitch

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Snitch struct{}

func (s *Snitch) Start() {
	var (
		err error
		lst net.Listener
		mux = http.NewServeMux()
		srv = http.Server{Addr: "0.0.0.0:9009", Handler: mux}
	)

	if lst, err = net.Listen("tcp", "0.0.0.0:9009"); err != nil {
		log.Fatal(fmt.Sprintf("snitch failed to listen: %s", err))
	}

	mux.Handle("/metrics", promhttp.Handler())

	log.Println("snitch starts")
	srv.Serve(lst)
}

func (s *Snitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
