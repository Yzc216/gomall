package mtl

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"
	"net/http"
)

var Registry *prometheus.Registry

func InitMetric(serviceName string, metricsPort string, registryAddr string) (registry.Registry, *registry.Info) {
	//初始化prometheus注册表
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	//初始化consul注册
	r, _ := consul.NewConsulRegister(registryAddr)
	//解析监控端口
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)

	//注册服务到 Consul
	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}

	if err := r.Register(registryInfo); err != nil {
		log.Printf("failed to register service: %v", err)
	}

	//注册服务注销钩子
	server.RegisterShutdownHook(func() {
		if err := r.Deregister(registryInfo); err != nil {
			log.Printf("Failed to deregister service: %v", err)
		}
	})

	//暴露 Prometheus 指标
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(metricsPort, nil) //nolint:errcheck

	return r, registryInfo
}
