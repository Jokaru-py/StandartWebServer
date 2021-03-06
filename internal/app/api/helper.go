package api

import (
	"StandartWebServer/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

//Попытка откунфигурировать API инстанс(Logger)
func (a *API) configreLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}

	a.logger.SetLevel(log_level)
	return nil
}

//Попытка настроить роутер
func (a *API) configreRouterField() {
	a.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello!"))
	})
}

func (a *API) configreStorageField() error  {
	storage := storage.New(a.config.Storage)
	if err := storage.Open(); err != nil{
		return err
	}

	a.storage = storage
	return nil
}