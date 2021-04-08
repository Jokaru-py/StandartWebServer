package api

import (
	"StandartWebServer/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type API struct {
	//UNEXPORTED FIELD
	config *Config
	logger *logrus.Logger
	router *mux.Router
	storage *storage.Storage
}

//API constructor
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	if err := api.configreLoggerField(); err != nil {
		return err
	}
	api.logger.Info("Starting api server at port", api.config.BindAddr)

	//Конфигурация роутера.
	api.configreRouterField()

	//Конфигурация бд
	if err := api.configreStorageField(); err != nil {
		return err
	}
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
