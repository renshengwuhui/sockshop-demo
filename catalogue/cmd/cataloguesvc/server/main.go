package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	//"code.huawei.com/cse/go-chassis"
	"github.com/ServiceComb/go-chassis/core/lager"
	"github.com/ServiceComb/go-chassis/core/registry"
	"github.com/ServiceComb/go-chassis/examples/catalogue"
	//"github.com/ServiceComb/go-chassis/core/server"
	"github.com/ServiceComb/go-chassis"
	_ "github.com/ServiceComb/go-chassis/server/restful"
	_ "github.com/ServiceComb/go-chassis/third_party/forked/go-micro/server/highway"
	_ "github.com/ServiceComb/go-chassis/third_party/forked/go-micro/transport/tcp"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

const (
	ServiceName = "catalogue"
)

var (
	HTTPLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "request_duration_seconds",
		Help:    "Time (in seconds) spent serving HTTP requests.",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "route", "status_code", "isWS"})
)
var images string

func init() {
	prometheus.MustRegister(HTTPLatency)
}

func main() {

	images := os.Getenv("IMAGE_PATH")
	if images == "" {
		images = "./images"
	}
	fmt.Fprintf(os.Stderr, "images: %q\n", images)
	abs, err := filepath.Abs(images)
	fmt.Fprintf(os.Stderr, "Abs(images): %q (%v)\n", abs, err)
	pwd, err := os.Getwd()
	fmt.Fprintf(os.Stderr, "Getwd: %q (%v)\n", pwd, err)
	files, _ := filepath.Glob(images + "/*")
	fmt.Fprintf(os.Stderr, "ls: %q\n", files) // contains a list of all files in the current directory

	err = chassis.Init()
	if err != nil {
		log.Panicln(err)
	}
	registry.Enable()

	// Mechanical stuff.
	errc := make(chan error)

	// Data domain.
	mysql_ip := os.Getenv("mysql_ip")
	mysql_port := os.Getenv("mysql_port")
	mysql_user := os.Getenv("mysql_user")
	mysql_db := os.Getenv("mysql_db")
	mysql_password := os.Getenv("mysql_password")
	dsn := mysql_user + ":" + mysql_password + "@tcp(" + mysql_ip + ":" + mysql_port + ")/" + mysql_db
	if dsn == "" {
		dsn = "root:@tcp(localhost:3306)/socksdb"
	}
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		lager.Logger.Errorf(err, " err")
		os.Exit(1)
	}
	defer db.Close()

	// Check if DB connection can be made, only for logging purposes, should not fail/exit
	err = db.Ping()
	if err != nil {
		lager.Logger.Errorf(err, "Unable to connect to Database DSN %v", dsn)
	}

	// Service domain.
	var service catalogue.Service
	service = catalogue.NewCatalogueService(db)

	chassis.RegisterSchema("rest", service)
	//if err != nil {
	//	lager.Logger.Errorf(err, "CatalogueServer start failed.")
	//}
	chassis.Run()
	// Capture interrupts.
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	lager.Logger.Error("exit", <-errc)
}
