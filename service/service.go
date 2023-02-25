package service

import (
	"log"
	"time"

	"github.com/asim/go-micro/v3/registry"
	"github.com/opentracing/opentracing-go"

	"github.com/asim/go-micro/v3"
	"github.com/go-micro/plugins/v3/registry/consul"
	"github.com/jary-287/gopass-common/common"
	"github.com/jary-287/gopass-svc/proto/svc"
)

var Client svc.SvcService

func init() {
	consulRegister := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{"192.168.0.19:8500"}
		o.Timeout = 20 * time.Second
	})
	// 链路追踪
	t, io, err := common.NewTracer("base", "192.168.0.102:9999")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	service := micro.NewService()

	service.Init(micro.Registry(consulRegister))
	Client = svc.NewSvcService("service.svc", service.Client())

}
