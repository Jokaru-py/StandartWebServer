package api

type API struct {
}

//API constructor
func New() *API {
	return &API{}
}

func (api *API) Start() error {
	return nil
}
