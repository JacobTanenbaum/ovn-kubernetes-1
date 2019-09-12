package cluster

import (
	"fmt"
	"net"
	"net/http"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	utilwait "k8s.io/apimachinery/pkg/util/wait"

	"github.com/prometheus/client_golang/prometheus"
)

func startMetricsServer(bindAddress, proxyMode string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/proxyMode", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", proxyMode)
	})
	mux.Handle("/metrics", prometheus.Handler())
	go utilwait.Untl(func() {
		err := http.ListenAndServe(bindAddress, mux)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("starting metrics server failed: %v", err))
		}
	}, 5*time.Second, utilwait.NeverStop)

}
