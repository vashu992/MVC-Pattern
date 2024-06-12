package api

import (
	"github.com/vashu992/MVC-Pattern/controller"
	"github.com/vashu992/MVC-Pattern/store"
)

type ApiRoutes struct {
	Server controller.ServerOperation
}

func (api *ApiRoutes) StartApp(server controller.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgres{})
}
