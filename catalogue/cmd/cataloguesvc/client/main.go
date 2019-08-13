package main

import (
	"github.com/ServiceComb/go-chassis/client/rest"
	"github.com/ServiceComb/go-chassis/core"
	"github.com/ServiceComb/go-chassis/core/loadbalance"
	"github.com/ServiceComb/go-chassis/core/registry"
	_ "github.com/ServiceComb/go-chassis/server/highway"
	_ "github.com/ServiceComb/go-chassis/server/restful"
	"github.com/ServiceComb/go-chassis/core/lager"
	"golang.org/x/net/context"
	"log"
	"sync"
	"time"
	"net/http"
	"github.com/ServiceComb/go-chassis"
)

var wg sync.WaitGroup
var success int
var fail int

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/path/to/conf/folder
func main() {
	err := chassis.Init();if err != nil {
		lager.Logger.Errorf(err, "Chassis init failed.")
	}
	registry.Enable()
	loadbalance.Enable()
	n := 1
	wg.Add(n)

	restinvoker := core.NewRestInvoker()
	for m := 0; m < n; m++ {
		go callrest(restinvoker)
	}
	wg.Wait()
	k := 10
	start := time.Now()
	lager.Logger.Debugf("begin to run benchmark of TPS, please wait for a while.......")
	wg.Add(n * k)
	restinvoker = core.NewRestInvoker()
	for m := 0; m < n; m++ {
		go bench(restinvoker, k)
	}
	wg.Wait()
	//use fasthttp to fix the bug of max connection per host.
	lager.Logger.Debugf("success: %d, fail: %d, spent: %fs, TPS: %f", success, fail, time.Now().Sub(start).Seconds(), float64(success)/time.Now().Sub(start).Seconds())

}

func callrest(invoker *core.RestInvoker) {
	defer wg.Done()
	/////////////////////////////////////////////////////////////////////////
	req, _ := rest.NewRequest("GET", "cse://Catalogue/health")
	req.SetHeader("Content-Type", "application/json")
	//use the invoker like http client.
	resp1, err := invoker.ContextDo(context.TODO(), req)
	if err != nil {
		lager.Logger.Errorf(err, "call request fail.")
		return
	}
	log.Printf("Rest Server health[Get] %s", string(resp1.ReadBody()))

	/////////////////////////////////////////////////////////////////////////
	req, _ = rest.NewRequest("GET", "cse://Catalogue//catalogue/size")
	req.SetHeader("Content-Type", "application/json")
	//use the invoker like http client.
	resp1, err = invoker.ContextDo(context.TODO(), req)
	if err != nil {
		lager.Logger.Errorf(err, "call request fail.")
		return
	}
	log.Printf("Rest Server catalogue/size [Get] %s", string(resp1.ReadBody()))

	/////////////////////////////////////////////////////////////////////////
	req, _ = rest.NewRequest("GET", "cse://Catalogue/tags")
	req.SetHeader("Content-Type", "application/json")
	//use the invoker like http client.
	resp1, err = invoker.ContextDo(context.TODO(), req)
	if err != nil {
		lager.Logger.Errorf(err, "call request fail.")
		return
	}
	log.Printf("Rest Server tags[Get] %s", string(resp1.ReadBody()))

	/////////////////////////////////////////////////////////////////////////
	req, _ = rest.NewRequest("GET", "cse://Catalogue/catalogue/03fef6ac-1896-4ce8-bd69-b798f85c6e0b")
	req.SetHeader("Content-Type", "application/json")
	//use the invoker like http client.
	resp1, err = invoker.ContextDo(context.TODO(), req)
	if err != nil {
		lager.Logger.Errorf(err, "call request fail.")
		return
	}
	log.Printf("Rest Server catalogue/03fef6ac-1896-4ce8-bd69-b798f85c6e0b[Get] %s", string(resp1.ReadBody()))

	/////////////////////////////////////////////////////////////////////////
	req, _ = rest.NewRequest("GET", "cse://Catalogue/catalogue")
	req.SetHeader("Content-Type", "application/json")
	//use the invoker like http client.
	resp1, err = invoker.ContextDo(context.TODO(), req)
	if err != nil {
		lager.Logger.Errorf(err, "call request fail.")
		return
	}
	log.Printf("Rest Server catalogue[Get] %s", string(resp1.ReadBody()))



	req.Close()
	resp1.Close()
}
func bench(invoker *core.RestInvoker, n int) {
	for i := 0; i < n; i++ {
		marks(invoker)
	}
}

func marks(invoker *core.RestInvoker) {
	defer wg.Done()
	req, _ := rest.NewRequest("GET", "cse://Catalogue/health")
	req.SetHeader("Content-Type", "application/json")
	//use the invoker like http client.
	resp1, err := invoker.ContextDo(context.TODO(), req)
	if err != nil {
		lager.Logger.Errorf(err, "call request fail.")
		fail++
		return
	}
	if resp1.GetStatusCode() == http.StatusOK {
		success++
	} else {
		fail++
	}
	req.Close()
	resp1.Close()
}
