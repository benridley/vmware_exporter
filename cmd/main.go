package main

import (
	"flag"
	"net/http"

	vmwexporter "github.com/benridley/vmware_exporter/internal"
	"github.com/benridley/vmware_exporter/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

func main() {
	var (
		listenAddress = flag.String("web.listen-address", ":9536", "Address to listen on for web interface and telemetry.")
		metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	)
	flag.Parse()

	config, err := util.GetVSphereConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Connecting to vCenter")

	e, err := vmwexporter.NewVmwareExporter(config)
	if err != nil {
		log.Fatal(err)
	}
	prometheus.MustRegister(e)

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
			<head><title>VMWare Exporter</title></head>
			<body>
			<h1>VMWare Exporter</h1>
			<p><a href='` + *metricsPath + `'>Metrics</a></p>
			</body>
			</html>`))
	})
	log.Info("Listening on address:port => ", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
