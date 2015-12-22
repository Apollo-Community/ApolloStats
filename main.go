package main

import (
	"flag"
	"fmt"

	"github.com/lmas/ApolloStats/src"
)

var (
	f_debug = flag.Bool("debug", true, "Run in debug mode")
	f_addr  = flag.String("addr", "127.0.0.1:8000", "Server's listening address")
	f_db    = flag.String("db", "apollo:apollo@/apollo", "Database authentication string")
)

func main() {
	tmp := fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", *f_db)
	db, e := apollostats.OpenDB(tmp)
	if e != nil {
		panic(e)
	}

	i := apollostats.Instance{
		Debug: *f_debug,
		DB:    db,
	}

	i.Init()
	i.Serve(*f_addr)
}
