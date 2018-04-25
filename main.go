package main

import (
	"flag"
	"github.com/gorilla/mux"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"fmt"
	"runtime"
	"go-jwt-example/core"
	"go-jwt-example/api"
)

const (
	appVersion = "1.0.0"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	envName := flag.String("e", "development", "environment")
	//logLevelStr := flag.String("l", "info", "log level")
	port := flag.Int("p", 8081, "port")
	flag.Parse()

	//logLevel, err := determineLogLevel(*logLevelStr)
	//if err != nil {
	//	panic(err)
	//}

	core := core.New(*envName)
	defer core.Close()
	router := mux.NewRouter()

	api.New(core, router, appVersion)
	//web.New(core, router, appVersion, *envName)

	// start server
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)
	n.UseHandler(router)

	log.Infof("starting %v v%v", *envName, appVersion)
	n.Run(fmt.Sprintf(":%v", *port))
}

func determineLogLevel(logLevel string) (log.Level, error) {
	switch logLevel {
	case "debug":
		return log.DebugLevel, nil
	case "info":
		return log.InfoLevel, nil
	case "warn":
		return log.WarnLevel, nil
	case "error":
		return log.ErrorLevel, nil
	case "fatal":
		return log.FatalLevel, nil
	default:
		return log.DebugLevel, fmt.Errorf("invalid log level '%v'", logLevel)
	}
}



