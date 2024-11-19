package mtl

import (
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Register *prometheus.Registry

func InitMetrics(serviceName, metricsPort, registerAddr string) {
	Register = prometheus.NewRegistry()
	Register.MustRegister(collectors.NewGoCollector())
	Register.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	r, _ := consul.NewConsulRegister(registerAddr)
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)
	registerInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}
	_ = r.Register(registerInfo)
	server.RegisterShutdownHook(func() {
		r.Deregister(registerInfo)
	})
	http.Handle("/metrics", promhttp.HandlerFor(Register, promhttp.HandlerOpts{}))
	go http.ListenAndServe(metricsPort, nil)
}
