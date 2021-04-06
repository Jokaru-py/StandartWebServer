package api

type API struct {
	//UNEXPORTED FIELD
	config *Config
}

//API constructor
func New(config *Config) *API {
	return &API{
		config: config,
	}
}

func (api *API) Start() error {
	return nil
}
