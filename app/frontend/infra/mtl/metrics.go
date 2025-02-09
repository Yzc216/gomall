package mtl

import (
	"context"
	"fmt"
	"github.com/Yzc216/gomall/app/frontend/conf"
	frontendutils "github.com/Yzc216/gomall/app/frontend/utils"
	"github.com/Yzc216/gomall/common/utils"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/route"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
)

var Registry *prometheus.Registry

func InitMetric() route.CtxCallback {
	//prometheus注册表
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	//初始化hertz consul注册
	config := consulapi.DefaultConfig()
	config.Address = conf.GetConf().Hertz.RegistryAddr
	consulClient, _ := consulapi.NewClient(config)
	r := consul.NewConsulRegister(consulClient)

	//解析监控ip和端口
	localIp := utils.MustGetLocalIPv4()
	ip, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s%s", localIp, conf.GetConf().Hertz.MetricsPort))
	if err != nil {
		hlog.Error(err)
	}

	registryInfo := &registry.Info{
		Addr:        ip,
		ServiceName: "prometheus",
		Weight:      1,
		Tags:        map[string]string{"service": frontendutils.ServiceName},
	}

	err = r.Register(registryInfo)
	if err != nil {
		hlog.Error(err)
	}

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(conf.GetConf().Hertz.MetricsPort, nil)

	return func(ctx context.Context) {
		r.Deregister(registryInfo)
	}
}
