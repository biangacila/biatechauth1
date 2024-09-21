package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/biangacila/biatechauth1/constants"
	"github.com/biangacila/biatechauth1/infrastructure/adapters/authproviders"
	"github.com/biangacila/biatechauth1/infrastructure/adapters/cassandradb"
	routeHandlers "github.com/biangacila/biatechauth1/interfaces/https/handlers"
	"github.com/biangacila/biatechauth1/internal/utils"
	"github.com/biangacila/biatechauth1/store"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port         int
	Env          string
	Backend      string
	DbHostServer string
	DbName       string
	Version      string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
	Backend     string `json:"backend"`
}

func main() {
	var cfg Config
	flag.IntVar(&cfg.Port, "port", 8080, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "development", "Application environment (development | production)")
	flag.StringVar(&cfg.Backend, "backend", "/backend-biatechauth1/api", "Application environment webservice dispatch main route")
	flag.StringVar(&cfg.DbHostServer, "db-host", "voip.easipath.com", "Database cassandradb host server")
	flag.StringVar(&cfg.DbName, "dbname", "biatechauth01", "Keyspace name")
	flag.StringVar(&cfg.Version, "version", "1.0", "API version")
	flag.Parse()

	// Initialize cassandra database
	cassandradb.InitSession(constants.DbHost, constants.DbName)
	defer cassandradb.CloseSession()

	// Initialize stores
	store.InitTokens()

	// Initialize google authentication
	go authproviders.NewGoogleAuth()

	// Initialize services and controllers
	builder := routeHandlers.NewBuilders()
	controllerHandlers := builder.Build()
	r := routeHandlers.SetupServer(&controllerHandlers)

	// System health controller
	r.HandleFunc(cfg.Backend+"/status", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := AppStatus{
			Status:      "Available",
			Environment: cfg.Env,
			Version:     cfg.Version,
			Backend:     cfg.Backend,
		}
		js, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err1 := w.Write(js)
		if err1 != nil {
			return
		}
	})

	cors := handlers.AllowedOrigins(utils.GetAllowedOrigins())
	headersOk := handlers.AllowedHeaders(utils.GetAllowedHeaders())
	methodsOk := handlers.AllowedMethods(utils.GetAllowedMethods())

	log.Println("Starting auth01 at port", fmt.Sprintf(":%d", cfg.Port), " ...")
	_err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), handlers.CORS(cors, headersOk, methodsOk)(r))

	if _err != nil {
		log.Printf("\x1B[31mServer exit with error: %s\x1B[39m\n", _err)
		os.Exit(1)
	}
}
