package main

import (
	"database/sql"
	"fmt"
	"github.com/Dostonlv/simplebank/api"
	db "github.com/Dostonlv/simplebank/db/sqlc"
	"github.com/Dostonlv/simplebank/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	cfg, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	fmt.Println(cfg)
	conn, err := sql.Open("postgres", cfg.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	address := Address{
		HomeNumber: 12,
		Person: Person{
			Id:          1,
			Name:        "test",
			phoneNumber: "234234213",
		},
	}
	fmt.Println(address)

}

type Person struct {
	Id          int
	Name        string
	phoneNumber string
}

type Address struct {
	Person
	HomeNumber int
}
