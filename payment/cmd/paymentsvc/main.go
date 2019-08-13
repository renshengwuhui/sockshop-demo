package main

import (
	"fmt"
	"github.com/ServiceComb/go-chassis/core/lager"
	"os"
	"os/signal"
	"syscall"
	//"code.huawei.com/cse/go-chassis"
	"github.com/ServiceComb/go-chassis"
	"github.com/ServiceComb/go-chassis/examples/payment"
	_ "github.com/ServiceComb/go-chassis/server/restful"
	_ "github.com/ServiceComb/go-chassis/third_party/forked/go-micro/server/highway"
	_ "github.com/ServiceComb/go-chassis/third_party/forked/go-micro/transport/tcp"
	"github.com/emicklei/go-restful/log"
)

const (
	ServiceName = "payment"
)

func main() {
	var (
		declineAmount = 1000
	)
	chassis.Init()
	// Mechanical stuff.
	errc := make(chan error)
	var logger = log.Logger
	var service payment.Service
	{
		service = payment.NewAuthorisationService(float32(declineAmount))
		service = payment.LoggingMiddleware(logger, service)

	}
	chassis.RegisterSchema("rest", service)
	chassis.Run()

	// Capture interrupts.
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()
	lager.Logger.Error("exit", <-errc)
}
