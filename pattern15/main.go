package main

type Database interface {
	Users() []User
}

type User struct{}

type Postgres struct{}

func (*Postgres) Users() []User {
	return []User{}
}

type Mongo struct{}

func (*Mongo) Users() []User {
	return []User{}
}

func NewDB(db string) Database {
	switch db {
	case "postgres":
		return new(Postgres)
	case "mongo":
		return new(Postgres)
	}
	return nil
}

func main() {
	pg := NewDB("postgres")
	pgUser := pg.Users()
	_ = pgUser

	mg := NewDB("mongo")
	mgUser := mg.Users()
	_ = mgUser
}
