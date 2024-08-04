package main

// API приложения.
type API struct{}

func (api *API) Serve() {}

// БД приложения.
type Database struct{}

func (db *Database) Data() {}

// Сервер является фасадом для API и БД.
type Server struct {
	api *API
	db  *Database
}

func main() {
	var s Server
	s.api = new(API)
	s.db = new(Database)
	// С помощью экземпляра сервера осуществляется
	// доступ к API и БД.
	s.api.Serve()
	s.db.Data()
}
