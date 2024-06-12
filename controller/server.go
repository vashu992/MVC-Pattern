package controller

import (
	"fmt"

	"github.com/vashu992/MVC-Pattern/store"
)

type Server struct {
	PostgresDb store.StoreOperations
}

func (s *Server) NewServer(pgstore store.Postgres) {

	s.PostgresDb = &pgstore
	s.PostgresDb.NewStore()
	fmt.Printf("server = %v\n", s)
}

type ServerOperation interface {
	NewServer(pgstore store.Postgres)
}
