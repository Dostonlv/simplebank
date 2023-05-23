package db

import (
	"database/sql"
	"github.com/Dostonlv/simplebank/util"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	cfg, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	testDB, err = sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
